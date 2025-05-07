package adapter

import "testing"

func TestPrint(t *testing.T) {
    legacy := &LegacyPrinter{}
    adapter := &PrinterAdapter{Legacy: legacy}

    got := usePrinter(adapter)
    want := "Hello through the adapter!"

    if want != got {
	t.Errorf("want %s, got %s", want, got)
    }
}

