# LoanCalculator

## 银行贷款计算器

win平台双击打开`LoanCalculator.exe`

在控制台输入下面命令。


### 等额本息还款方式
命令例子：
```
a 100000 0.05 6
```
输出结果：
```
当前计算的条件：等额本息 本金100000.00 年利率0.05000 分期月数6.00
 期数1 本月应还16910.56 本金16493.90 利息416.67
期数2 本月应还16910.56 本金16562.62 利息347.94
期数3 本月应还16910.56 本金16631.63 利息278.93
期数4 本月应还16910.56 本金16700.93 利息209.63
期数5 本月应还16910.56 本金16770.52 利息140.05
期数6 本月应还16910.56 本金16840.40 利息70.17
总金额101463.39 本金100000.00 总利息1463.39

```

### 等额本金还款方式

命令例子：
```
e 100000 0.05 6
```
输出结果：
```
当前计算的条件：等额本金 本金10000.00 年利率0.05000 分期月数6.00
 期数1 本月应还1708.33 本金1666.67 利息41.67
期数2 本月应还1701.39 本金1666.67 利息34.72
期数3 本月应还1694.44 本金1666.67 利息27.78
期数4 本月应还1687.50 本金1666.67 利息20.83
期数5 本月应还1680.56 本金1666.67 利息13.89
期数6 本月应还1673.61 本金1666.67 利息6.94
总金额10145.83 本金10000.00 总利息145.83

```

