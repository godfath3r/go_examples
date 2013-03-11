package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

var letters []string = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "/", "|", "[", "]"}

var digits int = 4 
var word []int = make([]int, digits)

var br bool = false
var x string = getmd5("A8df")

var counts int = 0
func main() {
	cable := make(chan []int)
	
	getmd5("hell")
	getmd5("help")
	fmt.Println(x)
	count(digits-1, cable)
}

func check(word []int) {
	word_to_check := decode(word)
	if  x == getmd5(word_to_check){
		fmt.Println("FOUND! It's:", word_to_check)
		br = true
	}
}

func decode(wordnum []int) string{
	var let string
	for i:=0; i<digits; i++{
		let += letters[wordnum[i]]
	}
	return let
}

func count(dig int, cable chan []int){
	if (dig>=0 && br!=true){
		for i:=0; i < len(letters); i++ {
			word[dig]=i
			check(word)
//			fmt.Println(word)
		}
//		fmt.Println(word)
		increase(word, dig-1)
		count(dig, cable)
	}
}

func increase(word []int, cur int){
	if (word[0] < len(letters) && br!=true){
		if word[cur] < len(letters)-1{
			word[cur] += 1
//			fmt.Println("test")
		}else {
//			fmt.Println("hello")
			word[cur]=0
			if (cur-1>=0){
			increase(word, cur-1)
			} else{
				br = true
				fmt.Println("Done!")
			}
			
		}
	}
}

func getmd5(str string) string{
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
