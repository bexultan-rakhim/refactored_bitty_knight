package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// MARK:COLORS
func brightYellow() rl.Color {
	return rl.NewColor(uint8(255), uint8(240), uint8(31), 255)
}
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}
func ranDarkGreen() rl.Color {
	return rl.NewColor(uint8(rInt(0, 30)), uint8(rInt(40, 90)), uint8(rInt(0, 40)), 255)
}
func ranGreen() rl.Color {
	return rl.NewColor(uint8(rInt(0, 60)), uint8(rInt(140, 256)), uint8(rInt(0, 60)), 255)
}
func ranRed() rl.Color {
	return rl.NewColor(uint8(rInt(140, 256)), uint8(rInt(0, 60)), uint8(rInt(0, 60)), 255)
}
func ranPink() rl.Color {
	return rl.NewColor(uint8(rInt(200, 256)), uint8(rInt(10, 110)), uint8(rInt(130, 180)), 255)
}
func ranBlue() rl.Color {
	return rl.NewColor(uint8(rInt(0, 60)), uint8(rInt(0, 60)), uint8(rInt(140, 256)), 255)
}
func ranDarkBlue() rl.Color {
	return rl.NewColor(uint8(rInt(0, 20)), uint8(rInt(0, 20)), uint8(rInt(100, 160)), 255)
}
func ranCyan() rl.Color {
	return rl.NewColor(uint8(rInt(0, 120)), uint8(rInt(200, 256)), uint8(rInt(150, 256)), 255)
}
func ranYellow() rl.Color {
	return rl.NewColor(uint8(rInt(245, 256)), uint8(rInt(200, 256)), uint8(rInt(0, 100)), 255)
}
func ranOrange() rl.Color {
	return rl.NewColor(uint8(255), uint8(rInt(70, 170)), uint8(rInt(0, 50)), 255)
}
func ranBrown() rl.Color {
	return rl.NewColor(uint8(rInt(100, 150)), uint8(rInt(50, 120)), uint8(rInt(30, 90)), 255)
}
func ranGrey() rl.Color {
	return rl.NewColor(uint8(rInt(170, 220)), uint8(rInt(170, 220)), uint8(rInt(170, 220)), 255)
}
func ranDarkGrey() rl.Color {
	return rl.NewColor(uint8(rInt(90, 120)), uint8(rInt(90, 120)), uint8(rInt(90, 120)), 255)
}
func darkRed() rl.Color {
	return rl.NewColor(uint8(139), uint8(0), uint8(0), 255)
}

