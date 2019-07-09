package main

import (
    "fmt"
    "math"
)

/*
dp 背包九讲
*/

/* 1
有N 件物品和一个容量为V的背包。放入第i 件物品耗费的费用是C[i]，得到的价值是W[i]。求解将哪些物品装入背包可使价值总和最大。

这是最基础的背包问题，特点是：每种物品仅有一件，可以选择放或不放。用子问题定义状态：即F[i; v] 表示前i 件物品恰放入一个容量为v 的背包可以获得的最大价值。则其状态转移方程便是：
public class test1930 {
    public static void main(String[] args){
        int[] cost={0,2,2,6,5,4};//费用
        int[] value={0,6,3,5,4,6};//价值
        int N=5;//物品个数
        int V=10;//容量
        solvePackage(cost,value,N,V);
    }

    public static void solvePackage(int[] cost,int[] value,int N,int V){
        int[][] dp=new int[N+1][V+1];
        for(int i=1;i<=N;i++){
            for(int j=1;j<=V;j++){
                if(cost[i]<=j){
                    dp[i][j]=Math.max(dp[i-1][j], dp[i-1][j-cost[i]]+value[i]);
                }
                else{
                    dp[i][j]=dp[i-1][j];
                }
            }
        }
        System.out.println(dp[N][V]);
    }
}
*/

func solvePackage(cost []int,value []int,n int,v int){
    var dp [][]int
    for i:=0;i<n+1;i++{
        dp = append(dp,make([]int,v+1))
    }



    for i:=1;i<n;i++{
        for j:=1;j<=v;j++{
            if cost[i] <=j{
                dp[i][j] = int(math.Max(float64(dp[i-1][j]),float64(dp[i-1][j-cost[i]]+value[i])))
            }else{
                dp [i][j] = dp[i-1][j]
            }

        }
    }
    fmt.Println(dp)
}

func solvePackage1(cost []int,value []int,n int,v int){

    dp:=make([]int,v)



    for i:=1;i<=n;i++{
        for j:=v-1;j>=1;j--{
            if cost[i] <=j{
                dp[j] = int(math.Max(float64(dp[j]),float64(dp[j-cost[i]]+value[i])))
            }

        }
    }
    fmt.Println(dp)
}

// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9


func main(){
    cost :=[]int{0,2,2,6,5,4}//费用
    value :=[]int{0,6,3,5,4,6} //价值
    n:=5
    v:=10
    solvePackage(cost,value,n,v)
    solvePackage1(cost,value,n,v)
}