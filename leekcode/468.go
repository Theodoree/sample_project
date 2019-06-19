package main

import (
    "fmt"
    "strconv"
    "strings"
)

var IPv6Map =map[string]struct{}{
    `0`: struct{}{},
    `1`:struct{}{},
    `2`:struct{}{},
    `3`:struct{}{},
    `4`:struct{}{},
    `5`:struct{}{},
    `6`:struct{}{},
    `7`:struct{}{},
    `8`:struct{}{},
    `9`:struct{}{},
    `a`:struct{}{},
    `b`:struct{}{},
    `c`:struct{}{},
    `d`:struct{}{},
    `e`:struct{}{},
    `f`:struct{}{},
}

func validIPAddress(IP string) string {

    ipv4 := strings.Split(IP, `.`)

    if len(ipv4) == 4 {

        for _, v := range ipv4 {

            if len(v) != 0 && v[0] == '0'{
                return "Neither"
            }

            num, err := strconv.Atoi(v)
            if err !=nil{
                return "Neither"
            }
            if num > 255 {
                return "Neither"
            }
        }
        return "IPv4"
    }

    IPv6 := strings.Split(IP, `:`)

    if len(IPv6) == 8 {
        for _, v := range IPv6 {

            if len(v) > 4 {
                return "Neither"
            }

            for _,vl:=range v{
                if _,ok:=IPv6Map[strings.ToLower(string(vl))];!ok{
                    return "Neither"
                }
            }
        }

        return "IPv6"
    }

    return "Neither"
}

func main() {
    fmt.Println(validIPAddress(`"20EE:FGb8:85a3:0:0:8A2E:0370:7334"`))

}
