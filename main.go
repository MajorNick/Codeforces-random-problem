package main

import (
	
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"math/rand"
	"time"
)

func main(){

	var first,second,page int  = 900,1000,1
	problems := []string{}
	var url string
	var LastsFirst string
	for {
		url = fmt.Sprintf("https://codeforces.com/problemset/page/%d?tags=%d-%d",page,first,second)
		resp,err := http.Get(url)
		if err != nil{
			fmt.Println(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)

		s := string(body)

		p := "/problemset/problem/"
		indexes := []int{}
		startIndex := 0
		for {
			index := strings.Index(s[startIndex:], p)
			if index == -1 {
				break
			}
			indexes = append(indexes, startIndex+index)
			startIndex += index + len(p)
		
		}
		
		curProblems := []string{}
		for i,v := range indexes{
			if i%2 == 0{
				index := strings.IndexRune(s[v:],'>')
				index--
				curProblems = append(curProblems, s[v:v+index])
			}
		}
		if LastsFirst == curProblems[0]{
			break
		}
		page++
		
		LastsFirst = curProblems[0]
		problems = append(problems,curProblems...)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	i := r1.Intn(len(problems))
	res := "https://codeforces.com"+problems[i]
	fmt.Println(res)
}


