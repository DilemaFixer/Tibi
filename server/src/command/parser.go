package command


import (
	"strings"
	"fmt"
)
func ParseToCommand(str string) (*Command , error) {
	if str == "" || str == " " {
		return nil ,fmt.Errorf("Can't parse empty string to command struct")
	}

	parts := strings.Split(str , " ")
	if len(parts) <= 1 {
		return nil , fmt.Errorf("Error parsging command , invalid format")
	}

	var command , subcommand string 

	command = parts[0]
	subcommand = parts[1]
	commandObj := &Command {
		Command:command,
		Subcommand:subcommand,
		Significances:nil,
	}

	if len(parts) > 2 {
		parts = parts[2:]
		signs , err := parseSignificance(parts)

		if err != nil {
			return nil , err
		}

		commandObj.Significances = signs
	}
	fmt.Println(commandObj.Significances)
	return commandObj, nil
}

func parseSignificance(parts []string) ([]Significance , error){
	signs := make([]Significance , 1)
	for _ ,v := range parts {
		if strings.HasPrefix(v , ".") {
			v = trimRunes(v, 0)
			varSeterParts := strings.Split(v , "=")
			if len(varSeterParts) != 2 {
				return nil , fmt.Errorf("Invalid var seter declaration :", v)
			}
			sign := Significance { 
				Type:Prop, 
				Name:varSeterParts[0], 
				Content:varSeterParts[1],
			}
			signs = append(signs, sign)
			continue
		}
		
		if strings.HasPrefix(v , "--") {
			v = trimRunes(v, 2)
			sign := Significance {
				Type:Flag,
				Name:v,
				Content:"",
			}
			signs = append(signs, sign)
			continue
		}

		sign := Significance {
			Type:Data,
			Name:"",
			Content:v,
		}
		signs = append(signs, sign)
	}
	return signs, nil
}

func trimRunes(s string , count int) string {
	if count < 0 {
		return s
	}

    for i := range s {
        if i > count {
            return s[i:]
        }
    }
    return ""
}
