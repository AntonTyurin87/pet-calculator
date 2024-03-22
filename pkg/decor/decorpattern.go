// DecorPattern -  возвращает строку из заданного числа подстрок
package decor

func DecorPattern(DecorElemetn string, DecoreLen int) string {
	var resultString string
	for i := DecoreLen; i > 0; i-- {
		resultString += DecorElemetn
	}
	return resultString
}
