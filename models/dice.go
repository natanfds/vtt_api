package models

type DieRollResult struct {
	DieValue           int
	Valid              bool
	SuccessFail        int
	YingOrYang         string
	Color              string
	WasCriticalSuccess bool
	WasCriticalFail    bool
	MathOperation      string
	Modifiers          []string
}

type DiceCommandResult struct {
	Pattern                 string
	Results                 []DieRollResult
	RedOrBlue               string
	AmountOfCriticalSuccess int
	AmountOfCriticalFail    int
	Total                   float32
}

type RollOperation struct {
	Roll          []string
	MathOperation string
}
