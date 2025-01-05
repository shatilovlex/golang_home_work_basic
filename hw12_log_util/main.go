package main

import (
	"fmt"
	"log"

	"github.com/shatilovlex/golang_home_work_basic/hw12_log_util/internal"
)

func main() {
	var err error
	inputFile, verbose, outputFile := internal.InitEnv()
	internal.InitFlags(&inputFile, &verbose, &outputFile)

	statistic := internal.NewStatistic(verbose)

	err = statistic.BuildStatisticByFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	if outputFile != "" {
		err = statistic.WriteToFile(outputFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(statistic)
	}
}
