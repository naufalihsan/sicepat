package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"moul.io/http2curl"
	"net/http"
	"reflect"
	"time"
)

func Call(header *http.Header, body interface{}, result interface{}) *error {
	reqBody := []byte("")
	var err error
	var client = &http.Client{}

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())
	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	url := "https://partner.sicepatklik.com/sicepat/api/v1/index.php"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	if header != nil {
		request.Header = *header
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = "ARTAKA_DEV"
	claims["iat"] = time.Now().Unix()

	jwt, _ := token.SignedString([]byte("art4k4d3v"))

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	request.Header.Set("Content-Type", "application/json")

	command, _ := http2curl.GetCurlCommand(request)
	fmt.Println(command)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		panic(err)
	}

	fmt.Println(string(respBody))

	if err := json.Unmarshal(respBody, &result); err != nil {
		panic(err)
	}

	return nil

}
