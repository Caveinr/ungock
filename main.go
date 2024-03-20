package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"ungock/CQHTTP"
	"ungock/GORM"
	"ungock/MFRCScan"
	//"gorm.io/gorm"
)

func main() {
	s := New()
	s.Run()
}

func loadConfig(path string) Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		panic(err)
	}
	return config
}

func New() *Server {
	var s Server
	var config = loadConfig("config.toml")
	s.port = config.Port

	db := GORM.Engine(config.DB)

	s.r = gin.Default()
	s.r.POST("/MFRCScan", MFRCScan.Auth(config.Mfrcs), MFRCScan.Handle(config.Mfrcs, db))
	s.r.POST("/cqhttp", CQHTTP.Auth(config.Cqhttp), CQHTTP.Handle(config.Cqhttp, db))

	return &s
}

func (s *Server) Run() {
	s.r.Run(s.port)
}

func printResp(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))
	c.JSON(http.StatusOK, gin.H{})
}
