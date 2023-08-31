package main

import (
	db "realtime-chat/db"
	middlewares "realtime-chat/middlewares"
	handler "realtime-chat/modules/user/handlers"
	ws "realtime-chat/modules/ws/handlres"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middlewares.SetJSONContentTypeMiddleware())

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	// Middleware untuk mengaktifkan CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Ganti dengan alamat halaman Anda
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	handler.InitUserHttpHandler(router)
	ws.InitUserHttpHandler(router, wsHandler)

	router.Run("localhost:8080")
}
