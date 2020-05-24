package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigdigit=[][]string{
{"  000  "," 0   0 ","0     0","0     0","0     0"," 0   0 ","  000  "},
{" 1 ","11 "," 1 "," 1 "," 1 "," 1 ","111"},
}

func main() {
	if len(os.Args)==1 {
		fmt.Printf("usage: %s [-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
		}
	stringofdigits:=os.Args[1]
	sline:=[]string{}	
	if len(os.Args)==3 && (os.Args[1]=="-b" || os.Args[1]=="--bar") {
		sline=append(sline,"")
		stringofdigits=os.Args[2]
		}
	for row:=range bigdigit[0]{
		line:=""
		for column:=range stringofdigits {
			digit:=stringofdigits[column]-'0'
			if 0<=digit && digit<=9 {
				line+=" "+bigdigit[digit][row]+" "
				} else {
				log.Fatal("invalid number")
				}
			}
		sline=append(sline,line)	
		}
	if os.Args[1]=="-b" || os.Args[1]=="--bar" {
		line:=""
		for range sline[1] {
			line+="*"
			}
		sline[0]=line
		sline=append(sline,line)
		}
	for row:=range sline{
		fmt.Println(sline[row])
		}
}