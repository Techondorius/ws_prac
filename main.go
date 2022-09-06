package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	// DBマイグレーション
	// model.Connectionがエラー発生しなくなるまで=DBが立ち上がるまで待機
	// (docker composeで立ち上げると必ずdbのほうが立ち上がり遅い)

	//_, dbConErr := model.Connection()
	//for dbConErr != nil {
	//	time.Sleep(time.Second)
	//	_, dbConErr = model.Connection()
	//}
	//migration.Mig()

	r := gin.Default()

	r.Use(logger())

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// ルーティング
	//routing.Routing(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})

	_ = r.Run()
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ByteBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
		log.Println("Endpoint: " + c.FullPath())
		log.Println("Body: " + string(ByteBody))

		q := c.Request.URL.Query()
		j, _ := json.Marshal(q)
		log.Println("Query Params: " + string(j))

		c.Next()
	}
}
