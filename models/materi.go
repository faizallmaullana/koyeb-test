package models

type MataPelajaran struct {
	ID      string `json:"id" gorm:"primary_key"`
	Tingkat string `json:"tingkat"`

	// foreignkey
	JurusanID string  `json:"jurusan_id"`
	Jurusan   Jurusan `json:"jurusan" gorm:"references:JurusanID"`
}

type MateriAjar struct {
	ID     string `json:"id" gorm:"primary_key"`
	Materi string `json:"materi"`
	Bobot  string `json:"bobot"`

	// foreignkey
	MataPelajaranID string        `json:"mata_pelajaran_id"`
	MataPelajaran   MataPelajaran `json:"mata_pelajaran" gorm:"references:MataPelajaranID"`
}
