package app

import (
	"encoding/json"
	"net/http"

	"github.com/Sonu875/goLearning/dto"
	"github.com/Sonu875/goLearning/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (s AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	} else {
		request.CustomerId = customerId
		account, err := s.service.NewAccount(request)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}
