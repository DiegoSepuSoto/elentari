package auth

import (
	"elentari/src/domain/models"
	"elentari/src/domain/models/requests"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func (r *authIluvatarRepository) Login(loginRequest *requests.LoginRequest) (*models.Student, error) {
	headers := http.Header{
		echo.HeaderContentType: []string{echo.MIMEApplicationJSON},
	}

	response, err := r.httpClientCall.
		Path("/v1/auth/login").
		Method(http.MethodPost).
		Headers(headers).
		Body(loginRequest).
		Do()

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("can't login student, wrong http code from response: %d", response.StatusCode))
	}

	var loggedStudent *models.Student
	err = json.Unmarshal(data, &loggedStudent)
	if err != nil {
		return nil, errors.New("can't unmarshal response from server")
	}

	return loggedStudent, nil
}
