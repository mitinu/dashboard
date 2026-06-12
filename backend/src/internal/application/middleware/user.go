package middleware

import (
	"backend/src/internal/config"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID             int64
	AccessStatusId int64
	jwt.RegisteredClaims
}

func UserMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Получаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Заголовок Authorization отсутствует"})
			return
		}

		// 2. Проверяем формат заголовка. Ожидается: "Bearer <токен>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат заголовка. Используйте 'Bearer <token>'"})
			return
		}

		tokenString := parts[1]

		// 3. Проверяем токен с использованием секрета из конфигурации (cfg.Pepper)
		secretKey := []byte(cfg.Pepper)
		claims, err := verifyJwtToken(tokenString, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Недействительный или истекший токен"})
			return
		}

		// 4. Записываем userID в контекст Gin, чтобы хэндлеры могли получить доступ к нему
		c.Set("ID", claims.ID)
		c.Set("AccessStatusId", claims.AccessStatusId)

		// Передаем запрос следующему обработчику в цепочке
		c.Next()
	}
}

func verifyJwtToken(tokenString string, secretKey []byte) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Извлекаем claims и проверяем, валиден ли токен
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
