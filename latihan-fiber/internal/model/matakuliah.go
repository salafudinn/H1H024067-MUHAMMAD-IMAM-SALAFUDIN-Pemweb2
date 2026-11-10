package model

import "time"

type Matakuliah struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Kode      string    `gorm:"size:20;uniqueIndex" json:"kode"`
	Nama      string    `gorm:"size:100" json:"nama"`
	Sks       int       `json:"sks"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"dibuat_pada"`
	UpdatedAt time.Time `json:"diubah_pada"`
}

