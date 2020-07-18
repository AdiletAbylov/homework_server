package model

// WinState constant
const WinState string = "win"

//LostState constant
const LostState string = "lost"

const SourceGame string = "game"
const SourcePayment string = "payment"
const SourceServer string = "server"

// Event struct holds data about event.
type Event struct {
	TransactionID string
	State         string
	Amount        float64 `json:",string"`
	SourceType    string  `json:"omitempty"`
}
