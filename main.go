package main

import (
  "fmt"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  _, err := sqlConnect()
  if err != nil {
      panic(err.Error())
  } else {
      fmt.Println("DB接続成功")
  }

  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(ctx *gin.Context){
      ctx.HTML(200, "index.html", gin.H{})
  })

  router.Run()
}

func sqlConnect() (database *gorm.DB, err error) {
  DBMS := "mysql"
  USER := "go_test"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
  return gorm.Open(DBMS, CONNECT)
}