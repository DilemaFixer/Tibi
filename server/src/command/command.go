package command

type SignificanceType int
type PropType int

const (
	Data SignificanceType = iota
	Flag
	Prop
)

const (
	Number PropType = iota
	String
)

type Command struct {
	Command       string
	Subcommand    string
	Significances []Significance
}

type Significance struct {
	Type    SignificanceType
	TProp   PropType
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
