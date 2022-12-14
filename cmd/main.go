package main

import (
	"simple_crud/pkg/books"
	"simple_crud/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// add env variable as needed

	port := viper.Get("PORT").(string)
	dburl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dburl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dburl": dburl,
		})
	})

	router.Run(port)
}
