// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"time"
)

// KeyBasedAuthData struct
type KeyBasedAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name" validate:"required"`
	APIKey       string `json:"apiKey" validate:"required"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// KeyBasedAuthDataItems struct
type KeyBasedAuthDataItems struct {
	KeyBasedAuthDataItems []KeyBasedAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthData) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthData) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&k)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthDataItems) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthDataItems) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&k)

	if err != nil {
		return "", err
	}

	return string(data), nil
}