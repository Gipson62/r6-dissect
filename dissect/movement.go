package dissect

type UtilityEvent struct {
	Operator      string  `json:"operator"`
	Username      string  `json:"username"`
	UtilityType   string  `json:"utilityType"`
	Action        string  `json:"action"`
	TimeInSeconds float64 `json:"time"`
}

type CameraDestruction struct {
	DestroyedBy   string  `json:"destroyedBy"`
	TimeInSeconds float64 `json:"time"`
}

