package datasources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/config"
	"incorta-cloud/upload-service/utils"
	"incorta-cloud/upload-service/utils/errors"
	"io"
	"net/http"
	"strings"
)

var backendHttpClient *http.Client

type BackendClient struct {
	Context *gin.Context
}

func getClient() *http.Client {
	backendHttpClient = &http.Client{}
	return backendHttpClient
}

func (b BackendClient) addHeaders(r *http.Request) {

	val, exists := b.Context.Get(utils.UserToken)
	if exists {
		valStr := fmt.Sprintf("%v", val)
		r.Header.Add("X-API-TOKEN", valStr)
	}

	r.Header.Set("Content-Type", "application/json")
}

func parseQueryParams(queryParams map[string]string) string {
	var params []string

	for k, v := range queryParams {
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}

	q := "?" + strings.Join(params, "&")

	return q

}

func handleStatusCode(statusCode int, data string) (string, error) {
	if statusCode == http.StatusForbidden {
		return "", &errors.TokenExpiredError{}
	}

	if statusCode == http.StatusBadRequest {
		return "", &errors.ValidationError{}
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {

		return "", &errors.APIError{Message: data}
	}

	return data, nil
}

func (b BackendClient) Get(path string, queryParams map[string]string) (string, error) {

	queryStr := parseQueryParams(queryParams)

	r, err := http.NewRequest("GET", config.GetConfig().GetString("BACKEND_URL")+"/"+path+queryStr, nil)

	if err != nil {
		return "", nil
	}

	b.addHeaders(r)

	resp, err := getClient().Do(r)

	if err != nil {
		return "", err
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)

	return handleStatusCode(resp.StatusCode, buf.String())

}

func (b BackendClient) Post(path string, queryParams map[string]string, body interface{}) (string, error) {

	queryStr := parseQueryParams(queryParams)

	buff := bytes.Buffer{}
	enc := json.NewEncoder(&buff)

	enc.Encode(body)
	buffer := strings.NewReader(buff.String())

	r, err := http.NewRequest("POST", config.GetConfig().GetString("BACKEND_URL")+"/"+path+queryStr, buffer)

	if err != nil {
		return "", nil
	}

	b.addHeaders(r)

	resp, err := getClient().Do(r)

	if err != nil {
		return "", err
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	// check errors

	return handleStatusCode(resp.StatusCode, buf.String())
}
