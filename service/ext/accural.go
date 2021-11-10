package ext

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spinel/gophermart/model"
)

const PathOrderStatus = "api/orders"

// ExtAccuralService ...
type ExtAccuralService struct {
	ctx     context.Context
	address string
	client  *http.Client
}

// ExtAccuralService is an external accural service builder.
func NewOrderWebService(ctx context.Context, address string) *ExtAccuralService {
	client := &http.Client{Timeout: 10 * time.Second}
	orderService := &ExtAccuralService{
		ctx:     ctx,
		address: address,
		client:  client,
	}

	return orderService
}

// OrderStatus external status responseof accural system.
func (ext ExtAccuralService) OrderStatus(ctx context.Context, number int) (model.ExtOrder, error) {
	var order model.ExtOrder

	requestPath := fmt.Sprintf("%s/%d", PathOrderStatus, number)
	err := ext.getJSON(requestPath, &order)
	if err != nil {
		log.Printf("external api: %s", err)
	}

	return order, nil
}

func (ext ExtAccuralService) getJSON(path string, target interface{}) error {
	requestURL := fmt.Sprintf("%s/%s", ext.address, path)
	r, err := ext.client.Get(requestURL)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	if r.StatusCode != 200 {
		return errors.New("no Content")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return json.NewDecoder(r.Body).Decode(target)
}
