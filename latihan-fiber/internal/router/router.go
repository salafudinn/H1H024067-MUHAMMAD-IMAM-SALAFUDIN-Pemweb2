package router

import (
	"github.com/gofiber/fiber/v3"
	"latihan-fiber/internal/handler"
	"latihan-fiber/internal/middleware"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	api.Post("/auth/registrasi", handler.Registrasi)
	api.Post("/auth/login", handler.Login)

	// Endpoint Program Studi (Modul 7 Tugas Praktikum #2)
	api.Get("/program-studi", handler.DaftarProgramStudi)

	// Rute terlindungi dengan autentikasi JWT
	terlindungi := api.Group("", middleware.Autentikasi())
	terlindungi.Get("/auth/profil", handler.Profil)
	terlindungi.Post("/auth/segarkan", handler.SegarkanToken) // Modul 8 Tugas Praktikum #1
	terlindungi.Post("/auth/logout", handler.Logout)          // Modul 8 Tugas Praktikum #3

	// Rute Mahasiswa CRUD (Modul 7 Langkah 10 & Modul 8 Langkah 10)
	mahasiswa := terlindungi.Group("/mahasiswa")
	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Post("/", middleware.HanyaPeran("admin", "dosen"), handler.BuatMahasiswa)
	mahasiswa.Put("/:id", middleware.HanyaPeran("admin", "dosen"), handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", middleware.HanyaPeran("admin"), handler.HapusMahasiswa)
	mahasiswa.Post("/:id/pulihkan", middleware.HanyaPeran("admin"), handler.PulihkanMahasiswa) // Modul 7 Tugas Praktikum #3

	// Rute Matakuliah CRUD (Modul 7 Tugas Praktikum #1)
	matakuliah := terlindungi.Group("/matakuliah")
	matakuliah.Get("/", handler.DaftarMatakuliah)
	matakuliah.Get("/:id", handler.DetailMatakuliah)
	matakuliah.Post("/", middleware.HanyaPeran("admin", "dosen"), handler.BuatMatakuliah)
	matakuliah.Put("/:id", middleware.HanyaPeran("admin", "dosen"), handler.PerbaruiMatakuliah)
	matakuliah.Delete("/:id", middleware.HanyaPeran("admin"), handler.HapusMatakuliah)
}
