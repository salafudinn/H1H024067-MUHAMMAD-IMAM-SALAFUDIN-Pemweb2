package handler

import (
	"github.com/gofiber/fiber/v3"
	"latihan-fiber/internal/repository"
)

func DaftarProgramStudi(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaProgramStudi()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data program studi")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

