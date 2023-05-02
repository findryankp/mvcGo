package middlewares

func InitMiddlewares() {
	AuthenticationCreate()
	LogCreate()
	CorsCreate()
	JwtCreate()
}
