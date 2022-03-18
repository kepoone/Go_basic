package main

import "fmt"

func main() {
	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30
	/* только для новых переменных (дважды использовать оператор := с той же переменной не получиться, выйдет ошибка)
	num := 31
	no new variables on left side of ^=
	*/

	num += 1
	fmt.Println("+=", num)

	// ++num нету
	num++
	fmt.Println("++", num)

	// camelCase - принятный стиль
	userIndex := 10
	// undex_score - не принято
	user_index := 10
	fmt.Println(userIndex, user_index)

	// объявление нескольких переменных
	var weight, height int = 10, 20

	// присваивание в существующие переменные
	weight, height = 11, 21

	// короткое присваивание
	// хотя-бы одна переменная должна быть новой! (то есть при weight, height = 12, 22 компилятор выдаст ошибку)
	weight, age := 12, 22

	fmt.Println(weight, height, age)
}
