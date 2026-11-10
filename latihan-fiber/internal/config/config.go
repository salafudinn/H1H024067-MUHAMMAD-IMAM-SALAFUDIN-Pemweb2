package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Konfigurasi struct {
	NamaAplikasi string
	Port         string
	Lingkungan   string
	ApiKey       string
	DbHost       string
	DbPort       string
	DbNama       string
	DbUser       string
	DbSandi      string
	JwtSecret    string
	JwtTtlMenit  int
}

func Muat() Konfigurasi {
	godotenv.Load()

	ttl, err := strconv.Atoi(ambil("JWT_TTL_MENIT", "120"))
	if err != nil {
		ttl = 120
	}

	return Konfigurasi{
		NamaAplikasi: ambil("APP_NAME", "Fiber API"),
		Port:         ambil("APP_PORT", "3000"),
		Lingkungan:   ambil("APP_ENV", "local"),
		ApiKey:       ambil("API_KEY", "rahasia123"),
		DbHost:       ambil("DB_HOST", "127.0.0.1"),
		DbPort:       ambil("DB_PORT", "3306"),
		DbNama:       ambil("DB_NAME", "pemweb2_fiber"),
		DbUser:       ambil("DB_USER", "root"),
		DbSandi:      ambil("DB_PASS", ""),
		JwtSecret:    ambil("JWT_SECRET", "rahasia-bawaan"),
		JwtTtlMenit:  ttl,
	}
}

func ambil(kunci string, bawaan string) string {
	nilai := os.Getenv(kunci)
	if nilai == "" {
		return bawaan
	}
	return nilai
}
