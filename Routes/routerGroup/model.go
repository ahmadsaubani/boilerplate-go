package routerGroup

type apiGroup struct {
	HealtCheck iHealthCheckGroup
	Example    iExampleGroup
}

var RoutesGroupCollection = apiGroup{
	HealtCheck: newHealthCheckGroup(),
	Example:    newExampleGroup(),
}
