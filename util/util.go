package util

import (
	"github.com/gizak/termui"
	"time"
	"math/rand"
)

func GetColorRand() termui.Attribute {
	return termui.ColorRGB(getColorRandInt(), getColorRandInt(), getColorRandInt())
}

func getColorRandInt() int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(5)
	return result
}

func Byte2GB(byte float64) float64 {
	return float64(byte / 1024 / 1024 / 1024)
}

func Byte2GBi(byte uint64) float64 {
	return Byte2GB(float64(byte))
}

func Byte2MB(byte float64) float64 {
	return float64(byte / 1024 / 1024)
}

func Byte2MBi(byte uint64) float64 {
	return Byte2MB(float64(byte))
}

func Byte2KB(byte float64) float64 {
	return float64(byte / 1024)
}

func Byte2KBi(byte uint64) float64 {
	return Byte2KB(float64(byte))
}



