package middleware

import (
    "github.com/gofiber/fiber/v3"
)

func KunciApi(kunciValid string) fiber.Handler {
    return func(c fiber.Ctx) error {
        kunciInput := c.Get("X-Api-Key")

        if kunciInput == "" || kunciInput != kunciValid {
            return fiber.NewError(fiber.StatusUnauthorized, "Akses ditolak: X-Api-Key tidak valid atau tidak ditemukan")
        }

        return c.Next()
    }
}