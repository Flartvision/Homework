package main

import (
	"bin/bins"
	"bin/file"
	"bin/storage"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	crBin()

}

func crBin() {
	var stat bool
	id := promptData("Введите id")
	name := promptData("Введите имя")
	sec := promptData("Сделать бин приватным? (Да или нет)")

	switch strings.ToLower(sec) {
	case "да":
		stat = true
	case "нет":
		stat = false
	default:
		fmt.Println("Неккоректно введены данные")
	}
	MyBin := bins.NewBin(id, name, stat)
	st := storage.NewStorage()
	st.AddBin(*MyBin)

	data, err := st.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON")
	}

	file.Wfile(data, "storage.json")
}

func promptData(q string) string {
	var usCh string

	fmt.Println(q)
	fmt.Scanln(&usCh)

	return usCh

}
