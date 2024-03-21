// AnswerNow - функция для оформления вывода текущего результата
package decor

import "fmt"

func AnswerNow(answer string) {
	fmt.Printf("Текущий результат %s\n", answer)
	fmt.Print("Введите число и нажмите Enter: ")
}
