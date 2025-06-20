package dice

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	"vtt_api/models"
)

type RollModifier struct {
	res *models.DiceCommandResult
}

func (r *RollModifier) rollAndKeep(kind string, amountToKeep int) {
	lenResults := len((*r.res).Results)

	sort.Slice((*r.res).Results, func(i, j int) bool {
		return ((*r.res).Results)[i].DieValue > ((*r.res).Results)[j].DieValue
	})

	amountToInvalidate := lenResults - amountToKeep

	if slices.Contains(TK_KEEP_HIGH, kind) {
		for i := range amountToInvalidate {
			(*r.res).Results[lenResults-i-1].Valid = false
		}
	} else {
		if slices.Contains(TK_KEEP_LOW, kind) {
			for i := range amountToInvalidate {
				(*r.res).Results[i].Valid = false
			}
		}

		if slices.Contains(TK_KEEP_MEDIUM, kind) {
			dicesToEliminate := (amountToInvalidate) / 2

			for i := range dicesToEliminate {
				(*r.res).Results[i].Valid = false
				(*r.res).Results[lenResults-i-1].Valid = false
			}
		}
	}
}

func (r *RollModifier) successFail(kind string, threshold int) {
	for i := range (*r.res).Results {
		index := i
		var value = 1
		selectedValue := (*r.res).Results[i].DieValue >= threshold
		if slices.Contains(TK_FAIL, kind) {
			value = -1
			selectedValue = (*r.res).Results[i].DieValue <= threshold
		}

		(*r.res).Results[i].SuccessFail = value

		if selectedValue {
			(*r.res).Results[index].SuccessFail = value
		}
	}
}

func (r *RollModifier) yingYang() {
	validCounts := 0
	for i := range (*r.res).Results {
		if (*r.res).Results[i].Valid {
			if validCounts == 0 {
				(*r.res).Results[i].YingOrYang = "Ying"
				validCounts++
			} else if validCounts == 1 {
				(*r.res).Results[i].YingOrYang = "Yang"
				validCounts++
			} else {
				(*r.res).Results[i].Valid = false
			}
		}
	}
}

func (r *RollModifier) redOrBlue() {
	validCounts := 0
	redDice := 0
	blueDice := 0
	for i := range (*r.res).Results {
		if (*r.res).Results[i].Valid {
			if validCounts == 0 {
				redDice = i
				(*r.res).Results[i].Color = "RED"
				validCounts++
			} else if validCounts == 1 {
				blueDice = i
				(*r.res).Results[i].Color = "BLUE"
				validCounts++
			} else {
				break
			}
		}
	}

	blueValue := (*r.res).Results[blueDice].DieValue
	redValue := (*r.res).Results[redDice].DieValue
	winner := "DRAW"
	if redValue > blueValue {
		winner = "RED"
	} else if redValue < blueValue {
		winner = "BLUE"
	}
	(*r.res).RedOrBlue = winner
}

func (r *RollModifier) checkAllRoolsEquals() (bool, int) {
	if len((*r.res).Results) == 0 {
		return false, 0
	}
	firstValid := 0
	foundValid := false

	for i, result := range (*r.res).Results {
		if result.Valid {
			firstValid = i
			foundValid = true
		}
	}

	if !foundValid {
		return false, 0
	}

	firstValue := (*r.res).Results[firstValid].DieValue

	for _, result := range (*r.res).Results[firstValid:] {
		if result.DieValue != firstValue && result.Valid {
			return false, 0
		}
	}
	return true, firstValue
}

func (r *RollModifier) criticalSuccessFail(kind string, threshold int) {

	for i, roll := range (*r.res).Results {
		if roll.DieValue >= threshold && slices.Contains(TK_CRIT_SUCCESS_VAL, kind) {
			(*r.res).AmountOfCriticalSuccess++
			(*r.res).Results[i].WasCriticalSuccess = true
		} else if roll.DieValue <= threshold && slices.Contains(TK_CRIT_FAIL_VAL, kind) {
			(*r.res).AmountOfCriticalFail++
			(*r.res).Results[i].WasCriticalFail = true
		} else {
			allValuesAreEquals, value := r.checkAllRoolsEquals()
			if !allValuesAreEquals {
				return
			}
			if value >= threshold && slices.Contains(TK_CRIT_SUCCESS_EQUALS, kind) {
				(*r.res).AmountOfCriticalSuccess++
				(*r.res).Results[i].WasCriticalSuccess = true
			} else if value <= threshold && slices.Contains(TK_CRIT_FAIL_EQUALS, kind) {
				(*r.res).AmountOfCriticalFail++
				(*r.res).Results[i].WasCriticalFail = true
			}
		}
	}
}

func (r *RollModifier) Apply() error {
	for _, roll := range (*r.res).Results {
		curAction := ""
		amount := 0
		waitingForAmount := false
		readyToApply := false

		for _, tk := range roll.Modifiers {

			if curAction == "" {
				if slices.Contains(ROLL_DEFINITION_TK, tk) {
					curAction = tk
				} else {
					return fmt.Errorf("token missplaced: %s", tk)
				}
			} else {
				if NUM_REGEX.MatchString(tk) && waitingForAmount {
					amount, _ = strconv.Atoi(tk)
					waitingForAmount = false
					readyToApply = true
				} else {

					actionNeedsAmount := !slices.Contains(ALL_MOD_WHO_NOT_NEED_AMOUNT, curAction)
					if actionNeedsAmount {
						amount = 0
						waitingForAmount = true
						readyToApply = false
					} else {
						readyToApply = true
					}
				}
				curAction = ""
			}

			if readyToApply {
				switch true {
				case slices.Contains(TK_KEEP, curAction):
					r.rollAndKeep(curAction, amount)
				case slices.Contains(TK_SUCESS_FAIL, curAction):
					r.successFail(curAction, amount)
				case slices.Contains(TK_YING_YANG, curAction):
					r.yingYang()
				case slices.Contains(TK_RED_OR_BLUE, curAction):
					r.redOrBlue()
				case slices.Contains(TK_CRT_VAL, curAction):
					r.criticalSuccessFail(curAction, amount)
				case slices.Contains(TK_CRT_EQUALS, curAction):
					r.criticalSuccessFail(curAction, amount)
				}
			}
		}
	}

	return nil
}
