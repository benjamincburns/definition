/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package settings

type Metrics struct {
	StatsdHost       string `json:"statsdHost"`
	SyslogNGProtocol string `json:"syslogNGProtocol"`
	SyslogNGIP       string `json:"syslogNGIP"`
	SyslogNGPort     int    `json:"syslogNGPort"`
}

var DefaultMetrics = Metrics{
	StatsdHost:       "localhost:8086",
	SyslogNGProtocol: "tcp",
	SyslogNGIP:       "localhost",
	SyslogNGPort:     514,
}
