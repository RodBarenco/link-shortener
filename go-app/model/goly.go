package model

import "github.com/google/uuid"

func GetAllGolies() ([]Goly, error) {
	var golies []Goly

	tx := db.Find(&golies)
	if tx.Error != nil {
		return []Goly{}, tx.Error
	}

	return golies, nil
}

func GetGoly(id uuid.UUID) (Goly, error) {
	var goly Goly

	tx := db.Where("id = ?", id).First(&goly)
	if tx.Error != nil {
		return Goly{}, tx.Error
	}

	return goly, nil
}

func CreateGoly(goly Goly) error {
	tx := db.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly Goly) error {
	tx := db.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uuid.UUID) error {
	tx := db.Unscoped().Delete(&Goly{}, id)
	return tx.Error
}

func FindByUrl(url string) (Goly, error) {
	var goly Goly
	tx := db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}
