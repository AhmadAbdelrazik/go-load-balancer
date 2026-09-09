package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// TODO (least-connections): implement per the lesson description.
//Input:
// POOL a b c
// PICK
// PICK
// PICK
// PICK
// DONE a
// PICK
// STATUS
//
// Output:
// OK
// a
// b
// c
// a
// OK
// a
// a:2
// b:1
// c:1

type Backend struct {
	Name        string
	Connections int
}

type LeastConnections struct {
	Backends []Backend
}

func (l *LeastConnections) Pool(backends []Backend) string {
	l.Backends = backends
	return "OK"
}

func (l *LeastConnections) Pick() string {
	var leastConnIdx int
	for i := range l.Backends {
		if l.Backends[i].Connections < l.Backends[leastConnIdx].Connections {
			leastConnIdx = i
		}
	}

	l.Backends[leastConnIdx].Connections++

	return l.Backends[leastConnIdx].Name
}

func (l *LeastConnections) Done(backend string) string {
	for i, b := range l.Backends {
		if b.Name != backend {
			continue
		}

		l.Backends[i].Connections = max(0, b.Connections-1)
		break
	}
	return "OK"
}

func (l *LeastConnections) Status() {
	sort.SliceStable(l.Backends, func(i, j int) bool {
		return l.Backends[i].Name < l.Backends[j].Name
	})

	for _, b := range l.Backends {
		fmt.Printf("%s:%d\n", b.Name, b.Connections)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	lc := &LeastConnections{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		switch {
		case strings.Contains(line, "POOL "):
			cleanLine := line[5:]
			backendStrs := strings.Split(cleanLine, " ")
			backends := make([]Backend, 0, len(backendStrs))
			for _, b := range backendStrs {
				backend := Backend{
					Name: b,
				}

				backends = append(backends, backend)
			}
			fmt.Println(lc.Pool(backends))
		case line == "PICK":
			fmt.Println(lc.Pick())
		case strings.Contains(line, "DONE "):
			backend := line[5:]
			fmt.Println(lc.Done(backend))
		case line == "STATUS":
			lc.Status()
		}
	}
}
