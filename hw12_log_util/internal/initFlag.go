package internal

import "github.com/spf13/pflag"

func InitFlags(inputFile *string, verbose *string, outputFile *string) {
	pflag.StringVarP(inputFile, "file", "f", *inputFile, "path to file with logs")
	pflag.StringVarP(verbose, "level", "l", *verbose, "log level for analysis (info|debug)")
	pflag.StringVarP(outputFile, "output", "o", "", "file to which the statistics will be recorded")
	pflag.Parse()
}
