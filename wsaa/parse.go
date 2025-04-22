package wsaa

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func parce(loginTicketResponse *string) (*model.WSAA, error) {
	// Parse XML response
	var response struct {
		Credentials struct {
			Token string `xml:"token"`
			Sign  string `xml:"sign"`
		} `xml:"credentials"`
		Header struct {
			ExpirationTime string `xml:"expirationTime"`
		} `xml:"header"`
	}

	err := xml.Unmarshal([]byte(*loginTicketResponse), &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing XML response: %v", err)
	}

	// Parse expiration time
	expirationTime, err := time.Parse("2006-01-02T15:04:05.999-07:00", response.Header.ExpirationTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing expiration time: %v", err)
	}

	// Create WSAA object with extracted data
	wsaa := &model.WSAA{
		Token:          response.Credentials.Token,
		Sign:           response.Credentials.Sign,
		ExpirationTime: expirationTime,
	}

	return wsaa, nil
}