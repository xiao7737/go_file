package numJewelsInStones

//返回S中存在J的个数，宝石与石头问题
//J：a   S ：aab   返回2

//暴力解法  O(N*N)
func numJewelsInStones(J string, S string) int {
	var i = 0
	for j := 0; j < len(J); j++ {
		for s := 0; s < len(S); s++ {
			if J[j] == S[s] {
				i++
			}
		}
	}
	return i
}

//利用map
func numJewelsInStones2(J string, S string) int {
	count := 0
	jmap := make(map[rune]int)
	for _, v := range J {
		jmap[v] = 1
	}
	for _, v := range S {
		if _, ok := jmap[v]; ok {
			count++
		}
	}
	return count
}
