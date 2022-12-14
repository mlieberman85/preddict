//
// Copyright 2022 The in-toto predicate dictionary Authors
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

package model

import (
	"github.com/mitchellh/mapstructure"
)

type Mapping struct {
	predicateToMap string
	indexes        []string
	mappings       map[string]string
}

type ModelWithIndexes[T any] struct {
	model   T
	indexes []string
}

type Document map[string]interface{}

type Converter[T any] interface {
	// Convert takes an arbitrary map and converts it via the given mapping
	Convert(d map[string]interface{}, m *Mapping) ModelWithIndexes[T]
}

func ConvertAny[T any](d map[string]interface{}, m *Mapping) (*ModelWithIndexes[T], error) {
	mapped := make(map[string]interface{})
	for k, v := range m.mappings {
		mapped[k] = d[v]
	}

	model := ModelWithIndexes[T]{
		model:   *new(T),
		indexes: m.indexes,
	}

	err := mapstructure.Decode(mapped, &model.model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
