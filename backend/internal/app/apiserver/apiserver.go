package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.RouterGroup
}

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
	starting()
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

func starting() {
	server := gin.Default()
	server.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})
	server.Run(":8080")
}
