package admin

// REGISTRATION

// LIST FUNCTION
// Registration

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt_auth "github.com/faizallmaullana/test-koyeb/Authentication"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func RegistrationGuru(c *gin.Context) {
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
		Password: "anonymous",
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
