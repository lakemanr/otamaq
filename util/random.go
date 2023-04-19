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

func randInt32(min, max int32) int32 {
	return min + utilRand.Int31n(max-min+1)
}

func RandomRestaurantName() string {
	return "restaurant " + randStirng(10)
}

func RandomDishName() string {
	return "dish " + randStirng(10)
}

func RandomClientName() string {
	return "client full name " + randStirng(10)
}

func RandomClientLogin() string {
	return "client login " + randStirng(10)
}

func RandomQuantity() int32 {
	return randInt32(1, 10)
}

func RandomOrderSize() int {
	return 1 + utilRand.Intn(10)
}

func RandomNumOrders() int {
	return 10 + utilRand.Intn(20)
}

func RandomID() int32 {
	return randInt32(1, 1000)
}
