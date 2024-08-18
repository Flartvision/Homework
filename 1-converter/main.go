package main

import "fmt"

const U2E = 0.90145
const U2R = 89.46

func main() {
	val, cur1, cur2 := getUserInput()
	res := calcCur()
	E2R := U2R / U2E
	fmt.Println(E2R)
}

func getUserInput() (string, string, string) {
	val := fmt.Sprintln("Введите  валюту (EUR, USD или RUB):")
	cur1 := fmt.Sprintln("Введите исходную валюту (EUR, USD или RUB):")
	cur2 := fmt.Sprint("Введите  целевую валюту (EUR, USD или RUB):")
	return val, cur1, cur2
}

func calcCur(val float64, cur1, cur2 string) float64 {

}
