package command

import "strconv"

type InputData struct {
	Type  DataType
	Value string
}

type InputVar struct {
	Name  string
	Value string
	Type  DataType
}

type Context struct {
	flags map[string]struct{}
	vars  map[string]InputVar
	datas []InputData
}

func NewContext() *Context {
	return &Context{
		flags: make(map[string]struct{}),
		vars:  make(map[string]InputVar),
		datas: make([]InputData, 0),
	}
}

func NewInputData(t DataType, v string) InputData {
	return InputData{
		Type:  t,
		Value: v,
	}
}

func (ctx *Context) addDataToContext(_type DataType, value string) {
	data := InputData{
		Type:  _type,
		Value: value,
	}
	ctx.datas = append(ctx.datas, data)
}

func (ctx *Context) addFlagToContexts(flag string) {
	ctx.flags[flag] = struct{}{}
}

func (ctx *Context) addVarToContext(name string, value string, _type DataType) {
	_var := InputVar{
		Name:  name,
		Value: value,
		Type:  _type,
	}
	ctx.vars[name] = _var
}

func (ctx *Context) IsFlagExist(flag string) bool {
	_, exist := ctx.flags[flag]
	return exist
}

func (ctx *Context) IsVarExist(name string) bool {
	_, exist := ctx.vars[name]
	return exist
}

func (ctx *Context) GetVarValueAsString(name string) string {
	value, exist := ctx.vars[name]

	if !exist {
		return ""
	}

	return value.Value
}

func (ctx *Context) GetVarValueAsInt(name string) int {
	value, exist := ctx.vars[name]

	if !exist {
		return 0
	}

	intValue, err := strconv.Atoi(value.Value)
	if err != nil {
		return 0
	}

	return intValue
}

func (ctx *Context) GetVarValueAsInt64(name string) int64 {
	value, exist := ctx.vars[name]

	if !exist {
		return 0
	}

	intValue, err := strconv.ParseInt(value.Value, 10, 64)
	if err != nil {
		return 0
	}

	return intValue
}

func (ctx *Context) GetVarValueAsFloat(name string) float64 {
	value, exist := ctx.vars[name]

	if !exist {
		return 0.0
	}

	floatValue, err := strconv.ParseFloat(value.Value, 64)
	if err != nil {
		return 0.0
	}

	return floatValue
}

func (ctx *Context) GetVarValueAsFloat32(name string) float32 {
	value, exist := ctx.vars[name]

	if !exist {
		return 0.0
	}

	floatValue, err := strconv.ParseFloat(value.Value, 32)
	if err != nil {
		return 0.0
	}

	return float32(floatValue)
}

func (ctx *Context) GetVarValueAsBool(name string) bool {
	value, exist := ctx.vars[name]

	if !exist {
		return false
	}

	boolValue, err := strconv.ParseBool(value.Value)
	if err != nil {
		return false
	}

	return boolValue
}

func (ctx *Context) GetVarValueAsIntWithError(name string) (int, error) {
	value, exist := ctx.vars[name]

	if !exist {
		return 0, NewVarNotFoundError(name)
	}

	intValue, err := strconv.Atoi(value.Value)
	if err != nil {
		return 0, NewParseError(name, "int", value.Value, err)
	}

	return intValue, nil
}

func (ctx *Context) GetVarValueAsFloatWithError(name string) (float64, error) {
	value, exist := ctx.vars[name]

	if !exist {
		return 0.0, NewVarNotFoundError(name)
	}

	floatValue, err := strconv.ParseFloat(value.Value, 64)
	if err != nil {
		return 0.0, NewParseError(name, "float64", value.Value, err)
	}

	return floatValue, nil
}

func (ctx *Context) GetVarType(name string) DataType {
	value, exist := ctx.vars[name]

	if !exist {
		return None
	}

	return value.Type
}

func RefinementToContext(signs []Significance) *Context {
	ctx := NewContext()
	if len(signs) == 0 {
		return ctx
	}

	for _, sign := range signs {
		switch sign.Type {
		case Data:
			ctx.addDataToContext(sign.TProp, sign.Content)
		case Flag:
			ctx.addFlagToContexts(sign.Name)
		case Prop:
			ctx.addVarToContext(sign.Name, sign.Content, sign.TProp)
		}
	}
	return ctx
}
