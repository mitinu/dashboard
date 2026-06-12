package handler

import (
	"backend/src/internal/application"
	"backend/src/internal/config"
	"backend/src/internal/domain"
	"backend/src/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	dUser    domain.User
	dSession domain.Session
	service  application.ReqAndAutoService
}

func (h *UserHandler) Login(cfg *config.Config, login string, Password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := h.dUser.GetByLogin(login)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}
		passwordService := application.NewPasswordService(cfg)
		result, err := passwordService.Verify(Password, user.PasswordHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			return
		}
		if !result {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}
		RefreshToken, err := h.service.CreateRefreshToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			logger.Error.Println(err)
			return
		}
		err = h.dSession.Create(user.ID, RefreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			logger.Error.Println(err)
			return
		}
		AccessToken, err := h.service.CreateJwtToken(user.ID, user.AccessStatusId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			logger.Error.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"refreshToken": RefreshToken,
			"accessToken":  AccessToken,
			"type":         "Bearer",
		})
	}

}
