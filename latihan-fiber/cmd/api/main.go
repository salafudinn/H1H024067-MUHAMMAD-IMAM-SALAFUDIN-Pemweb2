package main

import (
    "errors"
    "log"
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/logger"
    recoverer "github.com/gofiber/fiber/v3/middleware/recover"
    "github.com/gofiber/fiber/v3/middleware/requestid"
    "latihan-fiber/internal/config"
    "latihan-fiber/internal/middleware"
    "latihan-fiber/internal/router"
)

func main() {
    konf := config.Muat()

    app := fiber.New(fiber.Config{
        AppName:      konf.NamaAplikasi,
        ErrorHandler: penanganGalat,
    })

    app.Use(requestid.New())
    app.Use(logger.New())
    app.Use(recoverer.New())
    app.Use(middleware.WaktuProses())

	app.Use(func(c fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "sukses": false,
        "pesan":  "Endpoint tidak tersedia",
    	})
	})

    router.Daftarkan(app)

    log.Fatal(app.Listen(":" + konf.Port))
}



func penanganGalat(c fiber.Ctx, err error) error {
    kode := fiber.StatusInternalServerError
    pesan := "Terjadi kesalahan pada server"

    var galatFiber *fiber.Error
    if errors.As(err, &galatFiber) {
        kode = galatFiber.Code
        pesan = galatFiber.Message
    }

    return c.Status(kode).JSON(fiber.Map{
        "sukses": false,
        "pesan":  pesan,
    })
}