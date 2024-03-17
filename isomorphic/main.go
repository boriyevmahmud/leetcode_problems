package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("egg", "add"))

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		map1 = make(map[byte]byte)
		map2 = make(map[byte]byte)
	)

	for i := 0; i < len(s); i++ {
		s1 := s[i]
		s2 := t[i]

		if _, ok := map1[s1]; !ok {
			map1[s1] = s2
		} else if map1[s1] != s2 {
			return false
		}

		if _, ok := map2[s2]; !ok {
			map2[s2] = s1
		} else if map2[s2] != s1 {
			return false
		}

	}

	return true
}
