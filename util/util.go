package util

import (
	"github.com/gizak/termui"
	"math/rand"
	"time"
)

// ランダムなカラー値を取得する
func GetColorRand() termui.Attribute {
	return termui.ColorRGB(getColorRandInt(), getColorRandInt(), getColorRandInt())
}

func getColorRandInt() int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(5)
	return result
}

// Byte to GB
func Byte2GB(byte float64) float64 {
	return float64(byte / 1024 / 1024 / 1024)
}

// Byte to GB
func Byte2GBi(byte uint64) float64 {
	return Byte2GB(float64(byte))
}

// Byte to MB
func Byte2MB(byte float64) float64 {
	return float64(byte / 1024 / 1024)
}

// Byte to MB
func Byte2MBi(byte uint64) float64 {
	return Byte2MB(float64(byte))
}

// Byte to KB
func Byte2KB(byte float64) float64 {
	return float64(byte / 1024)
}

// Byte to KB
func Byte2KBi(byte uint64) float64 {
	return Byte2KB(float64(byte))
}
