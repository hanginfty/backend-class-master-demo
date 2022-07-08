package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hanginfty/simple-bank/db/sqlc"
)

type Server struct {
	store db.Store
	r     *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := Server{store: store}
	r := gin.Default()

	r.POST("/accounts", server.createAccount)
	r.GET("/accounts/:id", server.getAccount)

	server.r = r
	return &server
}

func (s *Server) Start(addr string) error {
	err := s.r.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
