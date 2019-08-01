package _00_899

import (
    "fmt"
    "strconv"
    "strings"
)

/*
811. 子域名访问计数

一个网站域名，如"discuss.leetcode.com"，包含了多个子域名。作为顶级域名，常用的有"com"，下一级则有"leetcode.com"，最低的一级为"discuss.leetcode.com"。当我们访问域名"discuss.leetcode.com"时，也同时访问了其父域名"leetcode.com"以及顶级域名 "com"。

给定一个带访问次数和域名的组合，要求分别计算每个域名被访问的次数。其格式为访问次数+空格+地址，例如："9001 discuss.leetcode.com"。

接下来会给出一组访问次数和域名组合的列表cpdomains 。要求解析出所有域名的访问次数，输出格式和输入格式相同，不限定先后顺序。

示例 1:
输入:
["9001 discuss.leetcode.com"]
输出:
["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]
说明:
例子中仅包含一个网站域名："discuss.leetcode.com"。按照前文假设，子域名"leetcode.com"和"com"都会被访问，所以它们都被访问了9001次。
示例 2
输入:
["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
输出:
["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
说明:
按照假设，会访问"google.mail.com" 900次，"yahoo.com" 50次，"intel.mail.com" 1次，"wiki.org" 5次。
而对于父域名，会访问"mail.com" 900+1 = 901次，"com" 900 + 50 + 1 = 951次，和 "org" 5 次。
*/

func subdomainVisits(cpdomains []string) []string {
    var resultMap = make(map[string]int)

    for _, v := range cpdomains {
        f := strings.Split(v, ` `)
        num, _ := strconv.Atoi(f[0])
        com := strings.Split(f[1], `.`)
        for len(com) >= 1 {
            resultMap[strings.Join(com, `.`)] += num
            com = com[1:]
        }
    }
    var result []string
    for k, v := range resultMap {
        result = append(result, fmt.Sprintf("%d %s", v, k))
    }

    return result
}


func subdomainVisits(cpdomains []string) []string {
    domainMap := make(map[string]int)
    for _, value := range cpdomains {
        temparr := strings.Split(value, " ")
        count, _ := strconv.Atoi(temparr[0])
        domain := temparr[1]
        for i := len(domain) - 1; i >=0 ; i-- {
            if domain[i] == '.' || i == 0 {
                var subdomain string
                if domain[i] == '.' {
                    subdomain = domain[i + 1:]
                } else {
                    subdomain = domain[i:]
                }

                if _, ok := domainMap[subdomain]; ok {
                    domainMap[subdomain] += count
                } else {
                    domainMap[subdomain] = count
                }
            }
        }
    }
    res := []string{}
    for domainName, cnt := range domainMap {
        res = append(res, strconv.Itoa(cnt) + " " + domainName)
    }
    return res
}