package main

import (
	"os"

	"github.com/taylormonacelli/dailyclose"
)

func main() {
	code := dailyclose.Execute()
	os.Exit(code)
}
