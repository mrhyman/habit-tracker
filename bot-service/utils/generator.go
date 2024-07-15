package utils

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() uuid.UUID
	NewString() string
}

type RealUUIDGenerator struct{}

func (g RealUUIDGenerator) NewString() string {
	return uuid.NewString()
}

func (g RealUUIDGenerator) New() uuid.UUID {
	return uuid.New()
}

type FakeUUIDGenerator struct {
	FixedUUID string
}

func (g FakeUUIDGenerator) NewString() string {
	return g.FixedUUID
}

func (g FakeUUIDGenerator) New() uuid.UUID {
	u, _ := uuid.Parse(g.FixedUUID)
	return u
}
