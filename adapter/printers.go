package adapter

import "fmt"

type ModernPrinter interface {
    Print(msg string) string
}

// Service that the client can't directly implement
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintLegacy(msg string) string {
    response := fmt.Sprintf("Legacy Printer: %s", msg)
    return response
}

// Class that can work with both client and service
type PrinterAdapter struct {
    Legacy *LegacyPrinter // Service
}

func (pa *PrinterAdapter) Print(msg string) string {
    // Adapter translates the call
    legacyMsg := pa.Legacy.PrintLegacy(msg)
    return legacyMsg
}

func usePrinter(p ModernPrinter) string {
    return p.Print("Hello through the adapter!")
}

