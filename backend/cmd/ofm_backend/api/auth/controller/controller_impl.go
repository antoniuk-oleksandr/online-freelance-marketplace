package controller

import (
	"errors"
	"log"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (authController *authController) ConfirmEmail(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid").(string)

	err := authController.authService.ConfirmEmail(uuid)
	if err != nil {
		if errors.Is(err, utils.ErrTempTokenExpired) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "The token has expired",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Email confirmed",
	})
}

func (ac *authController) ForgotPassword(ctx *fiber.Ctx) error {
	var forgotPasswordBody body.ForgotPassword

	if err := ctx.BodyParser(&forgotPasswordBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	if err := ac.authService.ForgotPassword(forgotPasswordBody.Email); err != nil {
		if err == utils.ErrUsernameDoesNotExist {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ErrUsernameDoesNotExist.Error(),
			})
		}

		if err == utils.ErrEmailDoesNotExist {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ErrEmailDoesNotExist.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The email was sent successfully",
	})
}

func (ac *authController) SignUp(ctx *fiber.Ctx) error {
	var user body.SignUpBody

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := ac.authService.SignUp(&user)
	if err != nil {
		if utils.ErrMailSend.Error() == err.Error() {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid email address",
			})
		}

		if errors.Is(err, utils.ErrEmailIsTaken) || errors.Is(err, utils.ErrUsernameIsTaken) {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create account",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The email was sent successfully",
	})
}

func (ac *authController) SignIn(ctx *fiber.Ctx) error {
	var signInBody body.SignInBody
	if err := ctx.BodyParser(&signInBody); err != nil {
		return err
	}
	accessToken, refreshToken, err := ac.authService.SignIn(signInBody)

	if err != nil {
		if err.Error() == fiber.ErrUnprocessableEntity.Error() {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
				fiber.Map{"error": "Invalid request body"},
			)
		}

		if err.Error() == fiber.ErrUnauthorized.Error() {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid credentials"},
			)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "An unexpected error occurred"},
		)
	}

	return ctx.JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (ac *authController) GoogleAuth(ctx *fiber.Ctx) error {
	var googleBody body.Google

	if err := ctx.BodyParser(&googleBody); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	if googleBody.Code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	accessToken, refreshToken, err := ac.authService.GoogleAuth(googleBody.Code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (ac *authController) RefreshToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong token",
		})
	}

	accessToken, err := ac.authService.RefreshToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"accessToken": accessToken,
	})
}

func (ac *authController) ResetPassword(ctx *fiber.Ctx) error {
	var resetPasswordBody body.ResetPassword
	if err := ctx.BodyParser(&resetPasswordBody); err != nil {
		return ctx.JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	email := ctx.Locals("email").(string)
	token := ctx.Locals("token").(string)

	err := ac.authService.ResetPassword(resetPasswordBody, email, token)
	if err != nil {
		log.Println("err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Password reset successfully!",
	})
}

func (ar *authController) SignOut(ctx *fiber.Ctx) error {
	var signOutBody body.SignOut

	if err := ctx.BodyParser(&signOutBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	if err := ar.authService.SignOut(signOutBody); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
