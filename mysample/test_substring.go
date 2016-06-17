package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("start!")

	//explaination := "渋谷の宇田川町のど真ん中にある“ダイニングバー　SHIFUKU”！　ビルの4階にあるこちらのお店は合コンや女子会、貸切パーティにおすすめの空間です！　楽しい時簡に一花添える店内の大型のプロジェクターの使用も可能で、いらした皆さんが盛り上がること間違いなし！　もちろんデートなどの普段使いもでき、豊富にそろえる本格焼酎や梅酒などの果実酒、さらに美味しい多国籍料理でおもてなしいたします。また、誕生日会にはメニューには無いスペシャルデザートのご用意もいたしますので、お気軽にお電話下さい!!"
	explaination := ""
	text2 := "渋谷の宇田川町のど真ん中にある“ダイニングバー　SHIFUKU”！　ビルの4"

	fmt.Println(utf8.RuneCountInString(explaination))

	explains := []rune(explaination)

	output := string(explains)
	if len(explains) > 25 {
		output = string(explains[0:25])
	}

	fmt.Printf("%s\n", output)

	fmt.Println(SubstringByUTF8(5, 10, explaination))
	fmt.Println(SubstringByUTF8(0, 2, text2))

}

func SubstringByUTF8(start int, end int, text string) string {

	unicodes := []rune(text)

	output := string(unicodes)
	if start >len(unicodes) {
		output = ""
	} else if len(unicodes) > end {
		output = string(unicodes[start:end])
	} else {
		output = string(unicodes[start:len(unicodes)])
	}
	return output
}
