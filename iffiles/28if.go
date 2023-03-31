package iffiles

import (
	"fmt"
	"log"
)

func Abr28() (year int, err error) {

	fmt.Println("Введите год")
	_, err = fmt.Scan(&year)
	if err != nil {
		log.Println("Ошибка ввода")
		return
	}
	if year%4 == 0 {
		if year%100 != 0 && year%4 == 0 {
			fmt.Println("В этом году 366 дней, год високостный")
			return
		} else if year%4 == 0 && year%100 == 0 && year%400 == 0 {
			fmt.Println("В этом году 366 дней, год високостный")
			return
		}
	} else {
		fmt.Println("В этом году 365 дней")
		return
	}
	return
}
