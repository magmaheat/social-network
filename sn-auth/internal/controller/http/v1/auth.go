package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sn-auth/internal/service"
)

type authRoutes struct {
	authService service.Auth
}

func newAuthRoutes(g *echo.Group, authService service.Auth) {
	r := &authRoutes{
		authService: authService,
	}

	g.POST("/sign-up", r.signUp)
	g.POST("/sign-up", r.signIn)
}

type signUpInput struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,password"`
}

func (r *authRoutes) signUp(c echo.Context) error {
	var input signUpInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := r.authService.CreateUser(c.Request().Context(), service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Id int `json:"id"`
	}

	return c.JSON(http.StatusCreated, response{
		Id: id,
	})

}

type signInInput struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,password"`
}

func (r *authRoutes) signIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := r.authService.GenerateToken(c.Request().Context(), service.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if err == service.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, "invalid username or password")
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return err
	}

	type response struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusOK, response{
		Token: token,
	})
}
