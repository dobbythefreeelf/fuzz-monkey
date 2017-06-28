package main

import (
  "encoding/json"
  "io/ioutil"
)

type Config struct {
  Endpoints []EndpointConfig `json:"endpoints"`
}

type EndpointConfig struct {
  Name string `json:"name"`
  Endpoint string `json:"endpoint"`
  MaxResponseTime int `json:"maxResponseTime"`
  MinResponseTime int `json:"minResponseTime"`
  SuccessCode int `json:"successCode"`
}

func GetConfig(configPath string) (*Config) {
  fileContents := loadConfigFile(configPath)
  return mapFileToObject(fileContents)
}

func loadConfigFile(configPath string) ([]byte) {
  file, err := ioutil.ReadFile(configPath)

  CheckError(err)

  return file
}

func mapFileToObject(contents []byte) (*Config) {
  config := &Config{}
  err := json.Unmarshal(contents, config)
  CheckError(err)
  return config
}
