// Greetings - функция для декорирования входа в исполняемую функцию
package decor

import (
	"fmt"
	"unicode/utf8"
)

func Greetings(name string) {

	var delimetr, indentationLeft, nameStringRight, nameString, resultString string

	lenName := utf8.RuneCountInString(name)

	//Верхний и нижний разделитель
	delimetr = DecorPattern(DecorElemetn, DecoreLen) + "\n"

	//Отступ перед наименованием действия
	indentationLeft = DecorPattern(EndentationElement, (DecoreLen-lenName)/2)

	//Отступ после наименования действия
	nameStringRight = DecorPattern(EndentationElement, DecoreLen-(DecoreLen-lenName)/2)

	//Строка с наименованием действия
	nameString = indentationLeft + name + nameStringRight + "\n"

	//Результирующая строка
	resultString = "\n" + delimetr + nameString + delimetr + FnishString

	fmt.Println(resultString)
}
