/**
 * Copyright 2020 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dsl

func ExecuteDSL(ips map[string]string, subj string) (string, error) {
	return HandleOneOf(ips, subj)
}
