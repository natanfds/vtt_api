package dice

import (
	"fmt"
	"slices"
)

func ValidateRollSintax(tokens []string) error {
	usedKeep := false
	kindOfCritical := ""
	diceIndex := 0
	waitingForAmount := false
	orderCounter := 0
	var justStrTk []string
	for i, el := range tokens {
		tkIsNum := NUM_REGEX.MatchString(el)
		tkIsFudge := slices.Contains(TK_FUDGE, el)

		//unique
		if !tkIsNum {
			if slices.Contains(justStrTk, el) {
				return fmt.Errorf("token %s repeated", el)
			}
			justStrTk = append(justStrTk, el)
		}
		// token exists
		if !(slices.Contains(ALL_TK, el) || tkIsNum) {
			return fmt.Errorf("token %s not allowed", el)
		}

		// fudge dice
		if slices.Contains(TK_DICE, el) {
			diceIndex = i
		}

		// fudge dice
		if tkIsFudge {
			if i-diceIndex > 1 {
				return fmt.Errorf("fudge dice must be right after dice definition")
			}
		}

		//rolls how need amount
		if waitingForAmount {
			amountOfDice := i == diceIndex+1
			if !tkIsNum || (amountOfDice && !(tkIsFudge || tkIsNum)) {
				return fmt.Errorf("amount for %s is required", el)
			}
		}

		onModNeedAmount := !slices.Contains(ALL_MOD_WHO_NOT_NEED_AMOUNT, el)
		isKindOfAmound := tkIsNum || tkIsFudge
		waitingForAmount = onModNeedAmount && !isKindOfAmound
		if i == len(tokens)-1 && waitingForAmount {
			return fmt.Errorf("amount for %s is required", el)
		}

		// roll and keep
		if slices.Contains(TK_KEEP, el) {
			if usedKeep {
				return fmt.Errorf("only one keep token allowed")
			}
			usedKeep = true
		}

		// critical
		if slices.Contains(TK_CRIT, el) {
			currentCriticalKind := ""
			switch true {
			case slices.Contains(TK_CRT_EQUALS, el):
				currentCriticalKind = "eq"
			case slices.Contains(TK_CRT_VAL, el):
				currentCriticalKind = "val"
			default:
				return fmt.Errorf("critical token unregistered on validation function")
			}

			if kindOfCritical == "" {
				kindOfCritical = currentCriticalKind
			} else {
				if currentCriticalKind != kindOfCritical {
					return fmt.Errorf("only one critical token allowed")
				}
			}
		}

		// order
		for j, group := range TK_ORDER {
			inGroup := slices.Contains(group, el)

			if inGroup {
				if j < orderCounter {
					return fmt.Errorf("token %s must be before %s", el, TK_ORDER[:j])
				} else {
					orderCounter = j
				}
			}
		}
	}

	return nil
}
