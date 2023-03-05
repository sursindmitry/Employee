package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"prj/internal/app/model/users/controller"
	"prj/internal/app/model/users/service"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.RouterGroup
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting api server")

	routing(s)
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func routing(s *APIServer) {
	server := gin.New()

	server.GET("/users", func(ctx *gin.Context) {
		s.logger.Info("GET request, /users, method FindAll()")
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		s.logger.Info("POST request, /users, method Save(*gin.Context)")
		ctx.JSON(200, userController.Save(ctx))
	})

	server.Run(":8080")
}
