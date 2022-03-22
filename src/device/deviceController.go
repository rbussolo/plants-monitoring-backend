package device

import (
	"errors"
	"log"
)

func GetDevice(userId int, name string) (Device, error) {
	var device Device

	if userId == 0 {
		return device, errors.New("user is required")
	}

	if name == "" {
		return device, errors.New("name is required")
	}

	// Get device of this user
	device = FindDeviceByName(userId, name)

	if device.Id == 0 { // Create a new device if not exist
		device = NewDevice(userId, name)
	}

	return device, nil
}

func CreateNewInfo(userId int, info Info) error {
	// Get device
	device, err := GetDevice(userId, info.Device)

	log.Printf("Device %v", device)

	if err != nil {
		return err
	}

	// Insert a new information for this device
	err = NewDeviceInformation(device.Id, info)

	return err
}

func ListDeviceInfo(userId int, deviceName string) ([]Info, error) {
	var info []Info

	if userId == 0 {
		return info, errors.New("user is required")
	}

	info, err := FindDeviceInfo(userId, deviceName)

	return info, err
}
