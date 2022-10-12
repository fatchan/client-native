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
	goerrors "errors"
	"strconv"

	"github.com/go-openapi/strfmt"
	parser "github.com/haproxytech/config-parser/v4"
	parser_errors "github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"

	"github.com/haproxytech/client-native/v5/misc"
	"github.com/haproxytech/client-native/v5/models"
)

type BackendSwitchingRule interface {
	GetBackendSwitchingRules(parentType string, parentName string, transactionID string) (int64, models.BackendSwitchingRules, error)
	GetBackendSwitchingRule(id int64, parentType string, parentName string, transactionID string) (int64, *models.BackendSwitchingRule, error)
	DeleteBackendSwitchingRule(id int64, parentType string, parentName string, transactionID string, version int64) error
	CreateBackendSwitchingRule(parentType string, parentName string, data *models.BackendSwitchingRule, transactionID string, version int64) error
	EditBackendSwitchingRule(id int64, parentType string, parentName string, data *models.BackendSwitchingRule, transactionID string, version int64) error
}

// GetBackendSwitchingRules returns configuration version and an array of
// configured backend switching rules in the specified frontend. Returns error on fail.
func (c *client) GetBackendSwitchingRules(parentType string, parentName string, transactionID string) (int64, models.BackendSwitchingRules, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}

	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}

	bckRules, err := ParseBackendSwitchingRules(parentType, parentName, p)
	if err != nil {
		return v, nil, c.HandleError("", parentType, parentName, "", false, err)
	}

	return v, bckRules, nil
}

// GetBackendSwitchingRule returns configuration version and a requested backend switching rule
// in the specified frontend. Returns error on fail or if backend switching rule does not exist.
func (c *client) GetBackendSwitchingRule(id int64, parentType string, parentName string, transactionID string) (int64, *models.BackendSwitchingRule, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}

	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}

	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return 0, nil, NewConfError(ErrValidationError, "unsupported parent type")
	}
	data, err := p.GetOne(parentSection, parentName, "use_backend", int(id))
	if err != nil {
		return v, nil, c.HandleError(strconv.FormatInt(id, 10), parentType, parentName, "", false, err)
	}

	bckRule := ParseBackendSwitchingRule(data.(types.UseBackend))
	bckRule.Index = &id

	return v, bckRule, nil
}

// DeleteBackendSwitchingRule deletes a backend switching rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) DeleteBackendSwitchingRule(id int64, parentType string, parentName string, transactionID string, version int64) error {
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

	if err := p.Delete(parentSection, parentName, "use_backend", int(id)); err != nil {
		return c.HandleError(strconv.FormatInt(id, 10), parentType, parentName, t, transactionID == "", err)
	}

	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}
	return nil
}

// CreateBackendSwitchingRule creates a backend switching rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) CreateBackendSwitchingRule(parentType string, parentName string, data *models.BackendSwitchingRule, transactionID string, version int64) error {
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

	if err := p.Insert(parentSection, parentName, "use_backend", SerializeBackendSwitchingRule(*data), int(*data.Index)); err != nil {
		return c.HandleError(strconv.FormatInt(*data.Index, 10), parentType, parentName, t, transactionID == "", err)
	}

	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}

	return nil
}

// EditBackendSwitchingRule edits a backend switching rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) EditBackendSwitchingRule(id int64, parentType string, parentName string, data *models.BackendSwitchingRule, transactionID string, version int64) error {
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
	if _, err := p.GetOne(parentSection, parentName, "use_backend", int(id)); err != nil {
		return c.HandleError(strconv.FormatInt(id, 10), parentType, parentName, t, transactionID == "", err)
	}

	if err := p.Set(parentSection, parentName, "use_backend", SerializeBackendSwitchingRule(*data), int(id)); err != nil {
		return c.HandleError(strconv.FormatInt(id, 10), parentType, parentName, t, transactionID == "", err)
	}

	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}

	return nil
}

func ParseBackendSwitchingRules(parentType string, parentName string, p parser.Parser) (models.BackendSwitchingRules, error) {
	br := models.BackendSwitchingRules{}
	var parentSection parser.Section
	if parentType == "frontend" {
		parentSection = parser.Frontends
	} else {
		return nil, NewConfError(ErrValidationError, "unsupported parent type")
	}

	data, err := p.Get(parentSection, parentName, "use_backend", false)
	if err != nil {
		if goerrors.Is(err, parser_errors.ErrFetch) {
			return br, nil
		}
		return nil, err
	}

	bRules, ok := data.([]types.UseBackend)
	if !ok {
		return nil, misc.CreateTypeAssertError("[]types.DeclareCapture")
	}
	for i, bRule := range bRules {
		id := int64(i)
		b := ParseBackendSwitchingRule(bRule)
		if b != nil {
			b.Index = &id
			br = append(br, b)
		}
	}
	return br, nil
}

func ParseBackendSwitchingRule(ub types.UseBackend) *models.BackendSwitchingRule {
	return &models.BackendSwitchingRule{
		Name:     ub.Name,
		Cond:     ub.Cond,
		CondTest: ub.CondTest,
	}
}

func SerializeBackendSwitchingRule(bRule models.BackendSwitchingRule) types.UseBackend {
	return types.UseBackend{
		Name:     bRule.Name,
		Cond:     bRule.Cond,
		CondTest: bRule.CondTest,
	}
}
