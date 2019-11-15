package main

import (
	"fmt"
	"github.com/urfave/cli"
	"math"
	"os"
	"strconv"
)

func main(){

	app:=cli.NewApp()
	app.Name="lc"
	app.Version="1.0.0"
	app.Commands=[]cli.Command{
		{
			Name:"等额本息",
			Aliases: []string{"a"},
			Usage:"依次输入本金、年利率、分期月数。如 a 70000 0.05184 24",
			Action: func(c *cli.Context)  {
				baseMoney ,_:=strconv.ParseFloat(c.Args().Get(0),64)
				rate ,_:=strconv.ParseFloat(c.Args().Get(1),64)
				month ,_:=strconv.ParseFloat(c.Args().Get(2),64)
				fmt.Printf("当前计算的条件：等额本息 本金%.2f 年利率%.5f 分期月数%.2f\n ",baseMoney,rate,month)
				averageCapitalPlusInterest(month, baseMoney, rate)
			},
		},
		{
			Name:"等额本金",
			Aliases: []string{"e"},
			Usage:"依次输入本金、年利率、分期月数。如 e 70000 0.05184 24",
			Action: func(c *cli.Context)  {
				baseMoney ,_:=strconv.ParseFloat(c.Args().Get(0),64)
				rate ,_:=strconv.ParseFloat(c.Args().Get(1),64)
				month ,_:=strconv.ParseFloat(c.Args().Get(2),64)
				fmt.Printf("当前计算的条件：等额本金 本金%.2f 年利率%.5f 分期月数%.2f\n ",baseMoney,rate,month)
				equaPrincipal(month, baseMoney, rate)
			},
		},
	}

	for{
		var cmd,baseMoney,rate,month string
		fmt.Printf("\n请输入命令或者回车Enter提示命令，或输入exit退出！\n")
		fmt.Scanln(&cmd,&baseMoney, &rate,&month)
		args:=[]string{"",cmd,baseMoney,rate,month}

		 if cmd=="exit"{
			os.Exit(0)
		}

		if cmd!="a" && cmd!="b"  {
			args=[]string{""}
		}
		app.Run(args)
	}

}
//等额本息
func averageCapitalPlusInterest(month float64, baseMoney float64, rate float64)  {
	balance:=baseMoney
	sumMoney:=0.0
	for i := 0.0; i < month; i++ {
		//月利率
		monthRate:=rate/12
		//本月本息
		monthMoney:=baseMoney*monthRate*(math.Pow(1+monthRate,month))/(math.Pow(1+monthRate,month)-1)
		//本月利息
		rateMoney:=monthRate*balance
		//本月本金
		monthBaseMoney:=monthMoney-rateMoney
		//剩余本金
		balance-=monthBaseMoney
		//合计已还金额
		sumMoney+=monthMoney
		fmt.Printf("期数%.f 本月应还%.2f 本金%.2f 利息%.2f\n", i+1, monthMoney,monthMoney-rateMoney,rateMoney)
	}
	fmt.Printf("总金额%.2f 本金%.2f 总利息%.2f",sumMoney, baseMoney, sumMoney-baseMoney)


}
//等额本金
func equaPrincipal(month float64, baseMoney float64, rate float64) {
	sumMoney := 0.0
	for i := 0.0; i < month; i++ {
		monthBaseMoney := baseMoney / month
		monthMoney := monthBaseMoney + (baseMoney-(monthBaseMoney)*i)*(rate/12)
		sumMoney += monthMoney
		fmt.Printf("期数%.f 本月应还%.2f 本金%.2f 利息%.2f\n", i+1, monthMoney, monthBaseMoney, monthMoney-monthBaseMoney)
	}
	fmt.Printf("总金额%.2f 本金%.2f 总利息%.2f", sumMoney, baseMoney, sumMoney-baseMoney)
}