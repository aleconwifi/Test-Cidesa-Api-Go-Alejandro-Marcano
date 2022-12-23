package db

import "article/model"

func Migrate() error {
	db, err := DBCon()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	// Initiate factories table
	if !db.Migrator().HasTable(&model.User{}) {
		err := db.AutoMigrate(&model.User{}, &model.Order{}, &model.Order{}, &model.Promotion{}, &model.ItemsOrders{})
		if err != nil {
			return err
		}

		/// Create default admin user ///

		err = db.Create(&model.User{
			Name:   "SuperAdmin",
			Type:   model.UserTypeAdmin,
			Active: true,
		}).Error
		if err != nil {
			return err
		}
	}

	defer dbc.Close()

	return nil
}
