package cluster

import (
	"context"
	"fmt"
)

// leaderElectorFactoryFunc is a LeaderElector factory method type.
type leaderElectorFactoryFunc func(context.Context, ...interface{}) (
	LeaderElector, error)

var leaderElectorFactories map[string]leaderElectorFactoryFunc

// RegisterLeaderElectorFactory will register a new LeaderElector factory
// method corresponding to the passed id.
func RegisterLeaderElectorFactory(id string, factory leaderElectorFactoryFunc) {
	if leaderElectorFactories == nil {
		leaderElectorFactories = make(
			map[string]leaderElectorFactoryFunc,
		)
	}

	leaderElectorFactories[id] = factory
}

// MakeLeaderElector will constuct a LeaderElector identified by id with the
// passed arguments.
func MakeLeaderElector(ctx context.Context, id string, args ...interface{}) (
	LeaderElector, error) {

	if _, ok := leaderElectorFactories[id]; !ok {
		return nil, fmt.Errorf("leader elector factory for '%v' "+
			"not found", id)
	}

	return leaderElectorFactories[id](ctx, args...)
}

var leaderElectorFactories map[string]leaderElectorFactoryFunc1

// RegisterLeaderElectorFactory will register a new LeaderElector factory
// method corresponding to the passed id.
func RegisterLeaderElectorFactory(id string, factory leaderElectorFactoryFunc1) {
	if leaderElectorFactories == nil {
		leaderElectorFactories = make(
			map[string]leaderElectorFactoryFunc,
		)
	}

	leaderElectorFactories[id1] = factory
}

