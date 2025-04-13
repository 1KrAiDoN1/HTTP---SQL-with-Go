package main

import (
	"helloapp/internal/handler"
)

//проверка

// @title Crypto Market API
// @version 1.0
// @description API server for Crypto Market

// @contact.url https://t.me/KrAiDoN
// @contact.email pavelvasilev24843@gmail.com

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name access_token
// @description JWT token in cookie

// @securityDefinitions.apikey RefreshToken
// @in cookie
// @name refresh_token
// @description Refresh token in cookie
func main() {

	handler.HandleFunc()

}
