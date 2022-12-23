package db

import (
	"article/model"
)

func CreatePromotion(Promotion *model.Promotion) (*model.Promotion, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	err = db.Create(Promotion).Error
	if err != nil {
		return nil, err
	}

	return Promotion, nil
}

func UpdatePromotion(updatedPromotion *model.Promotion) (*model.Promotion, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	Promotion := new(model.Promotion)

	err = db.Where("id = ?", updatedPromotion.ID).First(&Promotion).Error
	if err != nil {
		return nil, err
	}

	Promotion.Name = updatedPromotion.Name
	Promotion.Code = updatedPromotion.Code
	Promotion.Used = updatedPromotion.Used

	err = db.Save(&Promotion).Error
	if err != nil {
		return nil, err
	}

	return Promotion, nil
}

func GetPromotions() (*[]model.Promotion, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	Promotions := new([]model.Promotion)

	err = db.Find(&Promotions).Error
	if err != nil {
		return nil, err
	}

	return Promotions, nil
}

func GetPromotion(id uint64) (*model.Promotion, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	Promotion := new(model.Promotion)

	err = db.Where("id = ?", id).First(&Promotion).Error
	if err != nil {
		return nil, err
	}

	return Promotion, nil
}

func DeletePromotion(id uint64) error {
	db, dbc, err := GetDBCon()
	if err != nil {
		return err
	}
	defer dbc.Close()

	err = db.Where("id = ? ", id).Delete(&model.Promotion{}).Error

	return err
}

func GetOrderByPromotion(id uint64) (*model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	order := new(model.Order)

	err = db.Where("promotion_id = ? ", id).First(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}
