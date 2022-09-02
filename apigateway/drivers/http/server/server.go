package server

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *gin.Engine
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	broker := &Broker{
		config: config,
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	b.router = gin.Default()
	b.router.SetTrustedProxies([]string{"127.0.0.1"})
	binder(b, b.router)
	log.Println("starting server on port", b.config.Port)
	if err := b.router.Run(b.config.Port); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
