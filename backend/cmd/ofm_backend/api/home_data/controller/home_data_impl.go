package controller

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type homeController struct {
	homeService service.HomeService
}

func NewHomeController(
	homeService service.HomeService,
) HomeController {
	return &homeController{
		homeService: homeService,
	}
}

func (hc *homeController) GetHomeData(ctx *fiber.Ctx) error {
	homeData, err := hc.homeService.GetHomeData()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(homeData)
}
