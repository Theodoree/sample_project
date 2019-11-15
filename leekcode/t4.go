package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "sort"
    "strconv"
    "strings"
    "time"
    
    "github.com/360EntSecGroup-Skylar/excelize"
    
    // . "github.com/Theodoree/sample_project/leekcode/utils"
)

const (
    null = 0x7777777
)

type n struct {
    buf []byte `json:"buf"`
}



func main() {
    
    
    readXlsxBy("/Users/ted/Downloads/尚德消耗成本底稿(1).xlsx","/Users/ted/Downloads/json")
}

var funs = ``
var val = ``

func readXlsx() {
    xlsx, err := excelize.OpenFile("/Users/ted/Downloads/兰芝-微信直播商品-汇总.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    var result []interface{}
    rows := xlsx.GetRows("直播间选品")
    
    fmt.Println(rows[0])
    var unionIndex, titleIndex, priceIndex, imgIndex int
    
    for k, v := range rows[0] {
        switch v {
        // case "sku_id":
        //     skuIndex = k
        case "title":
            titleIndex = k
        case "price":
            priceIndex = k
        case "img":
            imgIndex = k
        case "union_url":
            unionIndex = k
            
        }
    }
    
    for i := 1; i < len(rows); i++ {
        row := rows[i]
        price := strings.Split(row[priceIndex], `￥`)
        if len(price) > 1 {
            row[priceIndex] = price[1]
        } else {
            row[priceIndex] = price[0]
        }
        
        Sku := struct {
            Title    string
            OriPrice string
            Price    string
            UnionUrl string
            Img      string
        }{
            Title:    row[titleIndex],
            OriPrice: row[priceIndex],
            Price:    row[priceIndex],
            UnionUrl: row[unionIndex],
            Img:      row[imgIndex],
        }
        result = append(result, &Sku)
    }
    
    b, _ := json.Marshal(result)
    fmt.Printf("%s", b)
    
}

type Stats struct {
    Date  string `json:"date,omitempty"`
    Click uint64 `json:"click,omitempty"`
    Pv    uint64 `json:"pv,omitempty"`
    
    CTR string  `json:"ctr,omitempty"`
    CPM string  `json:"cpm,omitempty"`
    CPC float64 `json:"cpc,omitempty"`
    
    Cost    float64   `json:"cost,omitempty"`
    Balance float64   `json:"balance,omitempty"`
    Deposit []float64 `json:"deposit,omitempty"`
}

func readXlsxBy(ExcelPath, WritePath string) {
    xlsx, err := excelize.OpenFile(ExcelPath)
    if err != nil {
        fmt.Println(err)
        
    }
    
    // 读取的表名必须为sheet
    rows := xlsx.GetRows(`9.01-9.30`)
    if len(rows) == 0 {
        return
    }
    
    recordMap := make(map[string]*Stats)
    var arr []*Stats
    
    buf, err := ioutil.ReadFile(WritePath)
    if err == nil {
        err = json.Unmarshal(buf, &arr)
        for _, v := range arr {
            
            recordMap[v.Date] = v
        }
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Date < arr[j].Date
    })
    
    for _, v := range arr {
        ctr, _ := strconv.ParseFloat(v.CTR, 10)
        fmt.Printf("date %s click %d pv %d ctr %s cpm %s  cpc %.2f  cost %.2f balance %.2f  deposit %v %v  \n", v.Date, v.Click, v.Pv, v.CTR, v.CPM, v.CPC, v.Cost, v.Balance, v.Deposit, (float64(v.Pv)*(ctr/100) -float64(v.Click)) > -1 && (float64(v.Pv)*(ctr/100) -float64(v.Click)) < 1  )
    }
    
    return
    var filterRows [][]string
    for _, v := range rows {
        // 这里查看有没有历史数据
        if _, ok := recordMap[convertToFormatDay(v[0])]; !ok {
            filterRows = append(filterRows, v)
        }
    }
    rows = filterRows
    var (
        depositIndex = -1
        ClickIndex   = -1
        CPCIndex     = -1
        costIndex    = -1
        balanceIndex = -1
    )
    return
    /*
       [{"date":"2019-08-28","click":115060,"pv":6374515,"ctr":"1.805","cpm":"10.830","cpc":0.6,"cost":69036,"balance":230964,"deposit":[100000,200000]},{"date":"2019-08-29","click":125841,"pv":7612885,"ctr":"1.653","cpm":"9.918","cpc":0.6,"cost":75504.6,"balance":155459.4},{"date":"2019-08-30","click":121281,"pv":5699295,"ctr":"2.128","cpm":"12.768","cpc":0.6,"cost":72768.6,"balance":82690.8},{"date":"2019-08-31","click":137514,"pv":6348753,"ctr":"2.166","cpm":"12.996","cpc":0.6,"cost":82508.4,"balance":182.400000000038}]
       [{"date":"2019-08-31","click":137514,"pv":6348753,"ctr":"2.166","cpm":"12.996","cpc":0.6,"cost":82508.4,"balance":182.400000000038},{"date":"2019-08-28","click":115060,"pv":6374515,"ctr":"1.805","cpm":"10.830","cpc":0.6,"cost":69036,"balance":230964,"deposit":[100000,200000]},{"date":"2019-08-29","click":125841,"pv":7612885,"ctr":"1.653","cpm":"9.918","cpc":0.6,"cost":75504.6,"balance":155459.4},{"date":"2019-08-30","click":121281,"pv":5699295,"ctr":"2.128","cpm":"12.768","cpc":0.6,"cost":72768.6,"balance":82690.8}]
    */
    rand.Seed(time.Now().UnixNano())
    for _, v := range rows {
        if v[0] == "" {
            continue
        }
        
        if v[0] == "日期" {
            for i, val := range v {
                switch val {
                case "充值":
                    depositIndex = i
                case "点击量":
                    ClickIndex = i
                    depositIndex = -1
                case "CPC":
                    CPCIndex = i
                case "消耗":
                    costIndex = i
                case "余额":
                    balanceIndex = i
                }
            }
            
        } else {
            record := &Stats{}
            
            v[0] = convertToFormatDay(v[0])
            
            if val, ok := recordMap[v[0]]; ok {
                record = val
            }
            
            record.Date = v[0]
            if depositIndex >= 0 {
                deposit, err := strconv.ParseFloat(v[depositIndex], 64)
                if err != nil {
                    fmt.Println("depositIndex", err)
                }
                record.Deposit = append(record.Deposit, deposit)
            }
            
            if ClickIndex >= 0 {
                click, err := strconv.ParseUint(v[ClickIndex], 10, 64)
                if err != nil {
                    fmt.Println(err)
                }
                record.Click = click
            }
            
            if CPCIndex >= 0 {
                CPC, err := strconv.ParseFloat(v[CPCIndex], 64)
                if err != nil {
                    fmt.Println("CPCIndex", err)
                }
                record.CPC = CPC
            }
            
            if costIndex >= 0 {
                cost, err := strconv.ParseFloat(v[costIndex], 64)
                if err != nil {
                    fmt.Println("costIndex", err)
                }
                record.Cost = cost
            }
            
            if balanceIndex >= 0 {
                balance, err := strconv.ParseFloat(v[balanceIndex], 64)
                if err != nil {
                    fmt.Println("balanceIndex", err)
                }
                record.Balance = balance
            }
            if record.CPC > 0 && record.Click > 0 && record.Cost > 0 {
                var rate = 1.9
                f := rand.Intn(30)
                rate *= (float64(f-15) / 100) + 1
                record.CTR = fmt.Sprintf("%.3f", rate)
                record.Pv = uint64(float64(record.Click) / (rate / 100))
                record.CPM = fmt.Sprintf("%.3f", record.Cost/float64(record.Pv)*1000)
            }
            
            recordMap[v[0]] = record
        }
        
    }
    
    var recordArr []*Stats
    
    for _, v := range recordMap {
        recordArr = append(recordArr, v)
    }
    
    buf, err = json.Marshal(&recordArr)
    if err != nil {
        fmt.Println(err)
    }
    
    err = ioutil.WriteFile(WritePath, buf, 0644)
    fmt.Println(err)
    
}

func convertToFormatDay(excelDaysString string) string {
    baseDiffDay := 38719
    curDiffDay := excelDaysString
    b, _ := strconv.Atoi(curDiffDay)
    realDiffDay := b - baseDiffDay
    realDiffSecond := realDiffDay * 24 * 3600
    baseOriginSecond := 1136185445
    resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
    return resultTime
}
