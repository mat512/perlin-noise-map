package main

// #cgo LDFLAGS: -lm
// #include "perlin.h"
import "C"

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialization
	screenWidth := int32(1920 / 2)
	screenHeight := int32(1080 / 2)

	rl.InitWindow(screenWidth, screenHeight, "Perlin noise map")

	// Camera
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0.0, 2.0, 0.0)
	camera.Target = rl.NewVector3(10.0, 0.0, 10.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	rl.SetCameraMode(camera, rl.CameraFirstPerson) // Set a free camera mode

	rl.SetTargetFPS(60)

	lineColor := rl.LightGray

	showFPS := false

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera) // Update camera

		if rl.IsKeyPressed(rl.KeyF1) {
			showFPS = !showFPS
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(camera)

		renderDistance := 100
		amplification := C.float(10)
		smoothness := C.float(0.05)

		// Draw lines
		for x := int(camera.Position.X) - renderDistance; x < int(camera.Position.X)+renderDistance; x++ {
			for z := int(camera.Position.Z) - renderDistance; z < int(camera.Position.Z)+renderDistance; z++ {
				topCornerHeight := float32(C.perlin(C.float(x)*smoothness, C.float(z)*smoothness) * amplification)
				rightCornerHeight := float32(C.perlin(C.float(x+1)*smoothness, C.float(z)*smoothness) * amplification)
				botCornerHeight := float32(C.perlin(C.float(x)*smoothness, C.float(z+1)*smoothness) * amplification)

				if x < int(camera.Position.X)+renderDistance-1 && z < int(camera.Position.Z)+renderDistance-1 {
					rl.DrawLine3D(rl.NewVector3(float32(x), topCornerHeight, float32(z)), rl.NewVector3(float32(x)+1, rightCornerHeight, float32(z)), lineColor)
					rl.DrawLine3D(rl.NewVector3(float32(x), topCornerHeight, float32(z)), rl.NewVector3(float32(x), botCornerHeight, float32(z)+1), lineColor)
					rl.DrawLine3D(rl.NewVector3(float32(x), botCornerHeight, float32(z)+1), rl.NewVector3(float32(x)+1, rightCornerHeight, float32(z)), lineColor)
				} else {
					// Draw lines at border
					if x != int(camera.Position.X)+renderDistance-1 {
						rl.DrawLine3D(rl.NewVector3(float32(x), topCornerHeight, float32(z)), rl.NewVector3(float32(x)+1, rightCornerHeight, float32(z)), lineColor)
					}
					if z != int(camera.Position.Z)+renderDistance-1 {
						rl.DrawLine3D(rl.NewVector3(float32(x), topCornerHeight, float32(z)), rl.NewVector3(float32(x), botCornerHeight, float32(z)+1), lineColor)
					}
				}
			}
		}

		rl.EndMode3D()

		if showFPS {
			rl.DrawFPS(10, 40)
		}

		coordinateText := fmt.Sprintf("X: %f Y: %f Z:%f", camera.Position.X, camera.Position.Y, camera.Position.Z)
		rl.DrawText(coordinateText, 10, 10, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
