package models

type Mahasiswa struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	NRP       string `json:"nrp"`
	JurusanID int    `json:"jurusan_id"`
	Angkatan  int    `json:"angkatan"`
}
