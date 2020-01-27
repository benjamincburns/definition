/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package parser

import (
	"testing"

	"github.com/whiteblock/definition/pkg/entity"
	mockConverter "github.com/whiteblock/definition/pkg/mocks/converter"
	mockSearch "github.com/whiteblock/definition/pkg/mocks/search"
	"github.com/whiteblock/definition/schema"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/sirupsen/logrus"
)

var testSystemComp = schema.SystemComponent{
	Count: 5,
	Type:  "foo",
	Args:  nil,
	Environment: map[string]string{
		"bar": "baz",
	},
	Sidecars: nil,
	Resources: schema.SystemComponentResources{
		Cpus:     1,
		Memory:   "",
		Storage:  "20GB",
		Networks: nil,
	},
}

func TestResources_FromSystemDiff(t *testing.T) {
	testSystemComp2 := testSystemComp
	testSystemComp2.Count = testSystemComp.Count + 3

	searcher := new(mockSearch.Schema)
	searcher.On("FindServiceByType", mock.Anything, testSystemComp.Type).Return(
		schema.Service{}, nil).Once()

	conv := new(mockConverter.Resource)

	conv.On("FromResources", mock.Anything).Return(
		entity.Resource{}, nil).Times(int(testSystemComp2.Count))

	res := NewResources(searcher, conv, logrus.New())

	ents, err := res.FromSystemDiff(schema.RootSchema{}, testSystemComp, testSystemComp2)
	require.NoError(t, err)
	require.NotNil(t, ents)
	assert.Len(t, ents, 3)

	searcher.AssertExpectations(t)
	conv.AssertExpectations(t)
}

func TestResources_SystemComponent(t *testing.T) {
	searcher := new(mockSearch.Schema)
	searcher.On("FindServiceByType", mock.Anything, testSystemComp.Type).Return(
		schema.Service{}, nil).Once()

	conv := new(mockConverter.Resource)

	conv.On("FromResources", mock.Anything).Return(
		entity.Resource{}, nil).Times(int(testSystemComp.Count))

	res := NewResources(searcher, conv, logrus.New())

	ents, err := res.SystemComponent(schema.RootSchema{}, testSystemComp)
	assert.NoError(t, err)
	require.NotNil(t, ents)
	assert.Len(t, ents, int(testSystemComp.Count))

	searcher.AssertExpectations(t)
	conv.AssertExpectations(t)
}

func TestResources_SystemComponentNamesOnly(t *testing.T) {
	res := NewResources(nil, nil, logrus.New())
	ents := res.SystemComponentNamesOnly(testSystemComp)
	require.NotNil(t, ents)
	assert.Len(t, ents, int(testSystemComp.Count))
}

func TestResources_Tasks(t *testing.T) {
	testTasks := make([]schema.Task, 5)

	searcher := new(mockSearch.Schema)
	searcher.On("FindTaskRunnerByType", mock.Anything, mock.Anything).Return(
		schema.TaskRunner{}, nil).Times(len(testTasks))

	conv := new(mockConverter.Resource)
	conv.On("FromResources", mock.Anything).Return(
		entity.Resource{}, nil).Times(len(testTasks))

	res := NewResources(searcher, conv, logrus.New())

	result, err := res.Tasks(schema.RootSchema{}, testTasks)
	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result, len(testTasks))

	searcher.AssertExpectations(t)
	conv.AssertExpectations(t)
}
