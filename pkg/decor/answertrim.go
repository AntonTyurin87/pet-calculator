// AnswerTrim - функция для подрезки незначащих нулей у ответа выводит результат в виде строки.
package decor

import (
	"strconv"
	"strings"
)

func AnswerTrim(result float64) string {
	answer := strconv.FormatFloat(result, 'g', -1, 64)
	if answer != "0" {
		answer = strings.TrimRight(answer, "0")
	}
	return answer
}
