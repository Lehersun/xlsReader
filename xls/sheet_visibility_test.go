package xls

import (
	"testing"

	"github.com/shakinm/xlsReader/xls/record"
)

func TestSheetVisibilityHelpers(t *testing.T) {
	t.Run("zero value sheet defaults to visible", func(t *testing.T) {
		s := Sheet{}

		if s.HiddenState() != 0 {
			t.Fatalf("expected hidden state 0, got %d", s.HiddenState())
		}
		if s.IsHidden() {
			t.Fatal("expected zero-value sheet to be visible")
		}
	})

	t.Run("hidden and very hidden are hidden", func(t *testing.T) {
		sheets := []Sheet{
			{boundSheet: &record.BoundSheet{Grbit: [2]byte{0x01, 0x00}}},
			{boundSheet: &record.BoundSheet{Grbit: [2]byte{0x02, 0x00}}},
			{boundSheet: &record.BoundSheet{Grbit: [2]byte{0x03, 0x00}}},
		}

		for _, s := range sheets {
			if s.HiddenState() == 0 {
				t.Fatalf("expected hidden state to be non-zero, got %d", s.HiddenState())
			}
			if !s.IsHidden() {
				t.Fatal("expected sheet to be hidden")
			}
		}
	})

	t.Run("visible is not hidden", func(t *testing.T) {
		s := Sheet{boundSheet: &record.BoundSheet{Grbit: [2]byte{0x00, 0x00}}}

		if s.HiddenState() != 0 {
			t.Fatalf("expected hidden state 0, got %d", s.HiddenState())
		}
		if s.IsHidden() {
			t.Fatal("expected sheet to be visible")
		}
	})
}
