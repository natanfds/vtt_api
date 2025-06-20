package dice

import (
	"regexp"

	"vtt_api/utils"
)

var (
	TK_DICE                = []string{"d"}
	TK_PLUS                = []string{"+"}
	TK_MINUS               = []string{"-"}
	TK_MULTIPLY            = []string{"*"}
	TK_DIVIDE              = []string{"/"}
	TK_FUDGE               = []string{"df"}
	TK_KEEP_HIGH           = []string{"kh"}
	TK_KEEP_MEDIUM         = []string{"km"}
	TK_KEEP_LOW            = []string{"kl"}
	TK_EXPLODE             = []string{"ex"}
	TK_REROLL              = []string{"re"}
	TK_SUCCESS             = []string{"su"}
	TK_FAIL                = []string{"fa"}
	TK_YING_YANG           = []string{"yy"}
	TK_RED_OR_BLUE         = []string{"rb"}
	TK_CRIT_SUCCESS_VAL    = []string{"csv"}
	TK_CRIT_FAIL_VAL       = []string{"cfv"}
	TK_CRIT_SUCCESS_EQUALS = []string{"cse"}
	TK_CRIT_FAIL_EQUALS    = []string{"cfe"}

	TK_KEEP          = utils.JoinSlices(TK_KEEP_HIGH, TK_KEEP_LOW, TK_KEEP_MEDIUM)
	TK_MATH_SIMBOLS  = utils.JoinSlices(TK_PLUS, TK_DIVIDE, TK_MINUS, TK_MULTIPLY)
	TK_SUCESS_FAIL   = utils.JoinSlices(TK_SUCCESS, TK_FAIL)
	TK_SPECIAL_ROLLS = utils.JoinSlices(TK_YING_YANG, TK_RED_OR_BLUE)
	TK_CRT_VAL       = utils.JoinSlices(TK_CRIT_SUCCESS_VAL, TK_CRIT_FAIL_VAL)
	TK_CRT_EQUALS    = utils.JoinSlices(TK_CRIT_FAIL_EQUALS, TK_CRIT_SUCCESS_EQUALS)
	TK_CRIT          = utils.JoinSlices(TK_CRT_EQUALS, TK_CRT_VAL)
	TK_ROLL          = utils.JoinSlices(TK_DICE, TK_FUDGE)

	TK_ORDER = [][]string{
		TK_ROLL,
		TK_KEEP,
		TK_CRIT,
		TK_EXPLODE,
		TK_REROLL,
		TK_SUCESS_FAIL,
		TK_SPECIAL_ROLLS,
	}

	ALL_TK = utils.JoinSlices(
		TK_DICE,
		TK_FUDGE,
		TK_REROLL,
		TK_EXPLODE,
		TK_KEEP,
		TK_SUCESS_FAIL,
		TK_SPECIAL_ROLLS,
		TK_CRIT,
		TK_MATH_SIMBOLS,
	)

	ROLL_MODIFIERS_TK = utils.JoinSlices(
		TK_KEEP,
		TK_SUCESS_FAIL,
		TK_SPECIAL_ROLLS,
		TK_CRIT,
	)

	BASIC_ROLL_TK = utils.JoinSlices(
		TK_DICE,
		TK_FUDGE,
		TK_REROLL,
		TK_EXPLODE,
	)

	ROLL_DEFINITION_TK = utils.JoinSlices(
		BASIC_ROLL_TK,
		ROLL_MODIFIERS_TK,
	)

	ALL_MOD_WHO_NOT_NEED_AMOUNT = utils.JoinSlices(
		TK_YING_YANG,
		TK_RED_OR_BLUE,
	)

	NUM_REGEX  = regexp.MustCompile(`^\d+$`)
	MATH_REGEX = regexp.MustCompile(`[\+\-\*/]`)
)
