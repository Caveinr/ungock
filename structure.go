package main

import (
	"github.com/gin-gonic/gin"
	"ungock/CQHTTP"
	"ungock/GORM"
	"ungock/MFRCScan"
)

type Server struct {
	r    *gin.Engine
	port string
	db   *GORM.DB
}

type Config struct {
	Port   string            `toml:"PORT"`
	Cqhttp CQHTTP.CQHTTP     `toml:"CQHTTP"`
	Mfrcs  MFRCScan.MFRCScan `toml:"MFRCScan"`
	DB     GORM.DB           `toml:"GORM"`
}
