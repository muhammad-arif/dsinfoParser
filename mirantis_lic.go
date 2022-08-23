package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type LicenseType struct {
	Details struct {
		MaxEngines             int       `json:"maxEngines"`
		Expiration             time.Time `json:"expiration"`
		Tier                   string    `json:"tier"`
		LicenseType            string    `json:"licenseType"`
		ScanningEnabled        bool      `json:"scanningEnabled"`
		MirantisLicenseSubject string    `json:"mirantis_license_subject"`
	} `json:"details"`
	LicenseConfig struct {
		KeyId         string `json:"key_id"`
		PrivateKey    string `json:"private_key"`
		Authorization string `json:"authorization"`
		Jwtclaims     struct {
			Iss string   `json:"iss"`
			Sub string   `json:"sub"`
			Aud []string `json:"aud"`
			Azp string   `json:"azp"`
			Iat int      `json:"iat"`
			Exp int      `json:"exp"`
		} `json:"jwtclaims"`
	} `json:"license_config"`
	LastRefreshError string `json:"last_refresh_error"`
}

// NewLicenseType : Is a constructor for LicenseType
func NewLicenseType() *LicenseType {
	return &LicenseType{}
}

// GetAll to collect the all info as json
func (l LicenseType) GetAll(i *json.RawMessage) LicenseType {
	var data LicenseType
	err := json.Unmarshal(*i, &data)
	if err != nil {
		fmt.Errorf("[license.go.GetAll()] : cannot unmarshal %s", err)
	}
	return data
}
