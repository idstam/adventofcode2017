package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Moon struct {
	P IntPoint
	V IntPoint
}

func main() {
	lines := fileToLines("input.txt")
	moons := LinesToMoons(lines)

	for i := 0; i < 1000; i++ {
		moons = UpdateVelocity(moons)
		//fmt.Println(moons)
		moons = UpdatePosition(moons)
		//fmt.Println(moons)
	}
	e := SumEnergy(moons)
	fmt.Println(e)

}
func SumEnergy(moons []Moon) int {
	ret := 0
	for _, m := range moons {
		p := IntAbs(m.P.X) + IntAbs(m.P.Y) + IntAbs(m.P.Z)
		k := IntAbs(m.V.X) + IntAbs(m.V.Y) + IntAbs(m.V.Z)
		ret += (p * k)
	}

	return ret
}
func UpdatePosition(moons []Moon) []Moon {
	ret := []Moon{}

	for _, o := range moons {
		m := o
		m.P.X += m.V.X
		m.P.Y += m.V.Y
		m.P.Z += m.V.Z
		ret = append(ret, m)
	}
	return ret

}

func UpdateVelocity(moons []Moon) []Moon {
	ret := []Moon{}

	for oi, o := range moons {
		m := o
		for ii, i := range moons {
			if oi == ii {
				continue
			}
			m.V.X += VelocityChange(m.P.X, i.P.X)
			m.V.Y += VelocityChange(m.P.Y, i.P.Y)
			m.V.Z += VelocityChange(m.P.Z, i.P.Z)
		}
		ret = append(ret, m)
	}
	return ret

}
func VelocityChange(a, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}
func LinesToMoons(lines []string) []Moon {
	ret := []Moon{}

	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "<", "")
		line = strings.ReplaceAll(line, ">", "")
		pairs := strings.Split(line, ",")

		p := strings.Split(pairs[0], "=")
		m := Moon{}
		m.P.X, _ = strconv.Atoi(p[1])

		p = strings.Split(pairs[1], "=")
		m.P.Y, _ = strconv.Atoi(p[1])

		p = strings.Split(pairs[2], "=")
		m.P.Z, _ = strconv.Atoi(p[1])
		ret = append(ret, m)
	}
	return ret
}
