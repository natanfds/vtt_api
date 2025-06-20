package dice

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"vtt_api/models"
	"vtt_api/utils"
)

func RollDices(message string) (models.DiceCommandResult, error) {

	args := strings.ToLower(strings.Join(strings.Split(message, "")[1:], ""))
	mathMatches := MATH_REGEX.FindAllStringIndex(args, -1)

	//Separando operações de rolagem
	var operations = []models.RollOperation{}
	if len(mathMatches) > 0 {
		for i, match := range mathMatches {
			roll := args[:match[0]]
			mathOperation := string(args[match[0]:match[1]])
			if i > 0 {
				roll = args[mathMatches[i-1][1]:match[0]]
			}

			rollTokens, err := separateTokensInOperation(roll)
			if err != nil {
				return models.DiceCommandResult{}, err
			}
			operations = append(operations, models.RollOperation{
				Roll:          rollTokens,
				MathOperation: mathOperation,
			})

			if i >= len(mathMatches)-1 {
				rollTokens, err := separateTokensInOperation(args[mathMatches[i][1]:])
				if err != nil {
					return models.DiceCommandResult{}, err
				}
				operations = append(operations, models.RollOperation{
					Roll:          rollTokens,
					MathOperation: "",
				})
			}
		}
	} else {
		rollTokens, err := separateTokensInOperation(args)
		if err != nil {
			return models.DiceCommandResult{}, err
		}
		operations = append(operations, models.RollOperation{
			Roll:          rollTokens,
			MathOperation: "",
		})
	}

	diceRollRes := []models.DieRollResult{}
	for _, op := range operations {
		rollRes, err := rollDices(op)
		if err != nil {
			return models.DiceCommandResult{}, err
		}
		diceRollRes = append(diceRollRes, rollRes...)
	}

	commandRes := models.DiceCommandResult{
		Results: diceRollRes,
		Pattern: args,
	}

	rollModifier := RollModifier{&commandRes}
	rollModifier.Apply()

	ExecuteRollMath(&commandRes)

	return commandRes, nil

}

func separateTokensInOperation(operation string) ([]string, error) {
	numMatches := NUM_REGEX.FindAllStringIndex(operation, -1)
	tokensRoll := [][]int{}

	for _, tk := range ROLL_DEFINITION_TK {
		if strings.Contains(operation, tk) {
			indexFound := utils.FindStringIndex(operation, tk)
			if len(indexFound) > 1 {
				return nil, fmt.Errorf("token %s found more than once", tk)
			}
			tokensRoll = utils.JoinSlices(tokensRoll, indexFound)
		}
	}

	allMatches := utils.JoinSlices(numMatches, tokensRoll)
	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i][0] < allMatches[j][0]
	})

	var res []string

	for _, match := range allMatches {
		res = append(res, operation[match[0]:match[1]])
	}

	return res, nil
}

func rollDices(rollOperation models.RollOperation) ([]models.DieRollResult, error) {
	var amount int
	var size int
	results := []models.DieRollResult{}
	var waitingForSize bool
	var startOfModifiers int
	var willExplode bool
	var willReroll bool
	var rerollUnder int

	//pure math operation
	if len(rollOperation.Roll) == 1 {
		justNum := NUM_REGEX.MatchString(rollOperation.Roll[0])
		if justNum {
			res := []models.DieRollResult{}
			value, _ := strconv.Atoi(rollOperation.Roll[0])
			res = append(res, models.DieRollResult{
				DieValue:      value,
				Valid:         true,
				MathOperation: rollOperation.MathOperation,
			})
			return res, nil
		} else {
			if !slices.Contains(TK_FUDGE, rollOperation.Roll[0]) {
				return nil, fmt.Errorf("too short")
			}
		}
	}

	for i, op := range rollOperation.Roll {
		isNum := NUM_REGEX.MatchString(op)

		if !(isNum || slices.Contains(BASIC_ROLL_TK, op)) {
			if i == 0 {
				return nil, fmt.Errorf("expected dice definition wrong operation received %s", op)
			} else {
				startOfModifiers = i + 1
				break
			}
		} else {
			if i == len(rollOperation.Roll)-1 {
				startOfModifiers = i + 1
			}
		}
		// end of process

		if slices.Contains(TK_EXPLODE, op) {
			willExplode = true
		}

		if slices.Contains(TK_REROLL, op) {
			willReroll = true
		}

		if willReroll {
			if isNum {
				rerollUnder, _ = strconv.Atoi(op)
			} else {
				return nil, fmt.Errorf("expected reroll threshold")
			}
		}

		if slices.Contains(TK_DICE, op) {
			waitingForSize = true
			if i == 0 {
				amount = 1
			}
			continue
		}

		if i == 0 {
			if isNum {
				amount, _ = strconv.Atoi(op)
			} else {
				return nil, fmt.Errorf("invalid initial operation %s", op)
			}
		} else {
			if waitingForSize {
				if isNum {
					size, _ = strconv.Atoi(op)
					waitingForSize = false
				} else if slices.Contains(TK_FUDGE, op) {
					size = -1
					waitingForSize = false
				} else {
					return nil, fmt.Errorf("expected size of dices %s", op)
				}
			}
		}
	}

	for range amount {
		//explode
		for {
			preResult := executeRollDie(
				size,
				rollOperation.Roll[startOfModifiers:],
				rollOperation.MathOperation,
			)
			results = append(results, preResult)

			//reroll
			if willReroll && rerollUnder >= preResult.DieValue {
				secondResult := executeRollDie(
					size,
					rollOperation.Roll[startOfModifiers:],
					rollOperation.MathOperation,
				)

				secondBigger := preResult.DieValue < secondResult.DieValue

				if secondBigger {
					preResult.Valid = false
				} else {
					secondResult.Valid = false
				}
				results = append(results, secondResult)
			}

			if !willExplode {
				break
			} else {
				if preResult.DieValue != size {
					break
				}
			}
		}
	}

	return results, nil
}

func executeRollDie(size int, modifiers []string, mathOperation string) models.DieRollResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := 1
	max := size
	if size == -1 {
		min = -1
		max = 1
	}

	randomNum := r.Intn(max-min+1) + min
	preResult := models.DieRollResult{
		DieValue:      randomNum,
		Valid:         true,
		Modifiers:     modifiers,
		MathOperation: mathOperation,
	}

	return preResult
}
