package util

import (
	"math/rand"
	"strings"
	"time"
)

var utilSource rand.Source
var utilRand rand.Rand

func init() {
	utilSource = rand.NewSource(time.Now().UnixNano())
	utilRand = *rand.New(utilSource)
}

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func randStirng(leanght int) string {

	runes := []rune(alphabet)

	var st strings.Builder

	for i := 0; i < leanght; i++ {
		st.WriteRune(runes[utilRand.Intn(leanght)])
	}

	return st.String()

}

func RandomRestaurantName() string {
	return "restaurant " + randStirng(10)
}

func RandomDishName() string {
	return "dish " + randStirng(10)
}
