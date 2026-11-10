package handler

type PermintaanMahasiswa struct {
	ProgramStudiID uint    `json:"program_studi_id" validate:"required,gte=1"`
	Nim            string  `json:"nim" validate:"required,min=5,max=20,nim_h"`
	Nama           string  `json:"nama" validate:"required,min=3,max=100"`
	Email          string  `json:"email" validate:"required,email,max=100"`
	Angkatan       int     `json:"angkatan" validate:"required,gte=2000,lte=2100"`
	Ipk            float64 `json:"ipk" validate:"gte=0,lte=4"`
	Aktif          *bool   `json:"aktif"`
}

type PermintaanRegistrasi struct {
	Nama      string `json:"nama" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	KataSandi string `json:"kata_sandi" validate:"required,min=8"`
	Peran     string `json:"peran"`
}

type PermintaanLogin struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type PermintaanMatakuliah struct {
	Kode     string `json:"kode" validate:"required,min=2,max=20"`
	Nama     string `json:"nama" validate:"required,min=3,max=100"`
	Sks      int    `json:"sks" validate:"required,gte=1,lte=6"`
	Semester int    `json:"semester" validate:"required,gte=1,lte=8"`
}

