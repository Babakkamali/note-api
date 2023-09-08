package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SendVerificationSMS(phoneNumber, token string) error {
	apiKey := os.Getenv("KAVENEGAR_API_KEY")
	template := "note-verify"
	
	endpoint := fmt.Sprintf("https://api.kavenegar.com/v1/%s/verify/lookup.json?receptor=%s&token=%s&template=%s", 
		apiKey, phoneNumber, token, template)

	response, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("failed to send SMS: %s", string(bodyBytes))
	}

	return nil
}