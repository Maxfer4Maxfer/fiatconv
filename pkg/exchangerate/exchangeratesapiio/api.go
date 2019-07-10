package exchangeratesapiio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	errors "fiatconv/pkg/exchanging"
)

const addr string = "https://api.exchangeratesapi.io"

// RateAPI represents a wrapper for the "exchangeratesapi.io" service
type RateAPI struct {
	addr   string
	client getter
}

type getter interface {
	Get(url string) (resp *http.Response, err error)
}

// New returns a new RateAPI instance
func New(client getter) *RateAPI {
	return &RateAPI{
		client: client,
	}
}

type responce struct {
	Rates map[string]float32 `json:"rates"`
	Err   string             `json:"error"`
}

// Rate returns a exchange rate for src to destanation currtnce
func (api *RateAPI) Rate(src string, dst string) (float32, error) {

	src = strings.ToUpper(src)
	dst = strings.ToUpper(dst)

	// call to API
	url := fmt.Sprintf("%v/latest?base=%v&symbols=%v", addr, src, dst)

	resp, err := api.client.Get(url)
	if err != nil {
		return 0, errors.ErrRateUnavailable
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.ErrRateUnavailable
	}

	// parse the responce
	apiResp := responce{}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return 0, errors.ErrRateUnavailable
	}

	// handle errors
	if strings.Contains(apiResp.Err, "Base") {
		return 0, errors.ErrSrcCurrensyNotFound
	}

	if strings.Contains(apiResp.Err, "Symbols") {
		return 0, errors.ErrDstCurrensyNotFound
	}

	return apiResp.Rates[dst], nil
}
