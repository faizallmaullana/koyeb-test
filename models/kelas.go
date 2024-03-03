package models

import "time"

type Jurusan struct {
	ID      string `json:"id" gorm:"primary_key"`
	Jurusan string `json:"jurusan"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Kelas struct {
	ID          string `json:"id" gorm:"primary_key"`
	Tingkat     string `json:"tingkat"`
	NomorKelas  string `json:"nomor_kelas"`
	TahunAjaran string `json:"tahun_ajaran"`

	// foreign keys
	JurusanID string  `json:"jurusan_id"`
	Jurusan   Jurusan `json:"jurusan" gorm:"references:JurusanID"`

	WaliKelasID string `json:"wali_kelas_id"`
	WaliKelas   Staff  `json:"wali_kelas" gorm:"references:WaliKelasID"`
}
