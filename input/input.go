package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)
type InputState struct {
    UseController        bool
    IsController         bool
    ControllerOn         bool
    ControllerDisconnect bool
    ControllerWasOn      bool
    KeypressT            int32
}

func checkController(core *CoreState) {

}
