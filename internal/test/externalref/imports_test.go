package externalref

import (
	"testing"

	"github.com/marvell/oapi-codegen/internal/test/externalref/packageA"
	"github.com/marvell/oapi-codegen/internal/test/externalref/packageB"
)

func TestParameters(t *testing.T) {
	b := &packageB.ObjectB{}
	_ = Container{
		ObjectA: &packageA.ObjectA{ObjectB: b},
		ObjectB: b,
	}
}
