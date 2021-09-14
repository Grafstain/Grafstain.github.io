package main

import (
	"fmt"
)

func main2() {
	//s := rand.Intn(10) + 1
	//fmt.Println(s)
	//fmt.Printf("Привет \n%-40s %s", "Ramzan", "Krasavchik")
	checkSignature(true)
	checkSignature(false)
}

func checkSignature(Signature bool) {
	switch Signature {
	case true:
		fmt.Println("It's True")
	default:
		fmt.Println("It's False")
	}
}
