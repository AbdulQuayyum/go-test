package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulQuayyum/go-test/api"
	"github.com/AbdulQuayyum/go-test/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetAmountBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.UserShawtiesAmountParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.AmountDetails
	tokenDetails = (*database).GetUserAmount(params.Username)
	if tokenDetails == nil {
		log.Error("User not found or no amount data")
		api.InternalErrorHandler(w)
		return
	}

	amount := tokenDetails.Amount
	message := ""
	if amount == 1 {
		message = fmt.Sprintf("%s has 1 shawtie", params.Username)
	} else {
		message = fmt.Sprintf("%s has %d shawties", params.Username, amount)
	}

	var response = api.UserShawtiesAmountResponse{
		Amount:  amount,
		Code:    http.StatusOK,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
