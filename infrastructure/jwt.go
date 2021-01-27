package infrastructure

import (
	"fmt"
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	req "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

type userClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

var signingKey = []byte("SecretKeyThatShouldBeLongEnough")

// GenerateToken returns a JSON response containing a new JWT
func GenerateToken(c *gin.Context) {
	claims := userClaims{
		"owner",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer: "api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "token": ss})
}

// VerifyToken verifies recieved token
func VerifyToken(c *gin.Context) {
	tokenString, err := req.HeaderExtractor{"Token"}.ExtractToken(c.Request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(403)
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		c.AbortWithStatus(403)
		return
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		log.Printf("%v %v %v", claims.Role, claims.StandardClaims.ExpiresAt, claims.StandardClaims.Issuer)
	} else {
		log.Println("L74")
		c.AbortWithStatus(403)
		return
	}
}
