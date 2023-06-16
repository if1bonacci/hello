package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
	"github.com/if1bonacci/lets-go-chat/pkg/tokenGenerator"
	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

type LoginResponse struct {
	Url string `json:"url"`
}

const ChatLink = "/websoket?token="

type AuthHandler struct {
	repo     repositories.MessageRepository
	userRepo *repositories.UserRepository
}

func ProvideAuthHandler(r repositories.MessageRepository, u *repositories.UserRepository) AuthHandler {
	return AuthHandler{
		repo:     r,
		userRepo: u,
	}
}

// Register godoc
// @Summary      Register
// @Description  Register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body AuthRequest true "auth body"
// @Success      200  {object}  RegisterResponse
// @Router       /user [post]
func (h *AuthHandler) Register(ctx echo.Context) (err error) {
	request := new(AuthRequest)

	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := h.userRepo.CreateUser(request.UserName, request.Password)
	h.userRepo.StoreUser(user)

	resp := &RegisterResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// Login godoc
// @Summary      Login
// @Description  Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body AuthRequest true "auth body"
// @Success      200  {object}  LoginResponse "success"
// @Router       /user/login [post]
func (h *AuthHandler) Login(ctx echo.Context) (err error) {
	u := new(AuthRequest)
	path := os.Getenv("URL") + ChatLink

	if err = ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := h.userRepo.GetUserByName(u.UserName)

	if !hasher.CheckPasswordHash(u.Password, user.Password) || u.UserName != user.UserName {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username/password")
	}

	ctx.Response().Header().Set("X-Rate-Limit", "3000")
	ctx.Response().Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())
	token := tokenGenerator.Generate()
	user.Token = token

	return ctx.JSON(http.StatusOK, LoginResponse{path + token})
}
