package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func HubungkanBasisData(konf Konfigurasi) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		konf.DbUser,
		konf.DbSandi,
		konf.DbHost,
		konf.DbPort,
		konf.DbNama,
	)

	koneksi, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		log.Fatal("Gagal terhubung ke basis data: ", err)
	}

	DB = koneksi
	log.Println("Koneksi basis data berhasil")
}

