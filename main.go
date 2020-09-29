package main

import (
	database "store-app/src/data"
	"store-app/src/data/models"
)

func main() {
	db := database.ConnectSQL()

	var result *models.Business
	db.Raw("SELECT * FROM businesses").Scan(&result)
	/*rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		log.Fatal("Error while retrieving rows")
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			firstName string
			lastName  string
			created   string
		)

		err := rows.Scan(&id, &firstName, &lastName, &created)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("id: %d, first name: %s, last name: %s, created at: %s\n", id, firstName, lastName, created)
	}

	defer db.Close()*/
}
