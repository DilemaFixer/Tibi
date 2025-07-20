package command

type VarNotFoundError struct {
	VarName string
}

func (e *VarNotFoundError) Error() string {
	return "variable '" + e.VarName + "' not found"
}

func NewVarNotFoundError(name string) *VarNotFoundError {
	return &VarNotFoundError{VarName: name}
}

type ParseError struct {
	VarName    string
	TargetType string
	Value      string
	Err        error
}

func (e *ParseError) Error() string {
	return "failed to parse variable '" + e.VarName + "' with value '" + e.Value + "' as " + e.TargetType + ": " + e.Err.Error()
}

func NewParseError(varName, targetType, value string, err error) *ParseError {
	return &ParseError{
		VarName:    varName,
		TargetType: targetType,
		Value:      value,
		Err:        err,
	}
}
