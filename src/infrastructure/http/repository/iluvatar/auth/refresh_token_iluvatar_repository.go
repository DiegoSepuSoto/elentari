package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	client "github.com/pzentenoe/httpclient-call-go"
	"io/ioutil"
	"net/http"
)

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

func (r *authIluvatarRepository) RefreshToken(accessToken string) (string, error) {
	headers := http.Header{
		echo.HeaderContentType:     []string{echo.MIMEApplicationJSON},
		client.HeaderAuthorization: []string{accessToken},
	}

	response, err := r.httpClientCall.
		Path("/v1/auth/refresh-token").
		Method(http.MethodPost).
		Headers(headers).
		Do()

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("can't refresh token, wrong http code from response: %d", response.StatusCode))
	}

	var refreshTokenResponse *RefreshTokenResponse
	err = json.Unmarshal(data, &refreshTokenResponse)
	if err != nil {
		return "", errors.New("can't unmarshal response from server")
	}

	return refreshTokenResponse.Token, nil
}
