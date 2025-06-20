package dice

import "vtt_api/models"

func ExecuteRollMath(rollCommandResult *models.DiceCommandResult) {
	var res float32
	nextOp := ""

	if len((*rollCommandResult).Results) == 1 {
		(*rollCommandResult).Total = float32((*rollCommandResult).Results[0].DieValue)
		return
	}

	for _, roll := range (*rollCommandResult).Results {
		switch nextOp {
		case "":
			res += float32(roll.DieValue)
		case "-":
			res -= float32(roll.DieValue)
		case "+":
			res += float32(roll.DieValue)
		case "*":
			res *= float32(roll.DieValue)
		case "/":
			res /= float32(roll.DieValue)
		}

		nextOp = roll.MathOperation
	}

	(*rollCommandResult).Total = float32(res)
}
