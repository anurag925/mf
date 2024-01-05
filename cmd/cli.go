package cmd

import (
	"flag"
	"log/slog"
)


type CLIDetails struct {
	Env *string
	Port *int
}

func ParseFlags() CLIDetails{
	slog.Info("Parsing cli flags ...")
	cliDetails := CLIDetails{}
	cliDetails.Env = flag.String("e", "development","Environment in which the application will work")
	cliDetails.Port = flag.Int("p", 1323,"Port on which the application will work")
	flag.Parse()
	slog.Info("Successfully parsed cli flags ...", slog.Any("flags", cliDetails))
	return cliDetails
}