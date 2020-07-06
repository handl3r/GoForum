package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {

}

func (server *Server) Run(address string) {
	
}
