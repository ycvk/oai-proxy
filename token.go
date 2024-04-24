package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func parseJWT(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var claims map[string]any
	err = sonic.Unmarshal(payload, &claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func isTokenExpired(token string) bool {
	claims, err := parseJWT(token)
	if err != nil {
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	expirationTime := time.Unix(int64(exp), 0)
	return time.Now().After(expirationTime)
}

func refreshAccessToken(config *Config) error {
	var data map[string]any
	resp, err := req.R().
		SetFormData(map[string]string{"refresh_token": config.RefreshToken}).
		SetSuccessResult(&data).
		Post(refreshURL)

	if err != nil {
		return errors.New("error fetching access token")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		config.AccessToken = data["access_token"].(string)
		err := writeToConfigYml(config)
		if err != nil {
			log.Println("error writing to config.yml")
		}
		return nil
	} else {
		return errors.New("error fetching access token")
	}
}

func writeToConfigYml(config *Config) error {
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile("config.yml", out, 0644)
	return err
}