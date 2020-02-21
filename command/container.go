/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package command

import (
	"fmt"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/whiteblock/utility/utils"
)

// Container represents a docker container, this is calculated from the payload of the Run command
type Container struct {
	// BoundCpus are the cpus which the container will be set with an affinity for.
	BoundCPUs []int `json:"boundCPUs,omitonempty"`

	// EntryPoint overrides the docker containers entrypoint if non-empty
	EntryPoint string `json:"entrypoint"`

	// Environment represents the environment kv which will be provided to the container
	Environment map[string]string `json:"environment"`

	// Labels are any identifier which are to be attached to the container
	Labels map[string]string `json:"labels"`

	// Name is the unique name of the docker container
	Name string `json:"name"`

	// Network is the primary network(s) for this container to be attached to
	Network string `json:"network"`

	// Ports to be opened for each container, each port associated.
	Ports map[int]int `json:"ports"`

	// Volumes are the docker volumes to be mounted on this container
	Volumes []Mount `json:"volumes"`

	// Cpus should be a floating point value represented as a string, and
	// is  equivalent to the percentage of a single cores time which can be used
	// by a node. Can be more than 1.0, meaning the node can use multiple cores at
	// a time.
	Cpus string `json:"cpus"`

	// Memory supports values up to Terrabytes (tb). If the unit is omitted, then it
	// is assumed to be bytes. This is not case sensitive.
	Memory string `json:"memory"`

	// Image is the docker image
	Image string `json:"image"`

	// Args are the arguments passed to the containers entrypoint
	Args []string `json:"args"`

	// IP is the IP address on the primary network
	IP string `json:"ip"`

	// AutoRemove flags a container to be automatically removed
	AutoRemove bool `json"autoremove"`

	Credentials
}

// GetMemory gets the memory value as an integer.
func (c Container) GetMemory() (int64, error) {
	return utils.Memconv(c.Memory, utils.Mibi)
}

// GetEnv gets the environment variables in the format which is
// expected by docker
func (c Container) GetEnv() (envVars []string) {
	for key, val := range c.Environment {
		envVars = append(envVars, fmt.Sprintf("%s=%s", key, val))
	}
	return envVars
}

//  GetPortBindings gets the port bindings in the proper docker container types
func (c Container) GetPortBindings() (nat.PortSet, nat.PortMap, error) {
	if c.Ports == nil {
		return nil, nil, nil
	}
	dockerPorts := []string{}
	for hostPort, containerPort := range c.Ports {
		dockerPorts = append(dockerPorts, fmt.Sprintf("0.0.0.0:%d:%d/tcp", hostPort, containerPort))
	}
	portSet, portMap, err := nat.ParsePortSpecs(dockerPorts)
	return nat.PortSet(portSet), nat.PortMap(portMap), err
}

// GetEntryPoint returns the properly formatted
// entrypoint if this container has one,
// otherwise returns nil
func (c Container) GetEntryPoint() strslice.StrSlice {
	if len(c.EntryPoint) == 0 {
		return nil
	}
	return strslice.StrSlice([]string{c.EntryPoint})
}

//GetCmd gets the commands
func (c Container) GetCmd() strslice.StrSlice {
	return strslice.StrSlice(c.Args)
}

// GetMounts gets the docker mounts for the docker container
func (c Container) GetMounts() []mount.Mount {
	out := []mount.Mount{}
	for _, vol := range c.Volumes {
		out = append(out, mount.Mount{
			Type:     mount.TypeVolume,
			Source:   vol.Name,
			Target:   vol.Directory,
			ReadOnly: vol.ReadOnly,
		})
	}
	return out
}
