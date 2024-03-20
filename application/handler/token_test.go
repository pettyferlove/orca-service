package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func BenchmarkToken_Create(b *testing.B) {
	loginInfo := LoginRequest{
		Username: "administrator",
		Password: "123456",
	}
	jsonData, _ := json.Marshal(loginInfo)

	for i := 0; i < b.N; i++ {
		func() {
			req, _ := http.NewRequest("POST", "http://127.0.0.1:4001/api/v1/tokens", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				b.Fatal(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					b.Fatal(err)
				}
			}(resp.Body)
		}()
	}
}
