package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/service"
	"net/http"
)

type authRouter struct {
	authService service.Auth
}

func NewAuthRouter(g *echo.Group, authService service.Auth) {
	r := &authRouter{
		authService: authService,
	}

	g.POST("/sign-up", r.signUp)
	g.POST("/sign-in", r.signIn)
}

type signUpInput struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,password"`
}

func (a *authRouter) signUp(c echo.Context) error {
	var input signUpInput

	if err := c.Bind(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := a.authService.CreateUser(c.Request().Context(), service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		if err == service.ErrUserAlreadyExists {
			NewErrorResponce(c, http.StatusBadRequest, err.Error())
			return err
		}
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
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

func (a *authRouter) signIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return err
	}

	if err := c.Validate(input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	token, err := a.authService.GenerateToken(c.Request().Context(), service.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if err == service.ErrUserNotFound {
			NewErrorResponce(c, http.StatusBadRequest, "invalid username or password")
			return err
		}

		NewErrorResponce(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusOK, response{
		Token: token,
	})
}
