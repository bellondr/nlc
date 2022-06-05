package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	s1, s2 = s2, s1
	s2Map :=make(map[string]int)
	s1Map :=make(map[string]int)
	for i := range s2 {
		val, ok := s2Map[string(s2[i])]
		if ok {
			s2Map[string(s2[i])] = val + 1
		} else {
			s2Map[string(s2[i])] = 1
		}
		s1Map[string(s2[i])] = 0
	}

	right := 0
	left := 0
	for right < len(s1) {
		val, ok := s1Map[string(s1[right])]
		if !ok {
			fmt.Println(string(s1[right]))
			for k := range s1Map {
				s1Map[k] = 0
			}
			left = right
			right++
			continue
		} else {
			s1Map[string(s1[right])] = val + 1
			right++
			if right - len(s2) >= left  {
				if containerSub(s1Map, s2Map) {
					return true
				}
				val, ok := s1Map[string(s1[right - len(s2)])]
				if ok {
					if val > 0 {
						s1Map[string(s1[right - len(s2)])] = val - 1
					}
				}
			}
		}

	}
	return false
}


func containerSub(s1Map map[string]int, s2Map map[string]int) bool {
	for k := range s1Map {
		val, ok := s2Map[k]
		if !ok {
			return false
		}
		if val != s1Map[k] {
			return false
		}
	}
	return true
}
