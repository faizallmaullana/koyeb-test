package Authentication

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	StaffID  string `json:"staff_id"`
	Role     string `json:"role"`
}

var secretKey = []byte("G%3jyF83%?9Bg7,uX;(g-}tug:0n!IFPeJ{7qK!s>@_MGmbl!t/Y7:/j+T|hIZR")

func GenerateToken(username string, staff_id string, role string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
		Username: username,
		StaffID:  staff_id,
		Role:     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWTAuthMiddleware is the middleware to authorize JWT tokens
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		tokenString := authHeader[7:] // Remove the "Bearer " prefix

		// Parse the token
		claims := &MyClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token is valid, proceed with the request
		c.Next()
	}
}

func JWTClaims(tokenString, role string) (gin.H, error) {
	// Parse the JWT token.
	parsedToken, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return gin.H{
			"status": "Unauthorized",
		}, err
	}

	// Check if the token is valid.
	if claims, ok := parsedToken.Claims.(*MyClaims); ok && parsedToken.Valid {

		// role (guru, admin, all)
		if claims.Role != role && role != "all" {
			return gin.H{
				"status": "Unauthorized",
			}, err
		}

		// Return user profile data.
		return gin.H{
			"username": claims.Username,
			"staff_id": claims.StaffID,
			"role":     claims.Role,
			"status":   "Authorized",
		}, nil
	}
	return gin.H{
		"status": "Unauthorized",
	}, err

	// usage

	// tokenString := c.GetHeader("Authorization")
	// profileData, err := jwt_auth.JWTClaims(tokenString)

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
	// 	return
	// }
}

// func JWTHandler(tokenString string) (gin.H, error) {
// 	// Parse the JWT token.
// 	parsedToken, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check if the token is valid.
// 	if claims, ok := parsedToken.Claims.(*MyClaims); ok && parsedToken.Valid {
// 		// Return user profile data.
// 		return gin.H{"username": claims.Username}, nil
// 	}

// 	return nil, err
// }

// func loginHandler(c *gin.Context) {
// 	var user User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	// Check if the user exists and the password is correct (in-memory authentication).
// 	validUser, ok := users[user.Username]
// 	if !ok || validUser.Password != user.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Generate a JWT token for the authenticated user.
// 	token, err := generateToken(user.Username)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
// 		return
// 	}

// 	// Return the JWT token in the response.
// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// func profileHandler(c *gin.Context) {
// 	// Extract the JWT token from the Authorization header.
// 	tokenString := c.GetHeader("Authorization")
// 	if tokenString == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token in Authorization header"})
// 		return
// 	}

// 	// Check if the token starts with "Bearer ".
// 	if !strings.HasPrefix(tokenString, "Bearer ") {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 		return
// 	}

// 	// Parse the JWT token.
// 	parsedToken, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
// 		return
// 	}

// 	// Check if the token is valid.
// 	if claims, ok := parsedToken.Claims.(*MyClaims); ok && parsedToken.Valid {
// 		// Return user profile data.
// 		c.JSON(http.StatusOK, gin.H{"username": claims.Username})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 	}
// }
