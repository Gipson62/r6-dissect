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

type PlayerPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type EntityPositions struct {
	EntityID  byte             `json:"entityId"`
	Team      string           `json:"team"`
	Positions []PlayerPosition `json:"positions"`
}
