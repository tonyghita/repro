package repro_test

import (
	"testing"

	"github.com/tonyghita/repro"
)

func TestF(t *testing.T) {
	t.Run("F", func(t *testing.T) {
		var a *repro.A

		if a.F() != nil {
			t.Fail()
		}
	})
}
