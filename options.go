package dktest

import (
	"time"
)

import (
	"github.com/docker/go-connections/nat"
)

// Options contains the configurable options for running tests in the docker image
type Options struct {
	// PullTimeout is the timeout used when pulling images
	PullTimeout time.Duration
	// Timeout is the timeout used when starting a container and checking if it's ready
	Timeout time.Duration
	// CleanupTimeout is the timeout used when stopping and removing a container
	CleanupTimeout time.Duration
	ReadyFunc      func(ContainerInfo) bool
	Env            map[string]string
	Entrypoint     []string
	Cmd            []string
	// If you prefer to specify your port bindings as a string, use nat.ParsePortSpecs()
	PortBindings nat.PortMap
	PortRequired bool
	LogStdout    bool
	LogStderr    bool
}

func (o *Options) init() {
	if o.PullTimeout <= 0 {
		o.PullTimeout = DefaultPullTimeout
	}
	if o.Timeout <= 0 {
		o.Timeout = DefaultTimeout
	}
	if o.CleanupTimeout <= 0 {
		o.CleanupTimeout = DefaultCleanupTimeout
	}
}

func (o *Options) env() []string {
	env := make([]string, 0, len(o.Env))
	for k, v := range o.Env {
		env = append(env, k+"="+v)
	}
	return env
}
