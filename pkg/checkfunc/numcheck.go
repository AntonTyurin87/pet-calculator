// NumCheck - считывает введённую строку и проверяет на то, может ли она быть числом.
// Возвращает число в формате Float64 и значение ошибки
// При возникновении ошибки в качестве числа возвращает ноль
package checkfunc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NumCheck() (float64, bool) {
	var input string

	fmt.Fscan(os.Stdin, &input)

	input = strings.ToUpper(input)

	if input == "Q" {
		return 0, false
	}

	input = strings.ReplaceAll(input, ",", ".")

	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		err = errors.New("!!! некорректный ввод !!! повторите ещё раз")
		fmt.Println(err)
		return 0, true
	}

	return num, true
}
