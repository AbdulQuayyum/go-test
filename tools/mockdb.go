package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"AbdulQuayyum": {
		AuthToken: "123ABC",
		Username:  "AbdulQuayyum",
	},
	"Emmanuel": {
		AuthToken: "456DEF",
		Username:  "Emmanuel",
	},
	"Sulaiman": {
		AuthToken: "789GHI",
		Username:  "Sulaiman",
	},
}

var mockShawtiesDetails = map[string]CoinDetails{
	"AbdulQuayyum": {
		Amount:   1,
		Success:  true,
		Username: "AbdulQuayyum",
	},
	"Emmanuel": {
		Amount:   10,
		Success:  true,
		Username: "Emmanuel",
	},
	"Sulaiman": {
		Amount:   0,
		Success:  true,
		Username: "Sulaiman",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserAmount(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockShawtiesDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
