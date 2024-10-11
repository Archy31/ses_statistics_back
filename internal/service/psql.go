package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"ses_back/internal/config"
	"ses_back/internal/models"
)

func connect2DB() *sql.DB {
	config.ReadConfig()
	psqlCfg := config.Config.Psql
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bishkek",
		psqlCfg.Host, psqlCfg.User, psqlCfg.Password, psqlCfg.DbName, psqlCfg.Port,
	)

	db, err := sql.Open(psqlCfg.DriverName, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

//func CreateTable() {
//	var db = connect2DB()
//	defer func(open *sql.DB) {
//		err := open.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(db)
//
//	query := `
//			  CREATE TABLE IF NOT EXISTS bulls (
//			  id SERIAL PRIMARY KEY,
//			  event_id VARCHAR(255) NOT NULL,
//			  epicenter VARCHAR(255) NOT NULL,
//			  mag NUMERIC(3,1) NOT NULL,
//			  ev_date VARCHAR(255) NOT NULL,
//			  ev_time VARCHAR(255) NOT NULL
//			)`
//
//	_, err := db.Exec(query)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func InsertIntoTable(row models.Bulls) int {
//	var db = connect2DB()
//	defer func(open *sql.DB) {
//		err := open.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(db)
//
//	query := `INSERT INTO bulls(event_id, epicenter, mag, ev_date, ev_time)
//			VALUES ($1, $2, $3, $4, $5) RETURNING id`
//
//	var pk int
//
//	err := db.QueryRow(query, row.Event_id, row.Epicenter, row.Mag, row.Ev_date, row.Ev_time).Scan(&pk)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return pk
//}

func DBselect(conditions string) ([]models.Bulls, error) {
	//config.ReadConfig()
	//psqlCfg := config.Config.Psql

	var db = connect2DB()
	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	var data []models.Bulls
	var id int32
	var eventId string
	var epicenter string
	var mag float32
	var evDate string
	var evTime string

	var query = `SELECT * FROM bulls ` + conditions
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&id, &eventId, &epicenter, &mag, &evDate, &evTime)
		if err != nil {
			return nil, err
		}
		data = append(data, models.Bulls{
			Id:        id,
			Event_id:  eventId,
			Epicenter: epicenter,
			Mag:       mag,
			Ev_date:   evDate,
			Ev_time:   evTime,
		})
	}

	return data, nil
}
