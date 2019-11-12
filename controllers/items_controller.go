package controllers

import (
	"net/http"
	"github.com/federicoleon/bookstore_oauth-go/oauth"
	"github.com/federicoleon/bookstore_items-api/domain/items"
	"github.com/federicoleon/bookstore_items-api/services"
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//Return error json to the user.
		return
	}

	//TODO: Unmarshal request into the item struct.
	item := items.Item{
		Seller: oauth.GetClientId(r),
	}

	_, err := services.ItemsService.Create(item)
	if err != nil {
		//Return error json to the user.
		return
	}

	//TODO: Return created item with HTTP status 201
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}
