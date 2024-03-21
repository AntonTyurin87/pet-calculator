// Add - cуммирование введённых чисел
// Результат суммирования в виде строки
// Если при суммировани вместо числа типа "int" или "float"
// будут введены данные типа "string", то дальнейшая сумма будет конкатенацией строк.
package add

import (
	"pet_calculator/pkg/checkfunc"
	"pet_calculator/pkg/decor"
)

func Add() {

	var result, number float64
	var answer string

	//Переменная для работы цикла вычислений
	repeat := true

	//Значение ответа по умолчанию
	answer = "0"

	decor.Greetings("Cуммирование чисел")

	//Цикл для повторения суммирования
	for repeat {
		decor.AnswerNow(answer)

		number, repeat = checkfunc.NumCheck()
		result += number

		answer = decor.AnswerTrim(result)

		decor.CompleteAction()
	}

	decor.Parting(answer)

	//os.Exit(0)
}
