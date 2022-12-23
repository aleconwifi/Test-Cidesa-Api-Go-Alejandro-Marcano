package db

import (
	"article/model"
)

func CreateItem(item *model.Item) (*model.Item, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	err = db.Create(item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func UpdateItem(updatedItem *model.Item) (*model.Item, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	item := new(model.Item)

	err = db.Where("id = ?", updatedItem.ID).First(&item).Error
	if err != nil {
		return nil, err
	}

	item.Name = updatedItem.Name
	item.Price = updatedItem.Price
	item.Available = updatedItem.Available

	err = db.Save(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func GetItems(name, availability string) (*[]model.Item, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	items := new([]model.Item)

	if name != "" && availability != "" {
		av := false
		if availability == "True" || availability == "true" {
			av = true
		}
		err = db.Where("name LIKE ?  AND available = ? ", "%"+name+"%", av).Find(items).Error
		if err != nil {
			return nil, err
		}
		return items, nil

	} else if name != "" || availability != "" {

		if availability != "" {
			av := false
			if availability == "True" || availability == "true" {
				av = true
			}
			err = db.Where("available = ? ", av).Find(items).Error
			if err != nil {
				return nil, err
			}

			return items, nil

		}

		if name != "" {
			err = db.Where("name LIKE ? AND available = ?", "%"+name+"%", true).Find(items).Error
			if err != nil {
				return nil, err
			}
			return items, nil
		}

	}

	err = db.Where("available = ?", true).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItem(id uint64) (*model.Item, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	item := new(model.Item)

	err = db.Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteItem(id uint64) error {
	db, dbc, err := GetDBCon()
	if err != nil {
		return err
	}
	defer dbc.Close()

	err = db.Where("id = ? ", id).Delete(&model.Item{}).Error

	return err
}

func GetItemOrders(id uint64) (*[]model.ItemsOrders, error) {
	db, dbc, err := GetDBCon()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	itemOrders := new([]model.ItemsOrders)

	err = db.Where("item_id  = ? ", id).Find(itemOrders).Error
	if err != nil {
		return nil, err
	}

	return itemOrders, err
}
