package router

func routeStatic() {
	router.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")
	router.Static("/assets", "./frontend/dist/assets")
	router.StaticFile("/", "./frontend/dist/index.html")
}
