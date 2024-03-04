package siswa

// SISWA

// LIST FUNCTION
// PostSiswa
// PostImageSiswa

// IN FRONTEND
// 1. (PostSiswa) post the data without image using header content type "application/json"
// 2. get response id
// 3. (PostImageSiswa) update the image using header content type "multipart/formdata"

// in the bottom of this page there a function get data but not complete yet

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputSiswa struct {
	Nisn            string    `json:"nisn"`
	Nama            string    `json:"nama"`
	NamaAyahKandung string    `json:"nama_ayah_kandung"`
	NamaIbuKandung  string    `json:"nama_ibu_kandung"`
	TanggalLahir    string    `json:"tanggal_lahir"`
	JenisKelamin    string    `json:"jenis_kelaminan"`
	Alamat          string    `json:"alamat"`
	CreatedAt       time.Time `json:"created_at"`
}

func PostSiswa(c *gin.Context) {
	var input InputSiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, "Invalid requestt")
		return
	}

	// parsing tanggal lahir
	date := input.TanggalLahir
	layout := "2006-01-02"

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

	// post to the db
	siswa := models.Siswa{
		ID:              uuid.New().String(),
		Nisn:            input.Nisn,
		Nama:            input.Nama,
		NamaAyahKandung: input.NamaAyahKandung,
		NamaIbuKandung:  input.NamaIbuKandung,
		TanggalLahir:    parsedTanggalLahir,
		JenisKelamin:    input.JenisKelamin,
		Alamat:          input.Alamat,
		CreatedAt:       time.Now().UTC().Add(7 * time.Hour),
	}

	models.DB.Create(&siswa)
	c.JSON(http.StatusOK, siswa)
}

func PostImageSiswa(c *gin.Context) {
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

// func GetTester(c *gin.Context) {
// 	var siswaList []models.Siswa
// 	if err := models.DB.Find(&siswaList).Error; err != nil {
// 		c.String(http.StatusInternalServerError, "Unable to fetch Siswa data from database")
// 		return
// 	}

// 	c.JSON(http.StatusOK, siswaList)
// }
