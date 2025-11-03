package promptdata

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptData(prompt ...any) string {
	fmt.Println(prompt...)
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	data = strings.TrimSpace(data)
	return data
}

func PrintErr(value any) {
	switch t := value.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Printf("Код ошибки: %d", t)
	case error:
		fmt.Println(t.Error())
	default:
		fmt.Println("Неизвестная ошибка!")
	}
}
