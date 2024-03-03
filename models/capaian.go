package models

type CapaianSiswa struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreignKey
	MateriAjarID string     `json:"materi_ajar_id"`
	MateriAjar   MateriAjar `json:"materi_ajar" gorm:"references:MateriAjarID"`

	SiswaID string `json:"siswa_id"`
	Siswa   Siswa  `json:"siswa" gorm:"references:SiswaID"`
}
