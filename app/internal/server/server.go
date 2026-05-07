package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	addr   string
}

func NewServer(addr string, handler Handler) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	handler.RegisterRoutes(router)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return &Server{
		server: server,
		addr:   addr,
	}
}

func (s *Server) Start() error {
	go func() {
		fmt.Printf("server has been started on %v\n", s.addr)
		if err := s.server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
