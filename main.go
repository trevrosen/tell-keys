package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

var drone = tello.NewDriver("8888")

// TODO: consume events from the drone driver
// TODO: unify both drone and keyboard into one Robot?
func main() {
	keys := keyboard.NewDriver()
	drone.Start()

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			handleKeyEvent(key)
		})
	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{},
		[]gobot.Device{keys},
		work,
	)

	robot.Start()
}

func handleKeyEvent(k keyboard.KeyEvent) {
	switch k.Key {
	case keyboard.ArrowUp:
		fmt.Println("FORWARD")
		drone.Forward(20)
	case keyboard.ArrowDown:
		fmt.Println("BACK")
		drone.Backward(20)
	case keyboard.ArrowLeft:
		fmt.Println("LEFT")
		drone.Left(20)
	case keyboard.ArrowRight:
		fmt.Println("RIGHT")
		drone.Right(20)
	case keyboard.X:
		fmt.Println("LAND")
		drone.Land()
	case keyboard.T:
		fmt.Println("TAKEOFF")
		drone.TakeOff()
	case keyboard.W:
		fmt.Println("UP")
	case keyboard.A:
		fmt.Println("rotate CLOCKWISE")
	case keyboard.D:
		fmt.Println("rotate COUNTERCLOCKWISE")
	default:
		fmt.Printf("(%v)\n", k.Char)
	}
}
