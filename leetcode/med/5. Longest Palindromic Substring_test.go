package med

import (
	"fmt"
	"testing"
)

func Test_5(t *testing.T) {

	fmt.Println(longestPalindrome(bigText5()))
	/*if longestPalindrome2("cbbd") != "bb" {
		t.Error()
	}
	if longestPalindrome("ac") != "a" {
		t.Error()
	}
	if longestPalindrome("babad") != "bab" {
		t.Error()
	}*/
}

func longestPalindrome(s string) string {
	_max := ""

	if len(s) == 1 {
		return s
	}

	if len(s) > 1 {
		buffer := []byte(s)
		_buffer := []byte(reverse(s))

		n := len(buffer)
		for i := range n {
			for j := range i + 1 {
				sub := string(buffer[j : n-i+j])
				bus := string(_buffer[n-(n-i+j) : n-j])

				if len(sub) <= len(_max) {
					break
				}

				if sub == bus {
					if len(sub) > len(_max) {
						_max = sub
					}
				}
			}
		}
	}
	return _max
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func bigText5() string {
	return "ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy"
}
