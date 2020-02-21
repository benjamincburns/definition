/**
 * Copyright 2020 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dsl

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/whiteblock/definition/pkg/namer"
)

const (
	OneOfFnName = "$_one_of"
)

var (
	ErrNoMatches      = errors.New("no matches for the given service-network pair")
	ErrInvalidNumArgs = errors.New("$_one_of expects exactly 2 arguments (service,network)")
	OneOfRegex        = regexp.MustCompile(`\$_one_of\(([A-z|_|\-|0-9]*),([A-z|_|\-|0-9]*)\)`)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetEnvMatches(ips map[string]string, service, network string) ([]string, error) {
	pattern, err := regexp.Compile(namer.ToEnv(service) + "_SERVICE[0-9]+_" + namer.ToEnv(network))
	if err != nil {
		return nil, err
	}
	out := []string{}
	for key, ip := range ips {
		if pattern.MatchString(key) {
			out = append(out, ip)
		}
	}
	if len(out) == 0 {
		return nil, ErrNoMatches
	}
	return out, nil
}

func HandleOneOf(globalIPs map[string]string, input string) (string, error) {
	if !strings.Contains(input, OneOfFnName) { // most common path shouldn't use regex
		return input, nil
	}
	matches := OneOfRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) != 3 {
			return "", ErrInvalidNumArgs
		}
		service := match[1]
		network := match[2]
		ips, err := GetEnvMatches(globalIPs, service, network)
		if err != nil {
			return "", err
		}
		ip := ips[rand.Intn(len(ips))]
		input = strings.Replace(input, match[0], ip, -1)
	}
	return input, nil

}
