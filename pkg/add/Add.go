// Add - cуммирование введённых чисел
// Результат суммирования в виде строки
// Если при суммировани вместо числа типа "int" или "float"
// будут введены данные типа "string", то дальнейшая сумма будет конкатенацией строк.
package add

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pet_calculator/pkg/decor"
)

func Add() {

	var result float64
	var input, answer string

	answer = "0"

	decor.Greetings("Cуммирование чисел")

	for input != "Q" {
		fmt.Printf("Текущий результат %s\n", answer)
		fmt.Print("Введите число и нажмите Enter: ")
		fmt.Fscan(os.Stdin, &input)

		input = strings.ReplaceAll(input, ",", ".")

		number, _ := strconv.ParseFloat(input, 64)
		result += number

		answer = strconv.FormatFloat(result, 'g', -1, 64)
		answer = strings.TrimRight(answer, "0")

		fmt.Println("Для завершения введите Q и нажмите Enter без числа")
		fmt.Println("**************************************************")
	}

	decor.Parting(answer)
}
