package utils

import (
	"math/rand"
	"time"
	"golang.org/x/crypto/sha3"
	"fmt"
)

var a = []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","a","b","c",
"d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","1","2","3","4","5","6","7","8","9","0"}



func CreateSalt()(s string){
	rand.Seed(time.Now().Unix())
	for i:=0;i<10;i++{
	ran:=rand.Intn(10000)
	s += a[ran%len(a)-1]
	}
	return s
}


func SaltToPassWord(PassWord,Salt string)string{
	PassWord+=Salt

	return fmt.Sprintf("%x",sha3.Sum256([]byte(PassWord)))
}