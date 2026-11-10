package repository

import (
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

type ProgramStudiRespon struct {
	model.ProgramStudi
	JumlahMahasiswa int64 `json:"jumlah_mahasiswa"`
}

func AmbilSemuaProgramStudi() ([]ProgramStudiRespon, error) {
	var prodis []model.ProgramStudi
	err := config.DB.Find(&prodis).Error
	if err != nil {
		return nil, err
	}

	var hasil []ProgramStudiRespon
	for _, p := range prodis {
		var count int64
		config.DB.Model(&model.Mahasiswa{}).Where("program_studi_id = ?", p.ID).Count(&count)
		hasil = append(hasil, ProgramStudiRespon{
			ProgramStudi:    p,
			JumlahMahasiswa: count,
		})
	}
	return hasil, nil
}

