package command

func (eg EndPointsGroup) AddEndPoint(name string, handler func(*Context) error) *EndPoint {
	value, exist := eg.router.way[eg.name]

	if !exist {
		value = make([]*EndPoint, 0)
	}
	newEndPoint := NewEndPoint(name, handler, eg)
	eg.router.way[eg.name] = append(value, newEndPoint)
	return newEndPoint
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

func (ep *EndPoint) AddVar(name string, tprop DataType, defaultValue string, info string) *EndPoint {
	ep.existVars = append(ep.existVars, NewVar(name, tprop, defaultValue != "", defaultValue, info))
	return ep
}

func (eg *EndPoint) EndFilling() *EndPointsGroup {
	return &eg.group
}
