package gcs

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BucketName       string `json:"bucket_name"`
	BackupBucketName string `json:"backup_bucket_name"`
}

func ParseConfig(configFilePath string) (map[string]Config, error) {
	configContents, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config map[string]Config
	err = json.Unmarshal(configContents, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ReadGCPServiceAccountKey(gcpConfigFilePath string) (string, error) {
	gcpConfigContents, err := ioutil.ReadFile(gcpConfigFilePath)
	if err != nil {
		return "", err
	}

	return string(gcpConfigContents), nil
}
