// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/clivern/walnut/core/migration"
	"github.com/clivern/walnut/core/model"
)

// CreateKeyBasedAuthData creates a new entity
func (db *Database) CreateKeyBasedAuthData(keyBasedAuthData *model.KeyBasedAuthData) *model.KeyBasedAuthData {
	db.Connection.Create(keyBasedAuthData)

	return keyBasedAuthData
}

// GetKeyBasedAuthDataByAPIKey gets an entity by uuid
func (db *Database) GetKeyBasedAuthDataByAPIKey(apiKey string) model.KeyBasedAuthData {
	keyBasedAuthData := model.KeyBasedAuthData{}
	db.Connection.Where("api_key = ?", apiKey).First(&keyBasedAuthData)

	return keyBasedAuthData
}

// DeleteKeyBasedAuthDataByID deletes an entity by id
func (db *Database) DeleteKeyBasedAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.KeyBasedAuthData{})
}
