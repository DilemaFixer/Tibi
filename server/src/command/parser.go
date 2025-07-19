package command

import (
	"fmt"
	"strings"
	"unicode"
)

func ParseToCommand(str string) (*Command, error) {
	strings.TrimSpace(str)

	if str == "" || str == " " {
		return nil, fmt.Errorf("Error parsing command , input is empty")
	}

	parts := split(str)
	command := NewMinimalCommand(parts[0])
	setSubcommandIfCan(command, parts)
	err := setSignificancesIfCan(command , parts)
	if err != nil {
		return nil , err
	}
	return command, nil
}

func split(str string) []string {
	str = strings.TrimSpace(str)

	cutPoints := make([]int, 1) // cutPoints[0] == 0 (string start)
	var inQuotes bool

	for i, char := range str {
		if char == '"' {
			inQuotes = !inQuotes
		}

		if char == ' ' && !inQuotes {
			if i > 0 && str[i-1] != ' ' {
				cutPoints = append(cutPoints, i)
			}
		}
	}

	cutPoints = append(cutPoints, len(str))
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
				segment[len(segment)-1] == '"' {
				segment = segment[1 : len(segment)-1]

			}
			result = append(result, segment)
		}
	}
	return result
}


func setSubcommandIfCan(command *Command, parts []string) {
	partsCount := len(parts)
	if partsCount > 1 {
		command.Subcommand = parts[1]
	}
}

func setSignificancesIfCan(command *Command, parts []string) error {
	if (len(parts) - 2) == 0 {
		return nil
	}

	signsStr := parts[2:]
	signs := make([]Significance, 0)

	for _, signStr := range signsStr {
		signLenght := len(signStr)

		if success, err := tryParseVarSeter(&signs, signStr, signLenght); err != nil {
			return err
		} else if success || tryParseFlag(&signs, signStr, signLenght) {
			continue
		}

		parseData(&signs, signStr)
	}

	command.Significances = signs
	return nil
}

func tryParseVarSeter(signs *[]Significance , str string , strlen int) (bool , error) {
	if !strings.HasPrefix(str, ".") {
		return false , nil
	}

	if strlen <= 1 {
		return false, fmt.Errorf("Error parsing var seter : invalid format")
	}
	
	str = removeCharsFromBegine(str, 1)
	values := strings.Split(str, "=")
	if len(values) != 2 {
		return false, fmt.Errorf("Error parsing var seter : invalid var seter declaration")
	}
	sign := NewSignificance(Prop, values[0], values[1])
	sign.TProp = getPropType(value[1]) 
	*signs = append(*signs, sign)
	return true, nil
}

func getPropType(value string) PropType {
	for _, char := range value {
		if !unicode.IsDigit() {
			return String
		}
	}

	return Number
}

func tryParseFlag(signs *[]Significance , str string , strlen int) bool {
	if !strings.HasPrefix(str, "--") || strlen <= 2 {
		return false
	}

	str = removeCharsFromBegine(str, 2)
	sign := NewSignificance(Flag, str, "")
	*signs = append(*signs, sign)
	return true
}

func parseData(signs *[]Significance , str string) {
	sign := NewSignificance(Data, "", str)
	*signs = append(*signs, sign)
}

func removeCharsFromBegine(str string, count int) string {
	if count <= 0 {
		return str
	}

	if strlen := len(str); strlen < count || strlen == 0 {
		return ""
	}
	result := str[count:]
	return result
}
