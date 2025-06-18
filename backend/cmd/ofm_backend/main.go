package main

import (
	"log"
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	chat_routes "ofm_backend/cmd/ofm_backend/api/chat/routes"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	home_data_routes "ofm_backend/cmd/ofm_backend/api/home_data/routes"
	my_profile_routes "ofm_backend/cmd/ofm_backend/api/my_profile/routes"
	order_routes "ofm_backend/cmd/ofm_backend/api/order/routes"
	payment_routes "ofm_backend/cmd/ofm_backend/api/payment/routes"
	search_routes "ofm_backend/cmd/ofm_backend/api/search/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/rsa_encryption"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvValues()

	s3Client := config.InitS3Client()
	rsa_encryption.TryToLoadRSAKeys()
	posgresqlDb := database.ConnectToPostgresDB()
	redisDb, err := database.ConnectToRedisDB()
	if err != nil {
		log.Fatal("redisDb err:", err)
	}
	defer func() {
		posgresqlDb.Close()
		redisDb.Close()
	}()

	app := fiber.New()
	app.Use(config.ConfigCors())
	app.Use(config.ConfigRateLimiter())

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup, posgresqlDb, redisDb, s3Client)
	freelance_routes.RegisterFreelanceRoutes(apiGroup, posgresqlDb)
	user_routes.RegisterUserRoutes(apiGroup, posgresqlDb)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	search_routes.RegisterSearchRoutes(apiGroup)
	home_data_routes.RegisterHomeDataRoutes(apiGroup, posgresqlDb)
	payment_routes.RegisterPaymentRoutes(apiGroup, posgresqlDb)
	order_routes.RegisterOrderRoutes(apiGroup, posgresqlDb, s3Client)
	my_profile_routes.RegisterMyProfileRoutes(apiGroup, posgresqlDb, s3Client)
	chat_routes.RegisterChatRoutes(apiGroup, posgresqlDb)

	app.Listen(":8000")
}
