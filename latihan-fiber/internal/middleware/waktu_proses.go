package middleware

import (
    "strconv"
    "time"
    "github.com/gofiber/fiber/v3"
)

func WaktuProses() fiber.Handler {
    return func(c fiber.Ctx) error {
        mulai := time.Now()
        err := c.Next()
        durasi := time.Since(mulai)
        c.Set("X-Waktu-Proses-Ms", strconv.FormatInt(durasi.Milliseconds(), 10))
        return err
    }
}