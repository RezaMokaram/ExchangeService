package http

import (
	"fmt"

	"github.com/RezaMokaram/ExchangeService/app"
	"github.com/RezaMokaram/ExchangeService/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userSvcGetter := userServiceGetter(appContainer, cfg)
	router.Post("/sign-up", setTransaction(appContainer.DB()), SignUp(userSvcGetter))
	router.Get("/send-otp", setTransaction(appContainer.DB()), SendSignInOTP(userSvcGetter))
	router.Post("/sign-in", setTransaction(appContainer.DB()), SignIn(userSvcGetter))
	router.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)
}
