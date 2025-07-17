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

