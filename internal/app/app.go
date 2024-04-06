package app

import (
	config "app/configs"
	"app/store"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Config *config.Config
var Log *logrus.Logger
var Store *store.Store

type APIServer struct {
	config *config.Config
	router *gin.Engine
	logger *logrus.Logger
	store  *store.Store
}

func New(router *gin.Engine) *APIServer {
	Config = config.Cfg
	Log = logrus.New()
	return &APIServer{
		config: Config,
		router: router,
		logger: Log,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		Log.Fatal("error configureLogger", err)
		return err
	}

	Log.Info("starting api server at port :", s.config.Port)

	if err := s.configureStore(); err != nil {
		Log.Fatal("error configureStore", err)
		return err
	}
	Log.Info("postgres connected!")

	if err := s.router.Run(); err != nil {
		Log.Fatal("error Run", err)
		return err
	}

	return nil
}

func (s *APIServer) configureStore() error {
	Store = store.New()
	err := Store.Open()
	if err != nil {
		return err
	}
	s.store = Store
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LoggerLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
