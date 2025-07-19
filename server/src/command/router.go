package command

type VarType int

const (
	Int VarType = iota
	String
	Float
)

type Router struct {
	way   map[string][]EndPoint
	cache map[string]EndPoint
}

type EndPointsGroup struct {
	router *Router
	name   string
}

type EndPoint struct {
	router        *Router
	name          string
	isFlagsEnable bool
	isVarsEnable  bool
	existFlags    []string
	existVars     []string
	handler       func([]Significance) error
}

type Var struct {
}

func NewRouter() *Router {
	return &Router{
		way:   make(map[string][]EndPoint, 0),
		cache: make(map[string]EndPoint, 0),
	}
}

func NewEndPoint(name string, handler func([]Significance) error, router *Router) EndPoint {
	return EndPoint{
		name:    name,
		handler: handler,
		router:  router,
	}
}

func (r *Router) NewEndPointGroup(name string) EndPointsGroup {
	_, exists := r.way[name]

	if !exists {
		r.way[name] = make([]EndPoint, 0)
	}

	return EndPointsGroup{
		name:   name,
		router: r,
	}
}

func (eg EndPointsGroup) NewEndPoint(name string, handler func([]Significance) error) EndPointsGroup {
	_, exist := eg.router.way[eg.name]

	if !exist {
		eg.router.way[eg.name] = make([]EndPoint, 0)
	}

	eg.router.way[eg.name] = append(eg.router.way[eg.name], NewEndPoint(name, handler, eg.router))
	return eg
}

func (eg EndPointsGroup) EndFilling() *Router {
	return eg.router
}
