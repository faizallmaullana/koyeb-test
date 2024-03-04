package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Siswa struct {
	Nisn         string `json:"nisn"`
	TanggalLahir string `json:"tanggal_lahir"`
}

func Tester(c *gin.Context) {
	var siswa models.Siswa
	if err := c.Bind(&siswa); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	// image
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to get image")
		return
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to read image")
		return
	}

	var siswa2 models.Siswa
	if err := models.DB.Where("id = ?", c.Param("id")).First(&siswa2).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "gagal membuka db",
		})
		return
	}

	siswa2.Image = imageData

	if err := models.DB.Save(&siswa2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal menyimpan gambar",
		})
		return
	}

	c.String(http.StatusOK, "Siswa %s saved successfully", siswa.Nama)
}

func Tester2(c *gin.Context) {
	var input Siswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, "Invalid requestt")
		return
	}

	siswa := models.Siswa{
		ID:   uuid.New().String(),
		Nisn: input.Nisn,
	}

	models.DB.Create(&siswa)
	c.JSON(http.StatusOK, siswa)
}

func GetTester(c *gin.Context) {
	var siswaList []models.Siswa
	if err := models.DB.Find(&siswaList).Error; err != nil {
		c.String(http.StatusInternalServerError, "Unable to fetch Siswa data from database")
		return
	}

	c.JSON(http.StatusOK, siswaList)
}
