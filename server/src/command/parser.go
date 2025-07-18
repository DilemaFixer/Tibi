package command

import (
	"strings"
	"fmt"
	"log"
)

func ParseToCommand(str string) (*Command , error) {
	if str == "" || str == " " {
		return nil , fmt.Errorf("Error parsing command , input is empty")
	}
	parts := split(str)
	
	command := NewMinimalCommand(parts[0])
	partsCount := len(parts)
	if partsCount > 1 {
		command.Subcommand = parts[1]
	}

	if (partsCount - 2) != 0 {
		signsStr := parts[2:]
		signs := make([]Significance, 0)
		//TODO: refactor this part of code , move it to new func
		for _, signStr := range signsStr {
			signLenght := len(signStr)
			if strings.HasPrefix(signStr, ".") && signLenght > 1 {
				signStr = removeCharsFromBegine(signStr, 1)
				values := strings.Split(signStr, "=")
				if len(values) != 2 {
					return nil , fmt.Errorf("Invalid var seter declaration")
				}
				sign := NewSignificance(Prop, values[0], values[1])
				signs = append(signs, sign)
				continue 
			}

			if strings.HasPrefix(signStr, "--") && signLenght > 2{
				signStr = removeCharsFromBegine(signStr, 2)
				sign := NewSignificance(Flag, signStr, "")	
				signs = append(signs, sign)
				continue
			}
			
			sign := NewSignificance(Data, "", signStr)	
			signs = append(signs, sign)
		}
		command.Significances = signs
	}

	return command , nil
}

func split(str string) []string{
	str = strings.TrimSpace(str)
	//TODO: check is str is nil or empty | len(str) == 0

	cutPoints := make([]int , 1) // cutPoints[0] == 0 (string start)
	var inQuotes bool

	for i, char := range str {
		if char == '"' {
			inQuotes = !inQuotes
		}

		if char == ' ' && !inQuotes {
			if i > 0 && str[i-1] != ' ' {
				cutPoints = append(cutPoints , i)
			}
		}
	}

	cutPoints = append(cutPoints , len(str))
	var result []string
	for i := 0; i < len(cutPoints)-1; i++ {
		start := cutPoints[i]
		end := cutPoints[i+1]

		for start < end && str[start] == ' ' {
			start++
		}

		for end > start && str[end-1] == ' ' {
			end--
		}

		if start < end {
			segment := str[start:end]
			if len(segment) > 2 && segment[0] == '"' && 
			segment[len(segment)-1] == '"'{
				segment = segment[1:len(segment)-1]

				}
			result = append(result , segment)
		}
	}
	return result
}

func removeCharsFromBegine(str string , count int) string {
	//TODO: check len(str) < count return "" | check is string is empty

	result := str[count:]
	return result
}

