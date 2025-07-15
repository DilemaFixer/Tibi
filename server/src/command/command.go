package command

type SignificanceType int

const (
	Data	SignificanceType  = iot
	Flags
)

type Command struct {
	Command string
	Subcommand string 
	Significances []Significance
}

type Significance struct {
	Type SignificanceType
	Content string
}

