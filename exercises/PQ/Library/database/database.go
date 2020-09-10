package database

import (
	"excercise-library/config"
	"excercise-library/ent"
	"fmt"
)

func Connect(configObj config.Config) (*ent.Client, error) {
	client, err := ent.Open(configObj.DB.DB_TYPE, ConnectionString(configObj.DB))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectionString(dbConfig config.Database) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True",
		dbConfig.DB_USER,
		dbConfig.DB_PASS,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME,
	)
}
