package command

import (
	"fmt"
	"unicode"
)

type Router struct {
	way          map[string][]*EndPoint
	cache        map[string]EndPoint
	errorHandler func(error)
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
	handler       func([]Significance) error
}

type Var struct {
	TProp            PropType
	HaveDefaultValue bool
	Name             string
	DefaultValue     string
	Info             string
}

func NewRouter(errorHandler func(error)) *Router {
	return &Router{
		way:          make(map[string][]*EndPoint, 0),
		cache:        make(map[string]EndPoint, 0),
		errorHandler: errorHandler,
	}
}

func NewEndPoint(name string, handler func([]Significance) error, group EndPointsGroup) *EndPoint {
	return &EndPoint{
		name:       name,
		handler:    handler,
		group:      group,
		existFlags: make([]string, 0),
		existVars:  make([]Var, 0),
	}
}

func NewVar(name string, tprop PropType, haveDefaultValue bool, defaultValue string, info string) Var {
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
		r.errorHandler(fmt.Errorf("Route error : can't find endpoint group with name %s", cmd.Command))
		return
	}

	var point *EndPoint

	for _, currentPoint := range group {
		if currentPoint.name == cmd.Subcommand {
			point = currentPoint
			break
		}
	}

	point.handler(cmd.Significances)
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
