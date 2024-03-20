// decorPattern -  возвращает строку из заданного числа подстрок
package decor

func decorPattern(decorElemetn string, decoreLen int) string {
	var resultString string
	for i := decoreLen; i > 0; i-- {
		resultString += decorElemetn
	}
	return resultString
}
