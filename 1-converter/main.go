package main

import (
	"fmt"
	"strings"
)

const U2E = 0.90145
const U2R = 89.46

func main() {
	cur1, valConv, cur2 := getUserInput()

	calcCur(valConv, cur1, cur2)

}

func getUserInput() (string, float64, string) {
	var cur1 string
	var valConv float64
	var cur2 string
	for {
		fmt.Println("Введите исходную валюту (eur, usd или rub):")
		fmt.Scanln(&cur1)

		if !checkCurIn(&cur1) {
			fmt.Println("Некорректно введена исходная валюта")
			continue
		}

		break
	}

	for {

		fmt.Println("Введите  валюту (Числовой параметр):")
		_, err := fmt.Scan(&valConv)
		fmt.Println(valConv)
		isFloat := fmt.Sprintf("%T", valConv)

		if err != nil || isFloat != "float64" {
			fmt.Println("Некорректно введено значение")
			continue
		}

		break
	}

	for {
		fmt.Println("Введите целевую валюту (eur, usd или rub):")
		fmt.Scanln(&cur2)

		if !checkCurIn(&cur2) {
			fmt.Println("Некорректно введена целевая валюта")
			continue
		}

		if cur1 == cur2 {
			fmt.Println("Исходная и целевая валюты совпадают.")
			continue
		}

		break
	}

	return strings.ToLower(cur1), valConv, strings.ToLower(cur2)

}

func calcCur(val float64, cur1, cur2 string) {
	E2R := U2R / U2E

	switch {
	case cur1 == "usd" && cur2 == "eur":
		fmt.Printf("%.3f \n", val*U2E)
	case cur1 == "eur" && cur2 == "usd":
		fmt.Printf("%.3f \n", val/U2E)
	case cur1 == "rub" && cur2 == "usd":
		fmt.Printf("%.3f \n", val/U2R)
	case cur1 == "usd" && cur2 == "rub":
		fmt.Printf("%.3f \n", val*U2R)
	case cur1 == "rub" && cur2 == "eur":
		fmt.Printf("%.3f \n", val/E2R)
	case cur1 == "eur" && cur2 == "rub":
		fmt.Printf("%.3f \n", val*E2R)
	}

}

func checkCurIn(pCur *string) bool {
	cur := strings.ToLower(*pCur)

	if cur == "usd" || cur == "eur" || cur == "rub" {
		return true
	}

	return false

}
