package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

type Statistic struct {
	CountRows  int
	StartTime  time.Time
	EndTime    time.Time
	CountError int
	verbose    string
	rows       []string
}

func NewStatistic(verbose string) Statistic {
	statistic := Statistic{}
	statistic.logLevel(verbose)
	return statistic
}

func (s Statistic) String() string {
	switch s.verbose {
	case "debug":
		return fmt.Sprintf(
			"rows:%s\n\nResult analyze:\n%d rows analyzed by period [%s - %s]\nFound %d errors",
			s.rows,
			s.CountRows,
			s.StartTime,
			s.EndTime,
			s.CountError,
		)
	default:
		return fmt.Sprintf(
			"Result analyze:\n%d rows analyzed by period [%s - %s]\nFound %d errors",
			s.CountRows,
			s.StartTime,
			s.EndTime,
			s.CountError,
		)
	}
}

func (s *Statistic) logLevel(verbose string) {
	s.verbose = verbose
}

func (s *Statistic) BuildStatisticByFile(inputFile string) error {
	var file *os.File
	var row string
	var err error
	file, err = os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	row, err = br.ReadString(byte('\n'))
	for err == nil {
		err = s.fillStatisticByRow(row)
		if err != nil {
			return err
		}
		row, err = br.ReadString(byte('\n'))
	}
	if errors.Is(err, io.EOF) {
		err = s.fillStatisticByRow(row)
		return err
	}

	return nil
}

func (s *Statistic) fillStatisticByRow(row string) error {
	s.CountRows++
	if s.verbose == "debug" {
		s.rows = append(s.rows, row)
	}

	logsFormat := `\[$timestamp\] rabbit_queues.$level: $_`
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)
	re := regexp.MustCompile(regexFormat)
	matches := re.FindStringSubmatch(row)

	for i, k := range re.SubexpNames() {
		if i == 0 || k == "_" {
			continue
		}

		switch k {
		case "timestamp":
			t, err := time.Parse(time.RFC3339, matches[i])
			if err != nil {
				return fmt.Errorf("cant parse datetime: %w", err)
			}
			if (s.StartTime == time.Time{}) {
				s.StartTime = t
			}
			s.EndTime = t
		case "level":
			if matches[i] == "ERROR" {
				s.CountError++
			}
		}
	}
	return nil
}

func (s *Statistic) WriteToFile(outputFile string) error {
	var file *os.File
	var written int
	var err error
	file, err = os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644) //nolint:gofumpt
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer file.Close()

	written, err = file.Write([]byte(s.String()))
	if err != nil {
		return fmt.Errorf("unable to write in file: %w", err)
	}

	fmt.Printf("Write %d bytes of data\n", written)

	return nil
}
