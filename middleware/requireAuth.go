package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jm61/jwt/initializers"
	"github.com/jm61/jwt/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("COUCOU C'EST LE MIDDLEWARE")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("401 Unauthorized")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	//fmt.Println("token: ", token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("31 401 Unauthorized")
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("38 401 Unauthorized")
		}
		c.Set("user", user)
		fmt.Println("user: ", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("45 401 Unauthorized")
	}
}

/* func RequireAuth(c *gin.Context) {
	fmt.Println("COUCOU C'EST LE MIDDLEWARE")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("18 401 Unauthorized")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["sub"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("31 401 Unauthorized")
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("38 401 Unauthorized")
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("45 401 Unauthorized")
	}
} */
