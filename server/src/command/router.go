package command

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

func (eg EndPointsGroup) AddEndPoint(name string, handler func([]Significance) error) *EndPoint {
	_, exist := eg.router.way[eg.name]

	if !exist {
		eg.router.way[eg.name] = make([]*EndPoint, 0)
	}
	newEndPoint := NewEndPoint(name, handler, eg)
	eg.router.way[eg.name] = append(eg.router.way[eg.name], newEndPoint)
	return newEndPoint
}

func (eg *EndPoint) EndFilling() *EndPointsGroup {
	return &eg.group
}

func (eg EndPointsGroup) EndFilling() *Router {
	return eg.router
}

func (ep *EndPoint) NewName(newName string) *EndPoint {
	ep.name = newName
	return ep
}

func (ep *EndPoint) SetFlagsEnable(isEnable bool) *EndPoint {
	ep.isFlagsEnable = isEnable
	return ep
}

func (ep *EndPoint) SetVarsEnable(isEnable bool) *EndPoint {
	ep.isVarsEnable = isEnable
	return ep
}

func (ep *EndPoint) AddFlag(flag string) *EndPoint {
	ep.existFlags = append(ep.existFlags, flag)
	return ep
}

func (ep *EndPoint) AddVar(name string, tprop PropType, defaultValue string, info string) *EndPoint {
	ep.existVars = append(ep.existVars, NewVar(name, tprop, defaultValue != "", defaultValue, info))
	return ep
}
