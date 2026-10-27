package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Konfigurasi struct {
    NamaAplikasi string
    Port         string
    Lingkungan   string
    ApiKey       string
}

func Muat() Konfigurasi {
    godotenv.Load()
    return Konfigurasi{
        NamaAplikasi: ambil("APP_NAME", "Fiber API"),
        Port:         ambil("APP_PORT", "3000"),
        Lingkungan:   ambil("APP_ENV", "local"),
        ApiKey:       ambil("API_KEY", "rahasia123"),
    }
}

func ambil(kunci string, bawaan string) string {
    nilai := os.Getenv(kunci)
    if nilai == "" {
        return bawaan
    }
    return nilai
}