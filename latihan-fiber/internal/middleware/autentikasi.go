package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"latihan-fiber/internal/auth"
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

func Autentikasi() fiber.Handler {
	konf := config.Muat()

	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Header Authorization tidak ditemukan")
		}

		bagian := strings.Split(header, " ")
		if len(bagian) != 2 || bagian[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "Format header Authorization tidak sesuai")
		}

		// Cek apakah token masuk daftar hitam (Tugas Praktikum Modul 8 #3)
		var blacklist model.TokenBlacklist
		if err := config.DB.Where("token = ?", bagian[1]).First(&blacklist).Error; err == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token sudah dicabut (telah logout)")
		}

		klaim, err := auth.PeriksaToken(bagian[1], konf.JwtSecret)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token tidak valid atau sudah kedaluwarsa")
		}

		id, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Isi token tidak sesuai")
		}

		c.Locals("id_pengguna", uint(id))
		c.Locals("peran", klaim.Peran)
		c.Locals("token", bagian[1])

		return c.Next()
	}
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peran, cocok := c.Locals("peran").(string)
		if cocok == false {
			return fiber.NewError(fiber.StatusForbidden, "Peran pengguna tidak diketahui")
		}

		for i := 0; i < len(peranDiizinkan); i++ {
			if peranDiizinkan[i] == peran {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Anda tidak berwenang mengakses sumber daya ini")
	}
}

