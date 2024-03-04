package siswa

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InputSiswa struct {
	Nisn            string    `json:"nisn"`
	Nama            string    `json:"nama"`
	NamaAyahKandung string    `json:"nama_ayah_kandung"`
	NamaIbuKandung  string    `json:"nama_ibu_kandung"`
	TanggalLahir    time.Time `json:"tanggal_lahir"`
	JenisKelaminan  string    `json:"jenis_kelaminan"`
	Alamat          string    `json:"alamat"`
}

func PostSiswa(c *gin.Context) {
	var input InputSiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
