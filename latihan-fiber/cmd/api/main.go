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

	config.HubungkanBasisData(konf)
	config.Migrasi()
	config.DataAwal()

	app := fiber.New(fiber.Config{
		AppName:      konf.NamaAplikasi,
		ErrorHandler: penanganGalat,
	})

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recoverer.New())
	app.Use(middleware.WaktuProses())

	router.Daftarkan(app)

	// Handler 404 untuk rute yang tidak ditemukan (ditempatkan setelah rute terdaftar)
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Endpoint tidak tersedia",
		})
	})

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

	if kode == fiber.StatusInternalServerError {
		log.Println("Galat internal:", err.Error())
	}

	return c.Status(kode).JSON(fiber.Map{
		"sukses": false,
		"pesan":  pesan,
	})
}
