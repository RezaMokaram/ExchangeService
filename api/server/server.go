package server

import (
	"log"

	"github.com/RezaMokaram/ExchangeService/config"
	_ "github.com/RezaMokaram/ExchangeService/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func NewServer() *echo.Echo {
	return echo.New()
}

func RunServer(e *echo.Echo, db *gorm.DB, cfg *config.APP) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	UserRoutes(e, db)
	PriceRoutes(e, db)
	TradeRoutes(e, db)
	BankRoutes(e, db)
	AdminRoutes(e, db)
	SupportRoutes(e, db)
	log.Fatal(e.Start(cfg.Host + ":" + cfg.Port))
}