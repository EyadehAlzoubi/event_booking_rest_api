package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,date_time, user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAll() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		var dateTime string
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		event.DateTime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTime)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	var dateTime string
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTime, &event.UserID)

	event.DateTime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTime)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {

	query := `
	DELETE FROM events
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err

}
