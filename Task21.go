package main

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string) string
}

type ModernPrinter interface {
	PrintMessage(s string) string
}

type LegacyPrinterImpl struct{}

func (lp *LegacyPrinterImpl) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s", s)
}

type ModernPrinterImpl struct{}

func (mp *ModernPrinterImpl) PrintMessage(s string) string {
	return fmt.Sprintf("Modern Printer: %s", s)
}

type PrinterAdapter struct {
	ModernPrinter ModernPrinter
}

func (pa *PrinterAdapter) Print(s string) string {
	return pa.ModernPrinter.PrintMessage(s)
}

func main() {
	legacyPrinter := &LegacyPrinterImpl{}
	modernPrinter := &ModernPrinterImpl{}

	fmt.Println(legacyPrinter.Print("Hello, Legacy Printer!"))

	fmt.Println(modernPrinter.PrintMessage("Hello, Modern Printer!"))

	adapter := &PrinterAdapter{ModernPrinter: modernPrinter}
	fmt.Println(adapter.Print("Hello, Printer Adapter!"))
}
