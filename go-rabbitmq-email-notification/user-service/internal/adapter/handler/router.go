package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/domain"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	conf *config.HTTP,
	jwtConf *config.JWT,
	userHandler *UserHandler,
	authHandler *AuthHandler,
) *Router {
	r := gin.New()

	pb := r.Group("/api/v1")
	us := pb.Group("/", AuthMiddleware(jwtConf))
	ad := us.Group("/", RoleMiddleware(domain.AdminRole))

	pb.POST("/login", authHandler.Login)
	pb.POST("/register", userHandler.RegisterUser)
	pb.GET("/confirm-email", userHandler.ConfirmEmail)

	ad.GET("/users", userHandler.GetUsers)

	return &Router{r}
}

func (r *Router) Serve(conf *config.HTTP) error {
	return r.r.Run(conf.Host + ":" + conf.Port)
}
