package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-user-investor/helper"
)

func GetAdminId(input AdminIdInput) (string, error) {
	adminID := helper.UserAdmin{}
	adminID.UnixAdmin = input.UnixID
	// fetch get /getAdminID from service api
	serviceAdmin := os.Getenv("SERVICE_ADMIN")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return adminID.UnixAdmin, errors.New("Service admin is empty")
	}
	resp, err := http.Get(serviceAdmin + "/api/v1/admin/getAdminID/" + adminID.UnixAdmin)
	if err != nil {
		return adminID.UnixAdmin, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return adminID.UnixAdmin, errors.New("Failed to get admin status or admin not found")
	}

	var response helper.AdminStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	} else if response.Data.StatusAccountAdmin == "deactive" {
		return "", errors.New("Admin account is deactive")
	} else if response.Data.StatusAccountAdmin == "active" {
		return adminID.UnixAdmin, nil
	} else {
		return "", errors.New("Invalid admin status")
	}
}