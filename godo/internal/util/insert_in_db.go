package util

func InsertTaskInDB(title *string, description *string, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare("INSERT INTO godo (title, description) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(*title, *description)
	if err != nil {
		return err
	}

	// Successfully inserted task into DB
	return nil
}
