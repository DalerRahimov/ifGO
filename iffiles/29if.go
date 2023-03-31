package iffiles

import (
	"fmt"
	"log"
)

func Abr29() (num int, err error) {

	fmt.Println("Введите число")

	_, err = fmt.Scan(&num)
	if err != nil {
		log.Println("Ошибка ввода")
		return
	}

	if num > 0 {
		if num%2 == 0 {
			fmt.Println("Положительное четное число")
			return
		} else {
			fmt.Println("Положительное нечетное число")
			return
		}
		return
	} else if num < 0 {
		if num%2 == 0 {
			fmt.Println("Отрицательное четное число")
			return
		} else {
			fmt.Println("Отрицательное нечетное число")
			return
		}

		return

	} else {
		fmt.Println("Error")
		return
	}
	return
}
