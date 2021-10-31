package route

func (r *Route) initUser() {
	user := r.Echo.Group("/api/user")
	user.Use(r.controller.Auth.Process)

	user.POST("/register", r.controller.RegisterUser)
	user.POST("/login", r.controller.UserLogin)

	user.POST("/orders", r.controller.Orders)
	user.GET("/orders", r.controller.OrdersList)
}
