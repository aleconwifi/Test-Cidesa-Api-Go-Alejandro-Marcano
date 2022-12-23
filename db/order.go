package db

import (
	"article/model"
)

func CreateOrder(order *model.Order) (*model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	err = db.Create(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func UpdateOrder(updatedOrder *model.Order) (*model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	order := new(model.Order)

	err = db.Where("id = ?", updatedOrder.ID).First(&order).Error
	if err != nil {
		return nil, err
	}

	order.NROOrder = updatedOrder.NROOrder
	order.Items = updatedOrder.Items
	order.PromotionID = updatedOrder.PromotionID
	order.Total = updatedOrder.Total
	order.TotalDiscount = updatedOrder.TotalDiscount
	err = db.Save(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrders(userID uint64) (*[]model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	orders := new([]model.Order)

	err = db.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil

}

func GetOrder(id uint64) (*model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	order := new(model.Order)

	err = db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func DeleteOrder(id uint64) error {
	db, dbc, err := GetDBCon()
	if err != nil {
		return err
	}
	defer dbc.Close()

	err = db.Where("id = ? ", id).Delete(&model.Order{}).Error

	return err
}
