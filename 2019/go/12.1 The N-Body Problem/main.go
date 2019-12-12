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

	//fmt.Println(LCM(4, 3, 2))

	lines := fileToLines("input.txt")
	original := LinesToMoons(lines)
	moons := LinesToMoons(lines)

	for i := 0; i < 10; i++ {
		moons = UpdateVelocity(moons)
		//fmt.Println(moons)
		moons = UpdatePosition(moons)
		//fmt.Println(moons)
	}
	e := SumEnergy(moons)
	fmt.Println(e)

	xRepeat := 0
	yRepeat := 0
	zRepeat := 0

	i := 2
	moons = LinesToMoons(lines)
	for xRepeat == 0 || yRepeat == 0 || zRepeat == 0 {

		moons = UpdateVelocity(moons)
		moons = UpdatePosition(moons)
		if xRepeat == 0 {
			xRepeat = IsBackAtStartX(original, moons, i)
		}
		if yRepeat == 0 {
			yRepeat = IsBackAtStartY(original, moons, i)
		}
		if zRepeat == 0 {
			zRepeat = IsBackAtStartZ(original, moons, i)
		}
		i++
	}

	fmt.Println(xRepeat, yRepeat, zRepeat)
	fmt.Println(LCM(int64(xRepeat), int64(yRepeat), int64(zRepeat)))

}

func LCM(in ...int64) int64 {
	if len(in) == 2 {
		return LCM2(in[0], in[1])
	}
	lcm := in[0]
	for i := 1; i < len(in); i++ {
		lcm = LCM2(lcm, in[i])
	}

	return lcm
}

func LCM2(a, b int64) int64 {

	gfA := factors(a)
	gfB := factors(b)
	cf := commonFactor(gfA, gfB)
	return (a * b) / cf
}
func commonFactor(a, b []int64) int64 {
	for _, af := range a {
		for _, bf := range b {
			if af == bf {
				return af
			}
		}
	}
	return 1
}
func factors(a int64) []int64 {
	ret := []int64{}
	for i := a; i > 0; i-- {
		if a%i == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
func IsBackAtStartX(original, moons []Moon, ret int) int {
	for i, o := range original {
		if o.P.X != moons[i].P.X {
			return 0
		}
	}
	return ret
}

func IsBackAtStartY(original, moons []Moon, ret int) int {
	for i, o := range original {
		if o.P.Y != moons[i].P.Y {
			return 0
		}
	}
	return ret
}
func IsBackAtStartZ(original, moons []Moon, ret int) int {
	for i, o := range original {
		if o.P.Z != moons[i].P.Z {
			return 0
		}
	}
	return ret
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
