/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package settings

import (
	"fmt"
)

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

func (metrics Metrics) StrMap() map[string]string {
	out := map[string]string{}
	if metrics.StatsdHost != "" {
		out["statsdHost"] = metrics.StatsdHost
	} else {
		out["statsdHost"] = DefaultMetrics.StatsdHost
	}

	if metrics.SyslogNGProtocol != "" {
		out["syslogNGProtocol"] = metrics.SyslogNGProtocol
	} else {
		out["syslogNGProtocol"] = DefaultMetrics.SyslogNGProtocol
	}

	if metrics.SyslogNGIP != "" {
		out["syslogNGIP"] = metrics.SyslogNGIP
	} else {
		out["syslogNGIP"] = DefaultMetrics.SyslogNGIP
	}

	if metrics.SyslogNGPort != 0 {
		out["syslogNGPort"] = fmt.Sprint(metrics.SyslogNGPort)
	} else {
		out["syslogNGPort"] = fmt.Sprint(DefaultMetrics.SyslogNGPort)
	}
	return out
}
