package handlers

import (
	"database/sql"
	models "holamundo/Models"
	"log"
)

func MyList(db *sql.DB) []models.Contact {

	query := "SELECT * FROM test_go"

	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	contacts := []models.Contact{}

	for row.Next() {

		test := models.Contact{}
		err := row.Scan(&test.Id, &test.Nombre, &test.Apellido, &test.Telefono)

		if err != nil {
			log.Fatal(err)
		}
		contacts = append(contacts, test)
	}
	return contacts
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO test_go (nombre, apellido, telefono) VALUES (?, ?, ?)"

	_, err := db.Exec(query, contact.Nombre, contact.Apellido, contact.Telefono)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nuevo usuario registrado")
}
