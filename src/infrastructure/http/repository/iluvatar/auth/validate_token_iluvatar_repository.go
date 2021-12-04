package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	client "github.com/pzentenoe/httpclient-call-go"
	"io/ioutil"
	"net/http"
)

type ValidateTokenResponse struct {
	Message string `json:"message"`
	Valid   bool   `json:"valid"`
}

func (r *authIluvatarRepository) ValidateToken(accessToken string) (bool, error) {
	headers := http.Header{
		echo.HeaderContentType: []string{echo.MIMEApplicationJSON},
		client.HeaderAuthorization: []string{accessToken},
	}

	response, err := r.httpClientCall.
		Path("/v1/auth/validate-token").
		Method(http.MethodPost).
		Headers(headers).
		Do()

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	if response.StatusCode != http.StatusOK {
		return false, errors.New(fmt.Sprintf("can't validate token, wrong http code from response: %d", response.StatusCode))
	}

	var validateTokenResponse *ValidateTokenResponse
	err = json.Unmarshal(data, &validateTokenResponse)
	if err != nil {
		return false, errors.New("can't unmarshal response from server")
	}

	return validateTokenResponse.Valid, nil
}
