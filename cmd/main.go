package main

import (
	"test/pkg/common/db"

	"test/pkg/books"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	r := gin.Default()
	h := db.Init(dbUrl)
	books.RegisRouters(r, h)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})
	r.Run(port)
}
