package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"latihan-fiber/internal/model"
	"latihan-fiber/internal/repository"
	"latihan-fiber/internal/validasi"
)

func DaftarMahasiswa(c fiber.Ctx) error {
	halaman, _ := strconv.Atoi(c.Query("halaman", "1"))
	perHalaman, _ := strconv.Atoi(c.Query("per_halaman", "10"))
	angkatan, _ := strconv.Atoi(c.Query("angkatan", "0"))

	if halaman < 1 {
		halaman = 1
	}

	if perHalaman < 1 || perHalaman > 100 {
		perHalaman = 10
	}

	filter := repository.FilterMahasiswa{
		Cari:       c.Query("cari", ""),
		Angkatan:   angkatan,
		Halaman:    halaman,
		PerHalaman: perHalaman,
	}

	daftar, total, err := repository.Ambil(filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data")
	}

	totalHalaman := int(total) / perHalaman
	if int(total)%perHalaman != 0 {
		totalHalaman = totalHalaman + 1
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
		"meta": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total":         total,
			"total_halaman": totalHalaman,
		},
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

func BuatMahasiswa(c fiber.Ctx) error {
	var permintaan PermintaanMahasiswa
	err := c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	galat := validasi.Periksa(permintaan)
	if galat != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}

	aktif := true
	if permintaan.Aktif != nil {
		aktif = *permintaan.Aktif
	}

	mahasiswa := model.Mahasiswa{
		ProgramStudiID: permintaan.ProgramStudiID,
		Nim:            permintaan.Nim,
		Nama:           permintaan.Nama,
		Email:          permintaan.Email,
		Angkatan:       permintaan.Angkatan,
		Ipk:            permintaan.Ipk,
		Aktif:          aktif,
	}

	err = repository.Simpan(&mahasiswa)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Gagal menyimpan data, periksa keunikan NIM dan email")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dibuat",
		"data":   mahasiswa,
	})
}

func PerbaruiMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}

	var permintaan PermintaanMahasiswa
	err = c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	if permintaan.Nama != "" {
		mahasiswa.Nama = permintaan.Nama
	}
	if permintaan.Email != "" {
		mahasiswa.Email = permintaan.Email
	}
	if permintaan.Angkatan > 0 {
		mahasiswa.Angkatan = permintaan.Angkatan
	}
	if permintaan.Ipk > 0 {
		mahasiswa.Ipk = permintaan.Ipk
	}
	if permintaan.ProgramStudiID > 0 {
		mahasiswa.ProgramStudiID = permintaan.ProgramStudiID
	}
	if permintaan.Aktif != nil {
		mahasiswa.Aktif = *permintaan.Aktif
	}

	err = repository.Perbarui(&mahasiswa)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Gagal memperbarui data")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil diperbarui",
		"data":   mahasiswa,
	})
}

func HapusMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}

	err = repository.Hapus(&mahasiswa)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus data")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dihapus",
	})
}

// PulihkanMahasiswa mengembalikan data mahasiswa yang terhapus secara soft delete (Tugas Praktikum Modul 7 #3)
func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	err = repository.Pulihkan(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memulihkan data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dipulihkan",
	})
}
