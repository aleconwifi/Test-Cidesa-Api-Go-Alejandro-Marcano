package db

import (
	"article/model"
)

// Only user can create user
func CreateUser(user *model.User) (*model.User, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	err = db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(id uint64) (*model.User, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	user := new(model.User)

	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers() (*[]model.User, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	users := new([]model.User)

	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(id uint64) error {
	db, dbc, err := GetDBCon()
	if err != nil {
		return err
	}
	defer dbc.Close()

	err = db.Where("id = ? ", id).Delete(&model.User{}).Error

	return err
}

// Update user updates the name of the user
func UpdateUser(updatedUser *model.User) (*model.User, error) {

	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	user := new(model.User)

	err = db.Where("id = ?", updatedUser.ID).First(&user).Error
	if err != nil {
		return nil, err
	}

	user.Name = updatedUser.Name

	err = db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetOrderByUser(id uint64) (*model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	order := new(model.Order)

	err = db.Where("user_id = ? ", id).First(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrdersByUser(id uint64) (*[]model.Order, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	orders := new([]model.Order)

	err = db.Where("user_id = ? ", id).Find(orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
