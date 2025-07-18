package command

type SignificanceType int

const (
	Data	SignificanceType = iota
	Flag
	Prop
)

type Command struct {
	Command string
	Subcommand string 
	Significances []Significance
}

type Significance struct {
	Type SignificanceType
	Content string
	Name string
}

func NewSignificance(type SignificanceType, name string, content string) Significance {
	return Significance {
		Type:type,
		Name:name,
		Content:content,
	}
}

func NewMinimalCommand(command string) *Command {
	return &Command {
		Command:command,
		Subcommand:"",
		Significances:nil,
	}
}
