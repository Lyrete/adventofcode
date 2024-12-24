package main

import (
	"aoc"
	"fmt"
	"maps"
	"slices"
	"strings"
)

func parse(input string) map[string][]string {
	ret := make(map[string][]string)
	for _, l := range strings.Split(input, "\n") {
		split := strings.Split(l, "-")
		if _, ok := ret[split[0]]; !ok {
			ret[split[0]] = []string{}
		}
		ret[split[0]] = append(ret[split[0]], split[1])
		if _, ok := ret[split[1]]; !ok {
			ret[split[1]] = []string{}
		}
		ret[split[1]] = append(ret[split[1]], split[0])
	}
	return ret
}

type lan struct {
	members map[string]struct{}
	network map[string][]string
}

func (l *lan) addNew(computer string) bool {
	for mem := range l.members {
		if !slices.Contains(l.network[computer], mem) {
			return false
		}
	}
	l.members[computer] = struct{}{}
	for _, connected := range l.network[computer] {
		l.addNew(connected)
	}
	return true
}

func (l *lan) getPassword() string {
	return strings.Join(slices.Sorted(maps.Keys(l.members)), ",")
}

func createPassowrdFromStringSlice(parts []string) string {
	slices.Sort(parts)
	return strings.Join(parts, ",")
}

func solve(input string) (int, string) {
	res, res2 := 0, ""
	instructions := parse(input)
	longest := 0
	foundSubnets := map[string]struct{}{}
	for k := range instructions {
		if k[0] == 't' {
			for i := 0; i < len(instructions[k]); i++ {
				for i2 := 0; i2 < len(instructions[k]); i2++ {
					if i2 == i {
						continue
					}
					if slices.Contains(instructions[instructions[k][i]], instructions[k][i2]) {
						pw := createPassowrdFromStringSlice([]string{k, instructions[k][i], instructions[k][i2]})
						if _, ok := foundSubnets[pw]; !ok {
							foundSubnets[pw] = struct{}{}
							res += 1
						}
					}
				}
			}
		}
		l := lan{members: map[string]struct{}{}, network: instructions}
		l.addNew(k)
		if len(l.members) > longest {
			longest = len(l.members)
			res2 = l.getPassword()
		}
	}

	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("23")))
}

const example = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`
