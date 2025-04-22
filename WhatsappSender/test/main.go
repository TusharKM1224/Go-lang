package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	whatsappAPIURL = "https://graph.facebook.com/v22.0/<PhoneID>/messages"
	accessToken    = "<Access Token>"
)

func sendPropertyInquiry(toPhoneNumber string, inquiryData map[string]string) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+ accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"messaging_product": "whatsapp",
			"recipient_type":    "individual",
			"to":                toPhoneNumber,
			"type":              "template",
			"template": map[string]interface{}{
				"name": "<template_name>", // Replace with your approved template name
				"language": map[string]string{
					"code": "en", // or "en_US" depending on your template
				},
				"components": []map[string]interface{}{
					{
						"type": "body",
						"parameters": []map[string]string{
							{"type": "text", "text": inquiryData["name"]},
							{"type": "text", "text": inquiryData["contact"]},
							{"type": "text", "text": inquiryData["email"]},
							{"type": "text", "text": inquiryData["budget"]},
							{"type": "text", "text": inquiryData["location"]},
							{"type": "text", "text": inquiryData["bedroom"]},
							{"type": "text", "text": inquiryData["amenities"]},
							{"type": "text", "text": inquiryData["date"]},
						},
					},
				},
			},
		}).
		Post(whatsappAPIURL)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status())
	fmt.Println("Response Body:", resp.String())
}

func main() {
	inquiryData := map[string]string{
		"name":      "John Doe",
		"contact":   "9000000000", // Should include country code if not in template
		"email":     "john@example.com",
		"budget":    "₹50L - ₹70L",
		"location":  "Bangalore",
		"bedroom":   "3 BHK",
		"amenities": "Gym, Pool, Parking",
		"date":      "2025-04-10",
	}

	sendPropertyInquiry("<recieverPhone>", inquiryData)
}
