package main

import (
	"fmt"
	"strings"
)

// Заменить switch на map

var curMap = map[string]map[string]float64{
	"EUR": {"RUB": 101.73, "USD": 1.11},
	"USD": {"RUB": 91.49, "EUR": 0.90},
	"RUB": {"EUR": 0.0098, "USD": 0.011},
}

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
		//fmt.Println(valConv)
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

	return strings.ToUpper(cur1), valConv, strings.ToUpper(cur2)

}

func calcCur(val float64, cur1, cur2 string) {
	fmt.Printf("Курс %s к %s составляет: %.2f %s. \n", cur1, cur2, curMap[cur1][cur2]*val, cur1)
}

func checkCurIn(pCur *string) bool {
	cur := strings.ToUpper(*pCur)

	if cur == "USD" || cur == "EUR" || cur == "RUB" {
		return true
	}

	return false

}
