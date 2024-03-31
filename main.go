package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func smoothMoveMouse(startX, startY, endX, endY int) {
	dx := endX - startX
	dy := endY - startY
	steps := int(math.Hypot(float64(dx), float64(dy))) / 10 // Keep the movement speed reasonable

	for i := 0; i < steps; i++ {
		x := startX + dx*i/steps
		y := startY + dy*i/steps

		robotgo.MoveMouse(x, y)
		time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond) // Short delay for smoother movement
	}
}

func main() {
	screenWidth, screenHeight := robotgo.GetScreenSize()

	x, y := robotgo.Location()

	for {
		// Increase the range for target coordinates to allow longer distance movements
		targetX := x + rand.Intn(600) - 300 // Now the mouse can move a longer distance
		targetY := y + rand.Intn(600) - 300

		if targetX < 0 {
			targetX = 0
		} else if targetX >= screenWidth {
			targetX = screenWidth - 1
		}
		if targetY < 0 {
			targetY = 0
		} else if targetY >= screenHeight {
			targetY = screenHeight - 1
		}

		smoothMoveMouse(x, y, targetX, targetY)
		x, y = targetX, targetY

		fmt.Printf("Moved mouse to (%d, %d)\n", x, y)

		if rand.Float32() < 0.1 { // 10% chance to right click
			robotgo.Click("right", false)
			fmt.Println("Performed a right click")
		}

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Short wait before the next move
	}
}
