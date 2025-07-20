package command

import (
	"fmt"
	"unicode"
)

type Router struct {
	way              map[string][]*EndPoint
	cache            map[string]EndPoint
	errorHandler     func(error)
	endPointSelector func(Command, *EndPoint) bool
}

type EndPointsGroup struct {
	router *Router
	name   string
}

type EndPoint struct {
	group         EndPointsGroup
	name          string
	isFlagsEnable bool
	isVarsEnable  bool
	existFlags    []string
	existVars     []Var
	handler       func(ctx *Context) error
}

type Var struct {
	TProp            DataType
	HaveDefaultValue bool
	Name             string
	DefaultValue     string
	Info             string
}

func NewRouter(errorHandler func(error)) *Router {
	return &Router{
		way:              make(map[string][]*EndPoint, 0),
		cache:            make(map[string]EndPoint, 0),
		errorHandler:     errorHandler,
		endPointSelector: defaultEndPointSelector,
	}
}

func NewEndPoint(name string, handler func(*Context) error, group EndPointsGroup) *EndPoint {
	return &EndPoint{
		name:       name,
		handler:    handler,
		group:      group,
		existFlags: make([]string, 0),
		existVars:  make([]Var, 0),
	}
}

func NewVar(name string, tprop DataType, haveDefaultValue bool, defaultValue string, info string) Var {
	return Var{
		Name:             name,
		TProp:            tprop,
		HaveDefaultValue: haveDefaultValue,
		DefaultValue:     defaultValue,
		Info:             info,
	}
}

func (r *Router) NewEndPointGroup(name string) EndPointsGroup {
	_, exists := r.way[name]

	if !exists {
		r.way[name] = make([]*EndPoint, 0)
	}

	return EndPointsGroup{
		name:   name,
		router: r,
	}
}

func (r *Router) Route(cmd Command) {
	if err := validateCmd(cmd); err != nil {
		r.errorHandler(err)
		return
	}

	group, exist := r.way[cmd.Command]

	if !exist {
		r.errorHandler(fmt.Errorf("Route error : can't find endpoint group with name '%s'", cmd.Command))
		return
	}

	if len(group) == 0 {
		r.errorHandler(fmt.Errorf("Route error : comand gourp '%s' is empty", cmd.Command))
		return
	}

	var targetPoint *EndPoint
	for _, currentPoint := range group {
		if r.endPointSelector(cmd, currentPoint) {
			targetPoint = currentPoint
			break
		}
	}

	if targetPoint == nil {
		r.errorHandler(fmt.Errorf("Routing error : not found command '%s' subcommand '%s'", cmd.Command, cmd.Subcommand))
		return
	}

	ctx := RefinementToContext(cmd.Significances)
	targetPoint.handler(ctx)
}

func defaultEndPointSelector(cmd Command, point *EndPoint) bool {
	if cmd.Subcommand == point.name {
		return true
	}
	return false
}

func validateCmd(cmd Command) error {
	if !hasVisibleChars(cmd.Command) {
		return fmt.Errorf("Invalid cmd : cmd.Command is empty")
	}

	if !hasVisibleChars(cmd.Subcommand) {
		return fmt.Errorf("Invalid cmd : cmd.Subcommand is empty")
	}

	return nil
}

func hasVisibleChars(s string) bool {
	for _, r := range s {
		if unicode.IsPrint(r) && !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
