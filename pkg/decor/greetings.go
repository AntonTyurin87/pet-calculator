// Greetings - функция для декорирования входа в исполняемую функцию
package decor

import (
	"fmt"
	"unicode/utf8"
)

func Greetings(name string) {

	var delimetr, indentationLeft, nameStringRight, nameString, finishString, resultString string

	lenName := utf8.RuneCountInString(name)

	//Верхний и нижний разделитель
	delimetr = decorPattern(DecorElemetn, DecoreLen) + "\n"

	//Отступ перед наименованием действия
	indentationLeft = decorPattern(EndentationElement, (DecoreLen-lenName)/2)

	//Отступ после наименования действия
	nameStringRight = decorPattern(EndentationElement, DecoreLen-(DecoreLen-lenName)/2)

	//Строка с наименованием действия
	nameString = indentationLeft + name + nameStringRight + "\n"

	//Последняя строка
	finishString = "Для завершения введите Q и нажмите Enter без числа"

	//Результирующая строка
	resultString = "\n" + delimetr + nameString + delimetr + finishString

	fmt.Println(resultString)
}
