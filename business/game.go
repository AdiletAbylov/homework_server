package business

import (
	"fmt"
	"hw_server/model"
)

// HandleGameEvent handles given event.
func HandleGameEvent(event model.GameEvent) {
	fmt.Printf("Business model %+v", event)
}
