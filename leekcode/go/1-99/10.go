package __99

import (
	"fmt"
	"regexp"
)

//func isMatch(s string, p string) bool {
//	if s == "" {
//		return false
//	}
//
//	pIndex := 0
//	var last string
//	var flag bool
//	for i := 0; i < len(s); i++ {
//
//		if flag {
//			if last == string(s[i]) {
//				continue
//			}
//			flag = false
//			pIndex++
//
//		}
//
//		if !flag && pIndex < len(p) {
//			switch string(p[pIndex]) {
//			case `*`:
//				last = string(p[pIndex-1])
//				if last == `.` {
//					return true
//				}
//				flag = true
//			case `.`:
//				pIndex++
//			default:
//				if s[i] != p[pIndex] {
//					if pIndex+1 < len(p) {
//						if p[pIndex+1] == '*'{
//							pIndex++
//							continue
//						}
//					}
//					return false
//
//				}
//				pIndex++
//
//			}
//		}
//	}
//
//	if pIndex == len(p)-1 {
//		return true
//	}
//
//	return false
//
//}



func isMatch(s string, p string) bool {

	if p != `` && s == `` && p != `.*`{
		return false
	}



	Regexp,_:=regexp.Compile(p)

	return  fmt.Sprintf("%s",Regexp.Find([]byte(s))) == s

}


func main() {

	fmt.Println(isMatch(`aab`, `c*a*b`))

}
