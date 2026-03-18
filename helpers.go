package main

import (
	"math"
	"math/rand"
)
// MARK:HELPERS
func absdiff(num1, num2 float32) float32 {
	num := float32(0)
	if num1 == num2 {
		num = 0
	} else {
		if num1 <= 0 && num2 <= 0 {
			num1 = getabs(num1)
			num2 = getabs(num2)
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		} else if num1 <= 0 && num2 >= 0 {
			num = num2 + getabs(num1)
		} else if num2 <= 0 && num1 >= 0 {
			num = num1 + getabs(num2)
		} else if num2 >= 0 && num1 >= 0 {
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		}
	}
	return num
}
func getabs(value float32) float32 {
	value2 := float64(value)
	value = float32(math.Abs(value2))
	return value
}
func remImg(slice []ximg, s int) []ximg {
	return append(slice[:s], slice[s+1:]...)
}
func remEnemy(slice []xenemy, s int) []xenemy {
	return append(slice[:s], slice[s+1:]...)
}
func remTxt(slice []xtxt, s int) []xtxt {
	return append(slice[:s], slice[s+1:]...)
}
func remProj(slice []xproj, s int) []xproj {
	return append(slice[:s], slice[s+1:]...)
}
func remFX(slice []xfx, s int) []xfx {
	return append(slice[:s], slice[s+1:]...)
}
func remBlok(slice []xblok, s int) []xblok {
	return append(slice[:s], slice[s+1:]...)
}
func flipcoin() bool {
	onoff := false
	choose := rInt(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}
func roll6() int {
	return rInt(1, 7)
}
func roll12() int {
	return rInt(1, 13)
}
func roll18() int {
	return rInt(1, 19)
}
func roll24() int {
	return rInt(1, 25)
}
func roll30() int {
	return rInt(1, 31)
}
func roll36() int {
	return rInt(1, 37)
}
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func rI32(min, max int) int32 {
	return int32(min + rand.Intn(max-min))
}
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}


