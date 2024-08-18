package main

import "fmt"

const U2E = 0.90145
const U2R = 89.46

func main() {
	E2R := U2R / U2E
	fmt.Println(E2R)
}
