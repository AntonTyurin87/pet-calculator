// InputCheck - функция для проверки входящей строки при выборе действия
package checkfunc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputCheck() string {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Некорректный ввод")
	}
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)

	return input
}
