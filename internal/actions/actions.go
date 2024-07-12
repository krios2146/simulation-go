package actions

import "github.com/krios2146/simulation-go/internal/world"

type Action interface {
	Perform(m *world.Map)
}
