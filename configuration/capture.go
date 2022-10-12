// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package configuration

import (
	"errors"
	"fmt"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	parser "github.com/haproxytech/config-parser/v4"
	parser_errors "github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"

	"github.com/haproxytech/client-native/v5/models"
)

type Capture interface {
	GetDeclareCaptures(parentType string, parentName string, transactionID string) (int64, models.Captures, error)
	GetDeclareCapture(index int64, parentType string, parentName string, transactionID string) (int64, *models.Capture, error)
	DeleteDeclareCapture(index int64, parentType string, parentName string, transactionID string, version int64) error
	CreateDeclareCapture(parentType string, parentName string, data *models.Capture, transactionID string, version int64) error
	EditDeclareCapture(index int64, parentType string, parentName string, data *models.Capture, transactionID string, version int64) error
}

// GetDeclareCaptures returns configuration version and an array of configured DeclareCapture lines in the specified frontend.
// Returns error on fail.
func (c *client) GetDeclareCaptures(parentType string, parentName string, transactionID string) (int64, models.Captures, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}
	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}
	captures, err := ParseDeclareCaptures(parentType, parentName, p)
	if err != nil {
		return v, nil, c.HandleError("", parentType, parentName, "", false, err)
	}
	return v, captures, nil
}

// GetDeclareCapture returns configuration version and a requested DeclareCapture line in the specified frontend.
// Returns error on fail or if DeclareCapture does not exist
func (c *client) GetDeclareCapture(index int64, parentType string, parentName string, transactionID string) (int64, *models.Capture, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}

	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return 0, nil, NewConfError(ErrValidationError, "unsupported parent type")
	}

	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}

	data, err := p.GetOne(parentSection, parentName, "declare capture", int(index))
	if err != nil {
		return v, nil, c.HandleError(strconv.FormatInt(index, 10), parentType, parentName, "", false, err)
	}

	declareCapture := ParseDeclareCapture(data.(types.DeclareCapture))
	declareCapture.Index = &index
	return v, declareCapture, nil
}

// DeleteDeclareCapture deletes a DeclareCapture line in the configuration. One of version or transactionID is mandatory.
// Returns error on fail, nil on success
func (c *client) DeleteDeclareCapture(index int64, parentType string, parentName string, transactionID string, version int64) error {
	p, t, err := c.loadDataForChange(transactionID, version)
	if err != nil {
		return err
	}
	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return NewConfError(ErrValidationError, "unsupported parent type")
	}

	if err := p.Delete(parentSection, parentName, "declare capture", int(index)); err != nil {
		return c.HandleError(strconv.FormatInt(index, 10), parentType, parentName, t, transactionID == "", err)
	}
	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}
	return nil
}

// CreateDeclareCapture creates a DeclareCapture line in the configuration. One of version or transactionID is mandatory.
// Returns error on fail, nil on success
func (c *client) CreateDeclareCapture(parentType string, parentName string, data *models.Capture, transactionID string, version int64) error {
	if c.UseModelsValidation {
		validationErr := data.Validate(strfmt.Default)
		if validationErr != nil {
			return NewConfError(ErrValidationError, validationErr.Error())
		}
	}
	p, t, err := c.loadDataForChange(transactionID, version)
	if err != nil {
		return err
	}
	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return NewConfError(ErrValidationError, "unsupported parent type")
	}

	if err := p.Insert(parentSection, parentName, "declare capture", SerializeDeclareCapture(*data), int(*data.Index)); err != nil {
		return c.HandleError(strconv.FormatInt(*data.Index, 10), parentType, parentName, t, transactionID == "", err)
	}
	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}
	return nil
}

// EditDeclareCapture edits a DeclareCapture line in the configuration. One of version or transactionID is mandatory.
// Returns error on fail, nil on success.
func (c *client) EditDeclareCapture(index int64, parentType string, parentName string, data *models.Capture, transactionID string, version int64) error {
	if c.UseModelsValidation {
		validationErr := data.Validate(strfmt.Default)
		if validationErr != nil {
			return NewConfError(ErrValidationError, validationErr.Error())
		}
	}
	p, t, err := c.loadDataForChange(transactionID, version)
	if err != nil {
		return err
	}
	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return NewConfError(ErrValidationError, "unsupported parent type")
	}

	if _, err := p.GetOne(parentSection, parentName, "declare capture", int(index)); err != nil {
		return c.HandleError(strconv.FormatInt(index, 10), parentType, parentName, t, transactionID == "", err)
	}
	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}
	return nil
}

func ParseDeclareCaptures(parentType string, parentName string, p parser.Parser) (models.Captures, error) {
	captures := models.Captures{}
	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return nil, NewConfError(ErrValidationError, "unsupported parent type")
	}

	data, err := p.Get(parentSection, parentName, "declare capture", false)
	if err != nil {
		if errors.Is(err, parser_errors.ErrFetch) {
			return captures, nil
		}
		return nil, err
	}
	items, ok := data.([]types.DeclareCapture)
	if !ok {
		return captures, fmt.Errorf("type assert error []types.DeclareCapture")
	}
	for i, c := range items {
		index := int64(i)
		capture := ParseDeclareCapture(c)
		capture.Index = &index
		captures = append(captures, capture)
	}
	return captures, nil
}

func ParseDeclareCapture(f types.DeclareCapture) *models.Capture {
	return &models.Capture{
		Type:   f.Type,
		Length: f.Length,
	}
}

func SerializeDeclareCapture(f models.Capture) types.DeclareCapture {
	return types.DeclareCapture{
		Type:   f.Type,
		Length: f.Length,
	}
}
