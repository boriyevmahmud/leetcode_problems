package main

import (
	"fmt"
)

func main() {
	s := "ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {

	if len(s) == 1 {
		return s
	}
	temp := ""
	subs := make(map[int]string)
	for i := 0; i < len(s); i++ {
		temp = string(s[i])
		for j := i + 1; j < len(s); j++ {
			temp1 := temp
			temp = temp + string(s[j])
			fmt.Println("temp: ", temp)
			fmt.Println("temp1: ", temp1)
			if !isPalindrome(temp) {
				if isPalindrome(temp1) {
					subs[len(temp1)] = temp1
				}
			} else {
				subs[len(temp)] = temp
			}
		}
	}
	max := 0
	for k := range subs {
		if max < k {
			max = k
		}
	}
	return subs[max]
}

func isPalindrome(s string) bool {
	reversed := reverseString(s)
	fmt.Println("reversed", reversed)

	return reversed == s
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
