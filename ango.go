package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	KEY        []byte
	word       = flag.String("w", "none", "暗号化もしくは復号化した文字列")
	decryption = flag.Bool("d", false, "復号化")
	cryption   = flag.Bool("c", false, "暗号化")
)

func init() {
	KEY = []byte("doshisha-encourage-password-gene")
}

func main() {
	flag.Parse()

	//fmt.Println(*word, *decryption, *cryption)

	if *decryption {
		replacer := strings.NewReplacer("!", "", "?", "", "4", "", "8", "", "1", "", "3", "", "2", "", "5", "", "6", "", "7", "", "9", "")
		text := replacer.Replace(*word)
		s := strings.NewReader(text)
		r := rot13Reader{s}
		fmt.Println("暗号化word:" + *word)
		fmt.Printf("元word:")
		io.Copy(os.Stdout, &r)
		return
	}

	if *cryption {
		b := []byte(*word)
		text := make([]byte, 2*len(b))
		rand.Seed(time.Now().UnixNano())
		for i, _ := range text {
			if i%2 == 0 {
				text[i] = b[i/2]
			} else {
				switch rand.Int() % 4 {
				case 0:
					text[i] = '!'
				case 1:
					text[i] = '1'
				case 2:
					text[i] = '2'
				case 3:
					text[i] = '3'
				case 4:
					text[i] = '4'
				case 5:
					text[i] = '5'
				case 6:
					text[i] = '6'
				case 7:
					text[i] = '7'
				case 8:
					text[i] = '8'
				case 9:
					text[i] = '9'
				case 10:
					text[i] = '?'
				}
			}
		}
		s := strings.NewReader(string(text))
		r := rot13Reader{s}
		fmt.Println("元word:" + *word)
		fmt.Printf("暗号化word:")
		io.Copy(os.Stdout, &r)
		return
	}
	fmt.Println("Add -d or -c ,so cannot error")
}

func rot13(c byte) byte {
	// 文字を ROT13 変換する関数
	switch {
	case ('A' <= c && c <= 'Z'):
		// 13 文字分ずらす
		return (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		// 13 文字分ずらす
		return (c-'a'+13)%26 + 'a'
	default:
		// 何もしない
		return c
	}
}

type rot13Reader struct {
	// io.Reader をラップする構造体
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	// バイト列を読み込む
	n, err = r.r.Read(p)
	if err != nil {
		// 読み込みに失敗した
		return 0, err
	}
	for i := range p {
		// 各文字に ROT13 を適用する
		p[i] = rot13(p[i])
	}
	return
}
