package models

import "time"

type Siswa struct {
	ID              string    `json:"id" gorm:"primary_key"`
	Nisn            string    `json:"nisn"`
	Nama            string    `json:"nama"`
	NamaAyahKandung string    `json:"nama_ayah_kandung"`
	NamaIbuKandung  string    `json:"nama_ibu_kandung"`
	TanggalLahir    time.Time `json:"tanggal_lahir"`
	JenisKelamin    string    `json:"jenis_kelamin"`
	Alamat          string    `json:"alamat"`
}

type PivotKelasSiswa struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreignkey
	SiswaID string `json:"siswa_id"`
	Siswa   Siswa  `json:"siswa" gorm:"references:SiswaID"`

	KelasID string `json:"kelas_id"`
	Kelas   Kelas  `json:"kelas" gorm:"references:KelasID"`
}
