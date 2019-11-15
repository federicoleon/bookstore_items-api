package controllers

import (
	"net/http"
	"github.com/federicoleon/bookstore_oauth-go/oauth"
	"github.com/federicoleon/bookstore_items-api/domain/items"
	"github.com/federicoleon/bookstore_items-api/services"
	"github.com/federicoleon/bookstore_items-api/utils/http_utils"
	"io/ioutil"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"encoding/json"
	"fmt"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}
