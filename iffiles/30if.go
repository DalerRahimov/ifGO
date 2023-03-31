package iffiles

import (
	"fmt"
	"log"
)

func Abr30() (num int, err error) {

	fmt.Println("Введите число")

	_, err = fmt.Scan(&num)
	if err != nil {
		log.Println("Ошибка ввода")
		return
	}

	if num > 99 {
		if num%2 == 0 {
			fmt.Println("Трехзначное четное число")
			return
		} else {
			fmt.Println("Трехзначное нечетное число")
			return
		}
		return
	} else if num > 9 {
		if num%2 == 0 {
			fmt.Println("Двухзначное четное число")
			return
		} else {
			fmt.Println("Двухзначное нечетное число")
			return
		}

		return
	} else if num > 0 {
		if num%2 == 0 {
			fmt.Println("Однозначное четное число")
			return
		} else {
			fmt.Println("Однозначное нечетное число")
			return
		}

	} else {
		fmt.Println("Это ноль")
		return

	}

	return
}
