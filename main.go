package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

const (
	CommonSpeed = 20 // speed in percent
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
		drone.Forward(CommonSpeed)
	case keyboard.ArrowDown:
		fmt.Println("BACK")
		drone.Backward(CommonSpeed)
	case keyboard.ArrowLeft:
		fmt.Println("LEFT")
		drone.Left(CommonSpeed)
	case keyboard.ArrowRight:
		fmt.Println("RIGHT")
		drone.Right(CommonSpeed)
	case keyboard.X:
		fmt.Println("LAND")
		drone.Land()
	case keyboard.T:
		fmt.Println("TAKEOFF")
		drone.TakeOff()
	case keyboard.W:
		fmt.Println("UP")
	case keyboard.A:
		fmt.Println("CLOCKWISE ROTATION")
		drone.Clockwise(CommonSpeed)
	case keyboard.D:
		fmt.Println("COUNTERCLOCKWISE ROTATION")
		drone.CounterClockwise(CommonSpeed)
	case keyboard.R:
		fmt.Println("CEASE ROTATION")
		drone.CeaseRotation()
	case keyboard.Spacebar:
		drone.Hover()
		fmt.Println("HOVERING")

	default:
		fmt.Printf("%d - (%v)\n", k.Key, k.Char)
	}
}
