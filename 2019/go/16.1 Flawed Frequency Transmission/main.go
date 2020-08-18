package main

import (
	"fmt"
)

func main() {
	input := "59773775883217736423759431065821647306353420453506194937477909478357279250527959717515453593953697526882172199401149893789300695782513381578519607082690498042922082853622468730359031443128253024761190886248093463388723595869794965934425375464188783095560858698959059899665146868388800302242666479679987279787144346712262803568907779621609528347260035619258134850854741360515089631116920328622677237196915348865412336221562817250057035898092525020837239100456855355496177944747496249192354750965666437121797601987523473707071258599440525572142300549600825381432815592726865051526418740875442413571535945830954724825314675166862626566783107780527347044"
	//input := "80871224585914546619083218645595"
	digits := StringToSlice(input)
	nums := StringToIntArray(digits)
	nums = IntRepeatArray(nums, 10000)
	pattern := []int{0, 1, 0, -1}

	nextPhase := nums
	for i := 1; i <= 100; i++ {
		nextPhase = CalcPhase(nextPhase, pattern)

	}
	//fmt.Print(nextPhase)
	fmt.Println("---")
	for i := 5977377; i <= 5977377+8; i++ {
		fmt.Print(nextPhase[i])
	}
	fmt.Println("---")

}

func CalcPhase(nums, pattern []int) []int {
	//ret := []int{}
	ret := make([]int, len(nums))
	for j := range nums {
		line := 0
		//phasePattern := GetPattern(nums, pattern, j+1)
		//fmt.Println(phasePattern)
		for i := range nums {
			v := nums[i] * GetPatternValue(i, pattern, j+1)
			line += v

		}

		ret[j] = IntAbs(line % 10)
	}
	return ret
}
func GetPatternValue(num int, pattern []int, phase int) int {
	i := 0
	skipp := true
	l := 0
	for l <= num {
		for ph := 0; ph < phase; ph++ {
			if skipp {
				skipp = false
				continue
			}
			p, _ := IntRingAt(i, pattern)
			if l == num {
				return p
			}
			l++
		}
		i++
	}
	return -1
}

func GetPattern(nums []int, pattern []int, phase int) []int {
	ret := []int{}
	i := 0
	skipp := true
	for len(ret) <= len(nums) {
		for ph := 0; ph < phase; ph++ {
			if skipp {
				skipp = false
				continue
			}
			p, _ := IntRingAt(i, pattern)
			ret = append(ret, p)

		}
		i++
	}
	return ret
}
