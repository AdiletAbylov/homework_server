package model

// GameEvent struct holds data about game event.
type GameEvent struct {
	TransactionID string
	State         string
	Amount        float32 `json:",string"`
	SourceType    string  `json:"omitempty"`
}
