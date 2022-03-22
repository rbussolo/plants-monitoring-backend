package device

import (
	"database/sql"

	d "backend/src/database"
)

type Device struct {
	Id     int
	UserId int
	Name   string
}

type Info struct {
	Id           int     `json:"id"`
	Device       string  `json:"device"`
	Timestamp    int64   `json:"timestamp"`
	SoilMoisture int     `json:"soilMoisture"`
	Temperature  float32 `json:"temperature"`
	Humidity     int     `json:"humidity"`
}

func NewDevice(userId int, name string) Device {
	db := d.GetInstance()

	var id int

	// Create a new user
	err := db.QueryRow("INSERT INTO devices(user_id, name) VALUES($1, $2) RETURNING id", userId, name).Scan(&id)
	if err != nil {
		panic(err)
	}

	device := Device{
		Id:     id,
		UserId: userId,
		Name:   name,
	}

	return device
}

func NewDeviceInformation(deviceId int, info Info) error {
	db := d.GetInstance()

	_, err := db.Exec("INSERT INTO device_info(device_id, timestamp, temperature, humidity, soil_moisture) VALUES($1, $2, $3, $4, $5)", deviceId, info.Timestamp, info.Temperature, info.Humidity, info.SoilMoisture)

	return err
}

func FindDeviceByName(userId int, name string) Device {
	db := d.GetInstance()

	var device Device
	err := db.Get(&device, "SELECT id, user_id as userId, name FROM devices WHERE name = $1", name)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return device
}

func FindDeviceInfo(userId int, name string) ([]Info, error) {
	db := d.GetInstance()

	var info []Info
	var err error
	var query string = `
		SELECT 
			i.id,
			d.name as device,
			i.timestamp,
			i.soil_moisture as soilMoisture,
			i.temperature,
			i.humidity 
		FROM devices d 
		INNER JOIN device_info i ON i.device_id = d.id
		WHERE d.user_id = $1
	`

	if name != "" {
		query += " AND d.name = $2"

		err = db.Select(&info, query, userId, name)
	} else {
		err = db.Select(&info, query, userId)
	}

	return info, err
}
