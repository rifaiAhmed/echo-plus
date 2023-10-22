package kreditPlusRepository

// fai

import (
	"test-echo/db"
	models "test-echo/model"
)

func CekLimit(NIK string) ([]models.Limit, error) {
	var obj []models.Limit
	db := db.CreateCon()
	err := db.Where("NIK = ?", NIK).Find(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func CekDataUser(NIK string) (models.Consumer, error) {
	var obj models.Consumer
	db := db.CreateCon()
	err := db.Where("NIK = ?", NIK).Find(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
