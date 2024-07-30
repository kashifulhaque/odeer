package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"odeer/internal/config"
	"odeer/internal/models"
	"strings"
)

func SendRequest(config *config.Config, messages []models.Message) (string, error) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", config.AccountID, config.ModelName)

	payload, err := CreatePayload(messages)
	if err != nil {
		return "", fmt.Errorf("failed to create payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.AuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	return ProcessResponse(resp)
}

func CreatePayload(messages []models.Message) (string, error) {
	messagesJSON, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`{
		"stream": true,
		"messages": %s
	}`, string(messagesJSON)), nil
}

func ProcessResponse(resp *http.Response) (string, error) {
	scanner := bufio.NewScanner(resp.Body)
	var assistantResponse strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			jsonData := strings.TrimPrefix(line, "data: ")
			if strings.Contains(jsonData, `"response"`) {
				var responseData map[string]interface{}
				err := json.Unmarshal([]byte(jsonData), &responseData)
				if err != nil {
					return "", fmt.Errorf("error unmarshalling JSON: %w", err)
				}
				if response, ok := responseData["response"].(string); ok {
					assistantResponse.WriteString(response)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return assistantResponse.String(), nil
}
