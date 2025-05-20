package handlers

import (
	"encoding/json"
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
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.UserShawtiesAmountResponse{
		Amount: (*tokenDetails).Amount,
		Code:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
