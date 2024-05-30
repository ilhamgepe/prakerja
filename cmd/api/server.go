package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	*gorm.DB
	srv *http.Server
}

func NewServer(addr string, db *gorm.DB) Server {
	r := gin.Default()
	setupRoutes(r, db)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", addr),
		Handler:      r.Handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &server{
		DB:  db,
		srv: srv,
	}
}

func (s *server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
