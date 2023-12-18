package main

import (
	qrterminal "github.com/mdp/qrterminal/v3"
	"os"
)

func main() {
	//qrterminal.Generate("https://github.com/mdp/qrterminal", qrterminal.L, os.Stdout)
	//qrterminal.Generate("hello world", qrterminal.L, os.Stdout)

	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE, // Inverted
		WhiteChar: qrterminal.BLACK,
		QuietZone: qrterminal.QUIET_ZONE,
	}
	qrterminal.GenerateWithConfig("hello world", config)
}
