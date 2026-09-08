package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Backend struct {
	Name    string
	Weight  int
	Counter int
}

// TODO (weighted-rr): implement per the lesson description.
type WeightedRoundRobin struct {
	Backends     []Backend
	TotalWeights int
}

func (r *WeightedRoundRobin) Pool(backends []Backend) string {
	var totalWeight int
	for _, b := range backends {
		totalWeight += b.Weight
	}

	r.Backends = backends
	r.TotalWeights = totalWeight
	return "OK"
}

func (r *WeightedRoundRobin) Pick() string {
	var highestIdx int

	for i, b := range r.Backends {
		r.Backends[i].Counter += b.Weight // step 1: add the weights to counter

		if r.Backends[i].Counter > r.Backends[highestIdx].Counter {
			highestIdx = i
		}
	}

	res := r.Backends[highestIdx].Name // step 2: use the highest counter

	r.Backends[highestIdx].Counter -= r.TotalWeights

	return res
}

func (r *WeightedRoundRobin) PickN(n int) string {
	results := make([]string, 0, n)
	for range n {
		results = append(results, r.Pick())
	}

	return strings.Join(results, ",")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	wrr := &WeightedRoundRobin{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		switch {
		case strings.Contains(line, "POOL "):
			cleanedLine := line[5:]
			backendStrings := strings.Split(cleanedLine, " ")
			backends := make([]Backend, 0, len(backendStrings))
			for _, b := range backendStrings {
				parts := strings.Split(b, ":")
				weight, _ := strconv.Atoi(parts[1])
				backend := Backend{
					Name:   parts[0],
					Weight: weight,
				}
				backends = append(backends, backend)
			}
			fmt.Println(wrr.Pool(backends))
		case strings.Contains(line, "PICKN "):
			n, _ := strconv.Atoi(line[6:])
			fmt.Println(wrr.PickN(n))
		case line == "PICK":
			fmt.Println(wrr.Pick())
		}
	}
}
