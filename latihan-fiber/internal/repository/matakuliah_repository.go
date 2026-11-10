package repository

import (
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

func AmbilSemuaMatakuliah() ([]model.Matakuliah, error) {
	var daftar []model.Matakuliah
	err := config.DB.Find(&daftar).Error
	return daftar, err
}

func AmbilMatakuliahPerId(id uint) (model.Matakuliah, error) {
	var mk model.Matakuliah
	err := config.DB.First(&mk, id).Error
	return mk, err
}

func SimpanMatakuliah(mk *model.Matakuliah) error {
	return config.DB.Create(mk).Error
}

func PerbaruiMatakuliah(mk *model.Matakuliah) error {
	return config.DB.Save(mk).Error
}

func HapusMatakuliah(mk *model.Matakuliah) error {
	return config.DB.Delete(mk).Error
}

