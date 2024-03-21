// CompleteAction - функция для декорирования завершения вычисления действия
package decor

import "fmt"

func CompleteAction() {
	fmt.Println(FnishString)
	fmt.Println(DecorPattern(DecorElemetn, DecoreLen))
}
