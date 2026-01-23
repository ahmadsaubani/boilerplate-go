package routerGroup

type apiGroup struct {
	HealtCheck iHealthCheckGroup
}

var RoutesGroupCollection = apiGroup{
	HealtCheck: newHealthCheckGroup(),
}
