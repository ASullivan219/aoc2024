package days

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DayOne struct {
	l               []int
	r               []int
	total           int
	similarityScore int
}

func NewDayOne() DayOne {
	file, err := os.Open("./inputs/dayOne.txt")
	if err != nil {
		slog.Error("Error opening day one input",
			"error", err,
		)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	l := make([]int, 0)
	r := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		l = append(l, left)
		r = append(r, right)
	}

	sort.Ints(l)
	sort.Ints(r)

	return DayOne{
		l:               l,
		r:               r,
		total:           0,
		similarityScore: 0,
	}
}

func (d *DayOne) Solve() {
	incidence_map := make(map[int]int)

	for i, _ := range d.l {
		distance := d.l[i] - d.r[i]
		if distance < 0 {
			distance *= -1
		}
		d.total += distance

		_, ok := incidence_map[d.r[i]]
		if ok {
			incidence_map[d.r[i]] += 1
		} else {
			incidence_map[d.r[i]] = 1
		}
	}

	for _, v := range d.l {
		occ, ok := incidence_map[v]
		if ok {
			d.similarityScore += v * occ
		}
	}

	fmt.Println("******** Day One ********")
	fmt.Printf("Part One: %d\n", d.total)
	fmt.Printf("Part Two: %d\n", d.similarityScore)
}
