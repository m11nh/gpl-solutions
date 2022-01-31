package anagram

func IsAnagram(s1 string, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	c1, c2 := map[rune]int{}, map[rune]int{}
	for _, r := range r1 {
		c1[r]++
	}
	for _, r := range r2 {
		c2[r]++
	}
	for k, v1 := range c1 {
		if v1 != c2[k] {
			return false	
		}
	}
	for k, v1 := range c2 {
		if v1 != c1[k] {
			return false	
		}
	}
	return true
}
