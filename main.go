package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RoundRobin struct {
	Backends []string
	Pointer  int
}

func (r *RoundRobin) Pool(backends []string) string {
	r.Backends = backends
	r.Pointer = 0
	return "OK"
}

func (r *RoundRobin) Pick() string {
	n := len(r.Backends)
	if n == 0 {
		return "EMPTY"
	}
	res := r.Backends[r.Pointer%n]
	r.Pointer++
	return res
}

func (r *RoundRobin) Reset() string {
	r.Pointer = 0
	return "OK"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	rr := &RoundRobin{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		var res string
		switch {
		case strings.HasPrefix(line, "POOL "):
			cleanLine, _ := strings.CutPrefix(line, "POOL ")
			backends := strings.Split(cleanLine, " ")
			res = rr.Pool(backends)
		case line == "PICK":
			res = rr.Pick()
		case line == "RESET":
			res = rr.Reset()
		}
		fmt.Println(res)
	}
}
