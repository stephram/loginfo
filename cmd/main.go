package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"os"

	log "github.com/sirupsen/logrus"
)

var (
	filename *string
	verbose  *bool
)

func init() {
	f := false
	verbose = &f
}

func main() {
	configureLogging()

	filename = flag.String("f", "./programming-task-example-data.log", "web log filename")
	verbose = flag.Bool("v", false, "verbose output")
	flag.Parse()

	if err := validateFilename(*filename); err != nil {
		log.WithError(err).Error("terminating")
		os.Exit(1)
	}
	log.Infof("parsing '%s'", *filename)

	ipaddrs := make(map[string]int)
	webpags := make(map[string]int)

	count, err := processLogFile(*filename, &ipaddrs, &webpags)
	if err != nil {
		log.WithError(err).Errorf("failed to process log file '%s'", *filename)
		os.Exit(1)
	}

	if *verbose {
		displayResults(&ipaddrs, &webpags)
	}
	fmt.Print("\n")

	log.Infof("Processed %d log entries", count)
	fmt.Printf("unique IPs        : %d\n", len(ipaddrs))
	fmt.Printf("most visited URLs : %s\n", getTopmost(&webpags, 3))
	fmt.Printf("most active IPs   : %s\n", getTopmost(&ipaddrs, 3))

	os.Exit(0)
}

func displayResults(ipaddrs *map[string]int, webpags *map[string]int) {
	fmt.Printf("\n")

	fmt.Printf("IP address usage\n")
	if ipaddrs != nil {
		for k, v := range *ipaddrs {
			fmt.Printf("%d : %s\n", v, k)
		}
		fmt.Print("\n")
	}

	fmt.Printf("URL usage\n")
	if webpags != nil {
		for k, v := range *webpags {
			fmt.Printf("%d : %s\n", v, k)
		}
		fmt.Print("\n")
	}
}

func processLogFile(filename string, ipaddrs *map[string]int, webpags *map[string]int) (int, error) {
	f, err := os.Open(filename) // nolint: gosec
	if err != nil {
		return 0, errors.Wrapf(err, "failed to open file '%s'", filename)
	}
	defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	fmt.Printf("%s\n", scanner.Text())
	// }

	count := 0
	fs := bufio.NewReader(f)

	for {
		c, err := processEntry(fs, ipaddrs, webpags)
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
		count += c
	}
	return count, nil
}

func processEntry(reader *bufio.Reader, ipaddrs *map[string]int, webpags *map[string]int) (int, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	s = strings.TrimSpace(s)
	if len(s) <= 0 {
		return 0, nil
	}
	t := strings.Split(s, " ")
	if len(t) < 7 {
		return 0, nil
	}
	if *verbose {
		fmt.Printf("%s, %s\n", t[0], t[6])
	}
	// What to do if only one accumulate call fails?
	if err = accumulate(ipaddrs, t[0]); err != nil {
		return 0, err
	}
	if err = accumulate(webpags, t[6]); err != nil {
		return 0, err
	}
	return 1, nil
}

func accumulate(accum *map[string]int, name string) error {
	if accum == nil {
		return errors.New("invalid accum argument")
	}

	if val, ok := (*accum)[name]; ok {
		(*accum)[name] = val + 1
		return nil
	}
	(*accum)[name] = 1
	return nil
}

func getTopmost(values *map[string]int, topmost int) []string {
	type sortableItem struct {
		key string
		val int
	}
	var entries []sortableItem

	for k, v := range *values {
		entries = append(entries, sortableItem{k, v})
	}
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].val > entries[j].val
	})
	var results []string
	for i := 0; i < topmost; i++ {
		results = append(results, entries[i].key)
	}
	return results
}

func validateFilename(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to stat filename '%s'", filename))
	}
	if fi.IsDir() {
		return fmt.Errorf("the supplied filename '%s' is a directory", filename)
	}
	return nil
}

func configureLogging() {
	log.SetReportCaller(false)
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: false})
	log.SetLevel(log.InfoLevel)
}
