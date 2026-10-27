package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v3"
    "latihan-fiber/internal/model"
)

func DaftarMahasiswa(c fiber.Ctx) error {
    angkatan := c.Query("angkatan", "")
    if angkatan == "" {
        return c.JSON(fiber.Map{
            "sukses": true,
            "data":   model.Daftar,
        })
    }

    nilaiAngkatan, err := strconv.Atoi(angkatan)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Parameter angkatan harus berupa angka")
    }

    hasil := []model.Mahasiswa{}
    for i := 0; i < len(model.Daftar); i++ {
        if model.Daftar[i].Angkatan == nilaiAngkatan {
            hasil = append(hasil, model.Daftar[i])
        }
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "data":   hasil,
    })
}

func DetailMahasiswa(c fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
    }

    for i := 0; i < len(model.Daftar); i++ {
        if model.Daftar[i].ID == id {
            return c.JSON(fiber.Map{
                "sukses": true,
                "data":   model.Daftar[i],
            })
        }
    }

    return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
}