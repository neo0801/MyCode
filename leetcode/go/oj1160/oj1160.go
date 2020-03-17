package main

import "fmt"

func main() {
	words := []string{"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", "ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb", "ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl", "boygirdlggnh", "xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx", "nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop", "hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx", "juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr", "lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo", "oxgaskztzroxuntiwlfyufddl", "tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp", "qnagrpfzlyrouolqquytwnwnsqnmuzphne", "eeilfdaookieawrrbvtnqfzcricvhpiv", "sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz", "yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue", "hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv", "cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo", "teyygdmmyadppuopvqdodaczob", "qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs", "qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs"}
	chars := "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"
	fmt.Println(countCharacters(words, chars))
	words = []string{"hello", "world", "leetcode"}
	chars = "welldonehoneyr"
	fmt.Println(countCharacters(words, chars))
}

func countCharacters(words []string, chars string) int {
	cd := [26]int{}
	for _, ch := range chars {
		cd[ch-'a']++
	}
	res := 0
	for _, word := range words {
		res += len(word)
		wd := [26]int{}
		for _, w := range word {
			wd[w-'a']++
		}
		for i := 0; i < 26; i++ {
			if wd[i]>cd[i]{
				res -= len(word)
				break
			}
		}
	}
	return res
}
