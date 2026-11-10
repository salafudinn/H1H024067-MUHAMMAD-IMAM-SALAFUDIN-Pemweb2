package validasi

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var pemeriksa = validator.New()

type PesanGalat struct {
	Field string `json:"field"`
	Pesan string `json:"pesan"`
}

func init() {
	// Validasi kustom untuk memastikan NIM diawali huruf H diikuti angka (Tugas Praktikum Modul 8 #2)
	pemeriksa.RegisterValidation("nim_h", func(fl validator.FieldLevel) bool {
		nilai := fl.Field().String()
		cocok, _ := regexp.MatchString(`^H[0-9]`, nilai)
		return cocok
	})
}

func Periksa(data any) []PesanGalat {
	err := pemeriksa.Struct(data)
	if err == nil {
		return nil
	}

	daftar := []PesanGalat{}
	for _, item := range err.(validator.ValidationErrors) {
		daftar = append(daftar, PesanGalat{
			Field: item.Field(),
			Pesan: susunPesan(item),
		})
	}
	return daftar
}

func susunPesan(item validator.FieldError) string {
	switch item.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Nilai minimal adalah " + item.Param()
	case "max":
		return "Nilai maksimal adalah " + item.Param()
	case "gte":
		return "Nilai tidak boleh kurang dari " + item.Param()
	case "lte":
		return "Nilai tidak boleh lebih dari " + item.Param()
	case "nim_h":
		return "NIM harus diawali huruf H diikuti angka"
	default:
		return "Nilai tidak memenuhi aturan " + item.Tag()
	}
}

