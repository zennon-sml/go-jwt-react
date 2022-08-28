package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zennon-sml/go-jwt-react/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/zennon-sml/go-jwt-react/"
)

var DB *gorm.DB

func MKCon() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro carregando o arquivo .env")
	}
	//https://www.phpmyadmin.co/ link pro php
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	banco := os.Getenv("DB_NAME")

	dominio := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usuario, senha, host, porta, banco)

	bancoDeDados, err := gorm.Open(mysql.Open(dominio), &gorm.Config{})
	if err != nil {
		log.Fatal("não foi possivel estabelecer conexão ", err)
	}
	DB = bancoDeDados
	//MIGRATIONS
	DB.Debug().AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return DB
}
