package command

import (
	"strings"
	"fmt"
)

func parseToCommand(str string) (Command , error) {
	parts := strings.Split(str , " ")
	fmt.Println(parts)
	return Command{}, nil
}
