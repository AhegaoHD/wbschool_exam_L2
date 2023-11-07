package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type GrepConfig struct {
	after        int
	before       int
	context      int
	count        bool
	ignoreCase   bool
	invertMatch  bool
	fixedStrings bool
	lineNum      bool
	pattern      string
	fileName     string
}

func parseFlags() *GrepConfig {
	cfg := &GrepConfig{}

	flag.IntVar(&cfg.after, "A", 0, "print +N lines after match")
	flag.IntVar(&cfg.before, "B", 0, "print +N lines before match")
	flag.IntVar(&cfg.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&cfg.count, "c", false, "print the count of matching lines")
	flag.BoolVar(&cfg.ignoreCase, "i", false, "ignore case distinctions")
	flag.BoolVar(&cfg.invertMatch, "v", false, "select non-matching lines")
	flag.BoolVar(&cfg.fixedStrings, "F", false, "interpret pattern as a fixed string")
	flag.BoolVar(&cfg.lineNum, "n", false, "print line number with output lines")

	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "usage: mygrep [flags] pattern file")
		os.Exit(1)
	}

	cfg.pattern = flag.Arg(0)
	cfg.fileName = flag.Arg(1)

	return cfg
}

func grep(cfg *GrepConfig, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	lineNumber := 0
	matchedLines := 0

	// Build the regular expression
	pattern := cfg.pattern
	if cfg.ignoreCase {
		pattern = "(?i)" + pattern
	}

	var re *regexp.Regexp
	var err error
	if cfg.fixedStrings {
		re, err = regexp.Compile(regexp.QuoteMeta(pattern))
	} else {
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		return fmt.Errorf("error compiling pattern: %w", err)
	}

	var beforeLines []string
	afterCount := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		match := re.MatchString(line)
		if cfg.invertMatch {
			match = !match
		}

		if cfg.context > 0 {
			cfg.after = cfg.context
			cfg.before = cfg.context
		}

		if match {
			if cfg.before > 0 && len(beforeLines) > 0 {
				fmt.Fprintln(output, strings.Join(beforeLines, "\n"))
				beforeLines = nil
			}
			if cfg.after > 0 {
				afterCount = cfg.after
			}
		}

		if match || afterCount > 0 {
			if cfg.count {
				matchedLines++
			} else {
				if cfg.lineNum {
					fmt.Fprintf(output, "%d:", lineNumber)
				}
				fmt.Fprintln(output, line)
			}
			if afterCount > 0 {
				afterCount--
			}
		}

		if cfg.before > 0 {
			if len(beforeLines) >= cfg.before {
				beforeLines = beforeLines[1:]
			}
			beforeLines = append(beforeLines, line)
		}

		if !match && afterCount > 0 {
			afterCount = 0
			fmt.Fprintln(output, "--")
		}
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error reading input: %w", scanner.Err())
	}

	if cfg.count {
		fmt.Fprintln(output, matchedLines)
	}

	return nil
}

func main() {
	cfg := parseFlags()

	file, err := os.Open(cfg.fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := grep(cfg, file, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
