// Parting - функция для декорирования выхода из исполняемой функции
package decor

import (
	"fmt"
)

func Parting(answer string) {

	var delimetr, answerString, resultString string

	//Верхний и нижний разделитель
	delimetr = decorPattern(DecorElemetn, DecoreLen) + "\n"

	//Строка с результатом операции
	answerString = fmt.Sprintf("Суммирование завершено\n\nТекущий результат %s\n", answer)

	//Результирующая строка
	resultString = "\n" + delimetr + answerString + delimetr

	fmt.Println(resultString)
}
