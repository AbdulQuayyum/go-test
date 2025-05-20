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

var mockAmountDetails = map[string]AmountDetails{
	"AbdulQuayyum": {
		Amount:   1,
		Username: "AbdulQuayyum",
	},
	"Emmanuel": {
		Amount:   10,
		Username: "Emmanuel",
	},
	"Sulaiman": {
		Amount:   0,
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

func (d *mockDB) GetUserAmount(username string) *AmountDetails {
	time.Sleep(time.Second * 1)

	var clientData = AmountDetails{}
	clientData, ok := mockAmountDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
