package controller

import (
	"errors"
	"log"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/helpers"
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}
	signResponse, signInData, err := ac.authService.SignIn(signInBody)
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

	helpers.CreateSignCookies(ctx, signResponse, signInBody.KeepSignedIn)

	return ctx.Status(fiber.StatusOK).JSON(signInData)
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Password reset successfully!",
	})
}

func (ac *authController) SignOut(ctx *fiber.Ctx) error {
	signOutBody := body.SignOut{
		AccessToken:  ctx.Cookies("accessToken", ""),
		RefreshToken: ctx.Cookies("refreshToken", ""),
	}

	if err := ac.authService.SignOut(signOutBody); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	helpers.ClearCookies(ctx)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

func (ac *authController) Session(ctx *fiber.Ctx) error {
	userId := int64(ctx.Locals("userId").(float64))
	if userId <= 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": utils.ErrNoDataFound,
		})
	}

	userSessionData, err := ac.authService.GetUserSessionData(userId)
	if err != nil {
		log.Println("err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.JSON(userSessionData)
}

func (ac *authController) CheckIfEmailIsAvailable(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidParameter.Error(),
		})
	}

	available, err := ac.authService.CheckIfEmailIsAvailable(email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"available": available,
	})
}

func (ac *authController) SignInWithGoogle(ctx *fiber.Ctx) error {
	var signInBody body.GoogleSignInBody

	err := ctx.BodyParser(&signInBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	signResponse, signInData, err := ac.authService.SignInWithGoogle(&signInBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	helpers.CreateSignCookies(ctx, signResponse, signInBody.KeepSignedIn)

	return ctx.Status(fiber.StatusOK).JSON(signInData)
}

func (ac *authController) SignUpWithGoogle(ctx *fiber.Ctx) error {
	var signUpBody body.GoogleSignUpBody

	err := ctx.BodyParser(&signUpBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	signResponse, signInData, err := ac.authService.SignUpWithGoogle(&signUpBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	helpers.CreateSignCookies(ctx, signResponse, signUpBody.KeepSignedIn)

	return ctx.Status(fiber.StatusOK).JSON(signInData)
}
