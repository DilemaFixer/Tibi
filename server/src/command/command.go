package command

type SignificanceType int
type DataType int

const (
	Data SignificanceType = iota
	Flag
	Prop
)

const (
	None DataType = iota
	String
	Number
)

type Command struct {
	Command       string
	Subcommand    string
	Significances []Significance
}

type Significance struct {
	Type    SignificanceType
	TProp   DataType
	Content string
	Name    string
}

func NewSignificance(_type SignificanceType, name string, content string) Significance {
	return Significance{
		Type:    _type,
		Name:    name,
		Content: content,
	}
}

func NewMinimalCommand(command string) *Command {
	return &Command{
		Command:       command,
		Subcommand:    "",
		Significances: nil,
	}
}
