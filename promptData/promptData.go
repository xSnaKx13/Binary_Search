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
