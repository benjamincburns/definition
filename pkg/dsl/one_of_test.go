/**
 * Copyright 2020 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testIps = map[string]string{
	"QUORUM1_SERVICE0":                         "172.24.0.5",
	"QUORUM1_SERVICE0_QUORUM_NETWORK":          "10.0.4.5",
	"QUORUM1_SERVICE0_SPLUNK_QUORUM1_SERVICE0": "172.24.0.6",
	"QUORUM2_SERVICE0":                         "172.24.4.5",
	"QUORUM2_SERVICE0_QUORUM_NETWORK":          "10.0.4.6",
	"QUORUM2_SERVICE0_SPLUNK_QUORUM2_SERVICE0": "172.24.4.6",
	"QUORUM3_SERVICE0":                         "172.24.8.5",
	"QUORUM3_SERVICE0_QUORUM_NETWORK":          "10.0.4.7",
	"QUORUM3_SERVICE0_SPLUNK_QUORUM3_SERVICE0": "172.24.8.6",
	"QUORUM4_SERVICE0":                         "172.24.12.5",
}

func TestHandleOneOf(t *testing.T) {
	res, err := HandleOneOf(testIps, "$_one_of(quorum1,quorum_network)")
	assert.NoError(t, err)
	assert.Equal(t, "10.0.4.5", res)

	_, err = HandleOneOf(testIps, "$_one_of(quorum5,quorum_network)")
	assert.Error(t, err)

	_, err = HandleOneOf(testIps, "$_one_of(quorum1,quorum_network2)")
	assert.Error(t, err)
}
