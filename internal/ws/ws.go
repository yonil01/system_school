package ws

import (
	"bytes"
	"foro-hotel/internal/logger"
	"io/ioutil"
	"net/http"
)

func CallApiRest(method, url string, jsonBytes []byte, header string, valueHeader string) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if header != "" {
		req.Header.Set(header, valueHeader)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error.Printf("no se pudo enviar la petición: %v  -- log: ", err)
		return resp.StatusCode, nil, err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Printf("no se pudo obtener response body: %v  -- log: ", err)
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, response, nil
}
