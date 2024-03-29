package admin

// REGISTRATION

// LIST FUNCTION
// Registration

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	jwt_auth "github.com/faizallmaullana/test-koyeb/Authentication"
	"github.com/faizallmaullana/test-koyeb/controllers/hashing"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type inputRegistration struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Token        int    `json:"token"`
	Nip          string `json:"nip"`
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Telpon       string `json:"telpon"`
}

// ALGORITHM
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

func Registration(c *gin.Context) {
	var input inputRegistration
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check unique username
	var unique_username models.Authentication
	if err := models.DB.Where(" username = ? ", input.Username).First(&unique_username).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username number already used"})
		return
	}

	// reparsing nomor telpon
	var nomorTelpon string

	if len(input.Telpon) > 0 && input.Telpon[:2] == "08" {
		nomorTelpon = input.Telpon
	} else if len(input.Telpon) > 0 && input.Telpon[:4] == "+628" {
		nomorTelpon = "0" + input.Telpon[3:]
	} else if len(input.Telpon) > 0 && input.Telpon[:3] == "628" {
		nomorTelpon = "0" + input.Telpon[2:]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone number",
		})
		return
	}

	// reparsing tanggal lahir
	date := input.TanggalLahir
	layout := "2006-01-02"

	parsedTanggalLahir, err := time.Parse(layout, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Time": err})
		return
	}

	// check password strength
	var password string
	isValid, err := hashing.CheckPasswordStrength(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if isValid {
		password, _ = hashing.HashPassword(input.Password)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password strength check failed."})
		return
	}

	// validate the token
	var CekToken models.Token
	if err := models.DB.First(&CekToken).Error; err != nil {
		if err.Error() == "record not found" {

			if input.Token != 111111 {
				c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not accepted"})
				return
			}

			// Seed the random number generator
			rand.Seed(time.Now().UnixNano())

			// Generate a random 6-digit number
			min := 100000 // minimum value of a 6-digit number
			max := 999999 // maximum value of a 6-digit number
			randomNum := min + rand.Intn(max-min+1)

			inputToken := models.Token{
				ID:    uuid.New().String(),
				Token: randomNum,
			}

			models.DB.Create(&inputToken)

			input.Token = randomNum
			CekToken.Token = randomNum
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
	}

	if input.Token != CekToken.Token {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not accepted"})
		return
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	min := 100000 // minimum value of a 6-digit number
	max := 999999 // maximum value of a 6-digit number
	randomNum := min + rand.Intn(max-min+1)

	CekToken.Token = randomNum

	models.DB.Model(&CekToken).Update(CekToken)

	// write to the database (users and profiles)
	dbStaff := models.Staff{
		ID:           uuid.New().String(),
		Nip:          input.Nip,
		Nama:         input.Nama,
		TanggalLahir: parsedTanggalLahir,
		JenisKelamin: input.JenisKelamin,
		Alamat:       input.Alamat,
		Telpon:       nomorTelpon,

		CreatedAt: time.Now().UTC().Add(7 * time.Hour),
		IsDeleted: false,
	}

	dbAuth := models.Authentication{
		ID:       uuid.New().String(),
		Username: strings.ToLower(input.Username),
		Password: password,
		Role:     input.Role,
		StaffID:  dbStaff.ID,
	}

	models.DB.Create(&dbAuth)
	models.DB.Create(&dbStaff)

	// generate authentication token using jwt
	tokenJWT, err := jwt_auth.GenerateToken(dbAuth.ID, dbAuth.Username, dbAuth.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	fmt.Println(tokenJWT)

	// return value
	c.JSON(http.StatusCreated, gin.H{
		"message":   "created success",
		"tokenAuth": tokenJWT,
		"user":      dbStaff,
	})
}
