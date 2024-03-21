// Run - функция, которая запускает калькулятор
package internal

import (
	"fmt"
	"pet_calculator/pkg/add"
	"pet_calculator/pkg/checkfunc"
	"pet_calculator/pkg/decor"
)

func Run() bool {
	fmt.Println(decor.DecorPattern(decor.DecorElemetn, decor.DecoreLen))
	fmt.Println("Выберете операцию из списка")
	fmt.Println(decor.FnishString)
	fmt.Println(decor.DecorPattern(decor.DecorElemetn, decor.DecoreLen))
	fmt.Println("SUM или 1")

	flag := true

	input := checkfunc.InputCheck()

	switch {
	case input == "SUM" || input == "1":
		add.Add()
	case input == "Q":
		flag = false
	default:
		fmt.Println("Введённая операция отсутствует")
	}
	return flag
}
