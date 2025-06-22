package dice

import (
	"slices"

	"vtt_api/models"
)

type SumSubEl struct {
	Value         int
	MathOperation string
}

func ExecuteRollMath(rollCommandResult *models.DiceCommandResult) {
	var res float32

	if len((*rollCommandResult).Results) == 1 {
		(*rollCommandResult).Total = float32((*rollCommandResult).Results[0].DieValue)
		return
	}

	sumSubEl := []SumSubEl{}
	stepsToJump := 0
	for i, roll := range (*rollCommandResult).Results {
		for stepsToJump > 0 {
			stepsToJump--
			continue
		}

		currentOperation := roll.MathOperation
		finishMulti := false
		for slices.Contains([]string{"*", "/"}, currentOperation) {
			sumSubEl = append(sumSubEl, SumSubEl{
				Value:         roll.DieValue,
				MathOperation: roll.MathOperation,
			})
			stepsToJump++
			nextRoll := (*rollCommandResult).Results[i+stepsToJump]
			switch currentOperation {
			case "*":
				sumSubEl[len(sumSubEl)-1].Value *= nextRoll.DieValue
			case "/":
				sumSubEl[len(sumSubEl)-1].Value /= nextRoll.DieValue
			}
			currentOperation = nextRoll.MathOperation
			finishMulti = true
		}

		if finishMulti {
			continue
		}

		if slices.Contains([]string{"+", "-", ""}, roll.MathOperation) {
			sumSubEl = append(sumSubEl, SumSubEl{
				Value:         roll.DieValue,
				MathOperation: roll.MathOperation,
			})
		}
	}

	sumOrSub := ""
	for i, el := range sumSubEl {
		if i == 0 {
			sumOrSub = el.MathOperation
			res += float32(el.Value)
			continue
		}

		switch sumOrSub {
		case "-":
			res -= float32(el.Value)
		case "+":
			res += float32(el.Value)
		case "":
			res += float32(el.Value)
		}

		sumOrSub = el.MathOperation
	}

	(*rollCommandResult).Total = float32(res)
}
