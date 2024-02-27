package admin

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	jwt_auth "github.com/faizallmaullana/test-koyeb/Authentication"
	"github.com/faizallmaullana/test-koyeb/controllers"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type inputRegristration struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	Token          string `json:"token"`
	Nip            string `json:"nip"`
	Nama           string `json:"nama"`
	TanggalLahir   string `json:"tanggal_lahir"`
	JenisKelaminan string `json:"jenis_kelamin"`
	Alamat         string `json:"alamat"`
	Telpon         string `json:"telpon"`
}

func Registration(c *gin.Context) {
	var input inputRegristration
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// unique username
	var auth models.Authentication
	if err := models.DB.Where("username = ?", input.Username).First(&auth).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already used"})
		return
	}

	if err := models.DB.First(&auth).Error; err != nil {
		if err.Error() == "record not found" {
			// unique nip
			var nip string
			var unique_nip models.Staff
			if err := models.DB.Where("nip = ?", input.Nip).First(&unique_nip).Error; err == nil {
				if input.Nip == "0" {
					nip = "0"
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "NIP already used"})
					return
				}
			}
			nip = input.Nip

			// check password strength
			var password string
			isValid, err := controllers.CheckPasswordStrength(input.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			} else if isValid {
				password, _ = controllers.HashPassword(input.Password)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Password strength check failed."})
				return
			}

			if input.Token != "tokenAdmin" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
				return
			}

			// write to db
			created := time.Now().UTC().Add(7 * time.Hour)

			date := input.TanggalLahir
			layout := "02-01-2006"

			location, err := time.LoadLocation("Asia/Bangkok")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error Location": err})
				return
			}

			parsedTanggalLahir, err := time.ParseInLocation(layout, date, location)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error Time": err})
				return
			}

			data_staff := models.Staff{
				ID:           uuid.New().String(),
				Nip:          nip,
				Nama:         input.Nama,
				TanggalLahir: parsedTanggalLahir,
				JenisKelamin: input.JenisKelaminan,
				Alamat:       input.Alamat,
				Telpon:       input.Telpon,
				CreatedAt:    created,
				IsDeleted:    false,
			}

			data_registration := models.Authentication{
				ID:       uuid.New().String(),
				Username: input.Username,
				Password: password,
				Role:     input.Role,
				StaffID:  data_staff.ID,
			}

			models.DB.Create(&data_registration)
			models.DB.Create(&data_staff)

			// generate new token
			rand.Seed(time.Now().UnixNano())
			code := fmt.Sprintf("%06d", rand.Intn(1000000))

			db_token := models.Token{
				ID:    uuid.New().String(),
				Token: code,
			}

			models.DB.Create(&db_token)

			// generate authentication token using jwt
			tokenJWT, err := jwt_auth.GenerateToken(input.Username, data_staff.ID, input.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":       "success",
				"authorization": tokenJWT,
				"role":          input.Role,
			})
			return
		} else {
			return
		}
	}

	// tokenString := c.GetHeader("Authorization")
	// profileData, err := jwt_auth.JWTClaims(tokenString)

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
	// 	return
	// }

	// unique nip
	var nip string
	var unique_nip models.Staff
	if err := models.DB.Where("nip = ?", input.Nip).First(&unique_nip).Error; err == nil {
		if input.Nip == "0" {
			nip = "0"
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "NIP already used"})
			return
		}
	}
	nip = input.Nip

	// check password strength
	var password string
	isValid, err := controllers.CheckPasswordStrength(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if isValid {
		password, _ = controllers.HashPassword(input.Password)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password strength check failed."})
		return
	}

	// token validation
	var valid_token models.Token
	if err := models.DB.First(&valid_token).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Can't access db token"})
		return
	}

	if input.Role != "guru" {
		if input.Token != valid_token.Token {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			return
		}
	}

	// write to db
	created := time.Now().UTC().Add(7 * time.Hour)

	date := input.TanggalLahir
	layout := "02-01-2006"

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Location": err})
		return
	}

	parsedTanggalLahir, err := time.ParseInLocation(layout, date, location)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Time": err})
		return
	}

	data_staff := models.Staff{
		ID:           uuid.New().String(),
		Nip:          nip,
		Nama:         input.Nama,
		TanggalLahir: parsedTanggalLahir,
		JenisKelamin: input.JenisKelaminan,
		Alamat:       input.Alamat,
		Telpon:       input.Telpon,
		CreatedAt:    created,
		IsDeleted:    false,
	}

	data_registration := models.Authentication{
		ID:       uuid.New().String(),
		Username: input.Username,
		Password: password,
		Role:     input.Role,
		StaffID:  data_staff.ID,
	}

	models.DB.Create(&data_registration)
	models.DB.Create(&data_staff)

	// regenerate token login
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	valid_token.Token = code

	models.DB.Save(&valid_token)

	// generate authentication token using jwt
	tokenJWT, err := jwt_auth.GenerateToken(input.Username, data_staff.ID, input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"authorization": tokenJWT,
		"role":          input.Role,
	})
}

// Get Input JSON ###
// Username must uniquely identify ###
// if first person registration, do another function ###
// unique nip (db staff) if non-ASN nip = 0 ###
// Password must strength ###
// Token checks validity (firstToken is "tokenAdmin") ###
// regenerate token code ###
// Create Data on DB ###
// regenerate token ###
// Generate Authentication ###
// Return Value ###
