package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

type Input struct {
	line   string
	fields []string
}

func main() {
	var inputWg sync.WaitGroup
	var outputWg sync.WaitGroup
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	inputCh := make(chan *Input, n)
	outputCh := make(chan string, n)
	// spawn convert worker gorutines
	for i := 0; i < n; i++ {
		go func() {
			inputWg.Add(1)
			defer inputWg.Done()
			for {
				input, ok := <-inputCh
				if !ok {
					break
				}
				outputCh <- convert(input)
			}
		}()
	}
	// output goroutine
	go func() {
		outputWg.Add(1)
		defer outputWg.Done()
		for {
			output, ok := <-outputCh
			if !ok {
				break
			}
			fmt.Println(output)
		}
	}()

	var fields []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "#") == 0 {
			if strings.Index(line, "#Fields:") == 0 {
				fields = parseFields(line)
			}
			continue
		}
		inputCh <- &Input{line: line, fields: fields}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	close(inputCh)
	inputWg.Wait()
	close(outputCh)
	outputWg.Wait()
}

func convert(in *Input) string {
	cols := strings.SplitN(in.line, "\t", len(in.fields))
	for i, name := range in.fields {
		cols[i] = name + ":" + cols[i]
	}
	return strings.Join(cols, "\t")
}

func parseFields(line string) []string {
	f := strings.Split(line, " ")
	fields := make([]string, len(f)-1)
	for i, name := range f[1:] {
		name = strings.Replace(name, "-", "_", -1)
		name = strings.Replace(name, "(", "_", -1)
		name = strings.Replace(name, ")", "", -1)
		fields[i] = strings.ToLower(name)
	}
	return fields
}
