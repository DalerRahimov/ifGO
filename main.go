package main

import (
	"HW6/iffiles"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Выберите задачу: 1)if28 2)if29 3)if30")
	var ex int
	var err error
	_, err = fmt.Scan(&ex)
	if err != nil {
		log.Println("Ошибка ввода")
		return
	}

	switch ex {
	case 1:
		iffiles.Abr28()

	case 2:
		iffiles.Abr29()

	case 3:
		iffiles.Abr30()

	default:
		fmt.Println("Ошибка")

	}

}
