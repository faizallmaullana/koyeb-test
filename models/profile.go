package models

import "time"

// staff
// siswa

type Staff struct {
	ID           string    `json:"id" grom:"primary_key"`
	Nip          string    `json:"nip"`
	Nama         string    `json:"nama"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Alamat       string    `json:"alamat"`
	Telpon       string    `json:"telpon"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
