package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
题目描述
1 总体说明
考生需要模拟实现一个简单的自动售货系统，实现投币、购买商品、退币、查询库存商品及存钱盒信息的功能。
系统初始化时自动售货机中商品为6种商品,商品的单价参见1.1规格说明，存钱盒内放置1元、2元、5元、10元钱币，商品数量和钱币张数通过初始化命令设置，参见2.1 系统初始化。
1.1规格说明
1. 商品:每种商品包含商品名称、单价、数量三种属性，其中商品名不重复。考生不能修改商品名称和单价，初始化命令设置商品数量。这些信息在考试框架中进行定义，考生在实现功能代码时可直接使用。
商品 名称
单价
数量
A1	2	X
A2	3	X
A3	4	X
A4	5	X
A5	8	X
A6	6	X
2. 存钱盒信息：钱币面额、张数两种属性。初始化命令设置各种面额钱币张数。这些信息在考试框架中进行定义，考生在实现功能代码时可直接使用。
钱币面额
张数
10元
X
5元
X
2元	X
1元	X
3. 退币原则 ：
1) 根据系统存钱盒内钱币的 信息 ，按钱币总张数最少的原则进行退币。
2) 如果因零钱不足导致不能退币，则尽最大可能退币，以减少用户损失。
例如：假设存钱盒内只有4张2元，无其它面额钱币。如果需要退币7元，系统因零钱不足无法退币，则继续尝试退币6元，最终系统成功退币3张2元,用户损失1元钱币。
4. 投币操作说明：每次投币成功，投入的钱币面额累加到投币余额；同时，本次投入的钱币放入存钱盒中，存钱盒相应面额钱币增加。
5. 投币余额：指当前自动售货机中用户剩余的可购买商品的钱币总额；例如：投入2元面额的钱币，投币余额增加2元；购买一件价格2元的商品，投币余额减少2元；
6. 投币余额约束：投币余额不能超过10元。
7. 退币操作说明：退币操作需要遵守 退币原则 ；退币成功后，投币余额清零，同时扣除存钱盒相应的金额。
8. 购买商品操作说明：一次仅允许购买一件商品；购买商品成功后，自动售货机中对应商品数量减1，投币余额扣除本次购买商品的价格。
2 操作说明
命令字与第一个参数间使用一个空格分隔，多条命令采用分号隔开。考试系统会对输入命令格式进行处理，考生不需要关注输入命令格式的合法性，只需要实现命令处理函数。
2.1 系统初始化
命令格式：
r A1 数量 -A2 数量 -A3 数量 -A4 数量 -A5 数量 -A6 数量 1 元张数 -2 元张数 -5 元张数 -10 元张数
参数名称
参数说明
类型
取值范围
A1数量
商品A1数量
整数
[0,10]
A2数量
商品A2数量
整数
[0,10]
A3数量
商品A3数量
整数
[0,10]
A4数量
商品A4数量
整数
[0,10]
A5数量
商品A5数量
整数
[0,10]
A6数量
商品A6数量
整数
[0,10]
1元张数
面额1元钱币张数
整数
[0,10]
2元张数
面额2元钱币张数
整数
[0,10]
5元张数
面额5元钱币张数
整数
[0,10]
10元张数
面额10元钱币张数
整数
[0,10]
商品和各种面额钱币取值范围只是作为初始化命令的限制，其它场景下不限制取值范围；考试框架已经实现取值范围的检查，考生不需要关注。
功能说明：设置自动售货机中商品数量和存钱盒各种面额的钱币张数；
约束说明：系统在任意阶段均可执行r初始化系统；考生不需要关注参数的合法性，不需要关注增加或缺少参数的场景；
输出说明：输出操作成功提示（执行完r命令后系统会自动输出操作结果，考生不需要再次调用输出函数），例：
命令	输出	含义
r 6-5-4-3-2-1 4-3-2-1;	S001:Initialization is successful	初始化成功
2.2 投币
命令格式：p 钱币面额
功能说明：
（1） 如果投入非1元、2元、5元、10元的钱币面额（钱币面额不考虑负数、字符等非正整数的情况），输出“E002:Denomination error”；
（2） 如果存钱盒中1元和2元面额钱币总额小于本次投入的钱币面额，输出“E003:Change is not enough, pay fail”，但投入1元和2元面额钱币不受此限制。
（3） 如果投币余额大于10元，输出“E004:Pay the balance is beyond the scope biggest”；
（4） 如果自动售货机中商品全部销售完毕，投币失败。输出“E005:All the goods sold out”；
（5） 如果投币成功，输出“S002:Pay success,balance=X”；
约束说明：
（1） 系统在任意阶段都可以投币；
（2） 一次投币只能投一张钱币；
（3） 同等条件下，错误码的优先级：E002 > E003 > E004 > E005；
输出说明：如果投币成功，输出“S002:Pay success,balance=X”。
例：
命令
输出
p 10;
S002:Pay success,balance=10
2.3 购买商品
命令格式：b 商品名称
功能说明：
（1） 如果购买的商品不在商品列表中，输出“E006:Goods does not exist”；
（2） 如果所购买的商品的数量为0，输出“E007:The goods sold out”；
（3） 如果投币余额小于待购买商品价格，输出“E008:Lack of balance”；
（4） 如果购买成功，输出“S003:Buy success,balance=X”；
约束说明：
（1） 一次购买操作仅能购买一件商品，可以多次购买；
（2） 同等条件下，错误码的优先级：E006 > E007 > E008；
输出说明：
如果购买成功，输出“S003:Buy success,balance=X”。
例:
命令
输出
b A1;
S003:Buy success,balance=8
2.4 退币
命令格式：c
功能说明：
（1） 如果投币余额等于0的情况下，输出“E009:Work failure”；
（2） 如果投币余额大于0的情况下，按照 退币原则 进行“找零”，输出退币信息；
约束说明：
（1） 系统在任意阶段都可以退币；
（2） 退币方式必须按照 退币原则 进行退币；
输出说明：如果退币成功，按照 退币原则 输出退币信息。
例，退5元钱币:
命令
输出
c;
1 yuan coin number=0
2 yuan coin number=0
5 yuan coin number=1
10 yuan coin number=0
2.5 查询
命令格式：q 查询类别
功能说明：
（1） 查询自动售货机中商品信息，包含商品名称、单价、数量。 根据商品数量从大到小进行排序；商品数量相同时，按照商品名称的先后顺序进行排序 。
例如：A1的商品名称先于A2的商品名称，A2的商品名称先于A3的商品名称。
（2） 查询存钱盒信息，包含各种面额钱币的张数；
（3） 查询类别如下表所示:
查询类别
查询内容
0
查询商品信息
1	查询存钱盒信息
如果“查询类别”参数错误，输出“E010:Parameter error”。“查询类别”参数错误时，不进行下面的处理；
输出说明：
“查询类别”为0时，输出自动售货机中所有商品信息（商品名称单价数量）例：
命令
输出
q 0;
A1 2 6
A2 3 5
A3 4 4
A4 5 3
A5 8 2
A6 6 0
“查询类别”为1时，输出存钱盒信息（各种面额钱币的张数），格式固定。例：
命令
输出
q 1;
1 yuan coin number=4
2 yuan coin number=3
5 yuan coin number=2
10 yuan coin number=1

输入描述:
依照说明中的命令码格式输入命令。

输出描述:
输出执行结果

示例1
输入
复制
r 1-1-1-1-1-1 10-5-2-1;p 1;q 1;
输出
复制
S001:Initialization is successful
S002:Pay success,balance=1
1 yuan coin number=11
2 yuan coin number=5
5 yuan coin number=2
10 yuan coin number=1
*/

var (
	prices   = map[string]int{"A1": 2, "A2": 3, "A3": 4, "A4": 5, "A5": 8, "A6": 6}
	moneySeq = []int{1, 2, 5, 10}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		cmds := strings.Split(string(line), ";")
		var machine *SellMachine

		for _, cmd := range cmds {
			if len(cmd) == 0 || cmd == "\n" {
				continue
			}
			info := strings.Split(cmd, " ")

			switch info[0] {
			case "r":
				if len(info) != 3 {
					return
				}
				machine, err = InitSellMachine(info[1], info[2])
				if err != nil {
					return
				}

			case "p":
				b, _ := strconv.Atoi(info[1])
				machine.UpdateBalance(b)
			case "b":
				machine.Sell(info[1])
			case "c":
				machine.Cancel()
			case "q":
				t, _ := strconv.Atoi(info[1])
				machine.Search(t)
			default:
				fmt.Printf("E010:Parameter error")
			}
		}
	}
}

type goods struct {
	Name  string
	Price int
	Num   int
}

type money struct {
	Currency int
	Num      int
}

type GoodsSlince []*goods
type MoneyScline []*money
type SellMachine struct {
	GoodsInfo GoodsSlince
	MoneyInfo MoneyScline
	Balance   int
}

func InitSellMachine(initGoods, initMoney string) (*SellMachine, error) {
	goodsInfo, err := InitGoods(initGoods)
	if err != nil {
		return nil, err
	}
	moneyInfo, err := InitMoney(initMoney)
	if err != nil {
		return nil, err
	}

	fmt.Printf("S001:Initialization is successful\n")
	return &SellMachine{
		GoodsInfo: goodsInfo,
		MoneyInfo: moneyInfo,
	}, nil
}

func (m *SellMachine) UpdateBalance(b int) {
	if isCurrencyValueErr(b) {
		fmt.Printf("%s\n", "E002:Denomination error")
	} else if m.isTooBigger(b) {
		fmt.Printf("%s\n", "E003:Change is not enough, pay fail")
	} else if b > 10 {
		fmt.Printf("%s\n", "E004:Pay the balance is beyond the scope biggest")
	} else if m.outOfSell() {
		fmt.Printf("%s\n", "E005:All the goods sold out")
	} else {
		m.Balance += b
		m.MoneyInfo.Add(b, 1)
		fmt.Printf("S002:Pay success,balance=%d\n", m.Balance)
	}
}

func isCurrencyValueErr(b int) bool {
	err := true
	for _, m := range moneySeq {
		if b == m {
			err = false
		}
	}
	return err
}

func (m *SellMachine) isTooBigger(b int) bool {
	value := 0
	for _, val := range m.MoneyInfo {
		if val.Currency == 1 || val.Currency == 2 {
			value += val.Currency * val.Num
		}
	}
	return value < b
}

func (m *SellMachine) outOfSell() bool {
	for _, g := range m.GoodsInfo {
		if g.Num > 0 {
			return false
		}
	}
	return true
}

func (m *SellMachine) Cancel() {
	if m.Balance <= 0 {
		fmt.Printf("E009:Work failure")
	} else {
		m.doCancel()
	}
}

func (m *SellMachine) doCancel() {
	refund := map[int]int{}
	for i := len(moneySeq) - 1; i >= 0; {
		if m.Balance >= moneySeq[i] && m.MoneyInfo[i].Num > 0 {
			refund[moneySeq[i]]++
			m.MoneyInfo[i].Num--
			m.Balance -= moneySeq[i]
		} else {
			i--
		}
	}
	for _, s := range moneySeq {
		fmt.Printf("%d yuan coin number=%d\n", s, refund[s])
	}
}

func (m *SellMachine) Search(t int) {
	if t == 0 {
		m.searchGoods()
	} else if t == 1 {
		m.searchMoney()
	} else {
		fmt.Printf("E010:Parameter error\n")
	}
}

func (m *SellMachine) searchMoney() {
	for _, s := range m.MoneyInfo {
		fmt.Printf("%d yuan coin number=%d\n", s.Currency, s.Num)
	}
}

func (m *SellMachine) searchGoods() {
	tmpGoodsInfo := m.GoodsInfo.Copy()
	sort.Sort(tmpGoodsInfo)
	for _, g := range tmpGoodsInfo {
		fmt.Printf("%s %d %d\n", g.Name, g.Price, g.Num)
	}
}

func (m *SellMachine) Sell(name string) {
	if m.goodsNotExist(name) {
		fmt.Printf("E006:Goods does not exist\n")
	} else if m.goodsOutOfSell(name) {
		fmt.Printf("E007:The goods sold out\n")
	} else if m.Balance < prices[name] {
		fmt.Printf("E008:Lack of balance\n")
	} else {
		m.doSell(name)
		fmt.Printf("S003:Buy success,balance=%d\n", m.Balance)
	}
}

func (m *SellMachine) goodsNotExist(name string) bool {
	for _, g := range m.GoodsInfo {
		if g.Name == name {
			return false
		}
	}
	return true
}
func (m *SellMachine) goodsOutOfSell(name string) bool {
	for _, g := range m.GoodsInfo {
		if g.Name == name && g.Num > 0 {
			return false
		}
	}
	return true
}
func (m *SellMachine) doSell(name string) {
	for _, g := range m.GoodsInfo {
		if g.Name == name && g.Num > 0 {
			g.Num--
			m.Balance -= g.Price
			break
		}
	}
}

func InitGoods(initCmd string) (GoodsSlince, error) {
	nums := strings.Split(initCmd, "-")
	if len(nums) != len(prices) {
		return nil, fmt.Errorf("Init with error goods cmd: %s", initCmd)
	}

	arrs := make(GoodsSlince, len(prices))
	for i := 0; i < len(prices); i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		arrs[i] = &goods{
			Name:  fmt.Sprintf("A%d", i+1),
			Price: prices[fmt.Sprintf("A%d", i+1)],
			Num:   num,
		}
	}
	return arrs, nil
}

func InitMoney(initCmd string) (MoneyScline, error) {
	nums := strings.Split(initCmd, "-")
	if len(nums) != len(moneySeq) {
		return nil, fmt.Errorf("Init with error money cmd: %s", initCmd)
	}

	arrs := make(MoneyScline, len(moneySeq))
	for i := 0; i < len(moneySeq); i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		arrs[i] = &money{
			Currency: moneySeq[i],
			Num:      num,
		}
	}
	return arrs, nil
}

func (g GoodsSlince) Copy() GoodsSlince {
	tmpGoodsInfo := make(GoodsSlince, len(g))
	for i, g := range g {
		tmpGoodsInfo[i] = &goods{
			Name:  g.Name,
			Price: g.Price,
			Num:   g.Num,
		}
	}
	return tmpGoodsInfo
}

func (g GoodsSlince) Len() int {
	return len(g)
}
func (g GoodsSlince) Less(i, j int) bool {
	if g[i].Num < g[j].Num {
		return true
	} else if g[i].Num == g[j].Num {
		return g[i].Name < g[j].Name
	} else {
		return false
	}
}

func (g GoodsSlince) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g GoodsSlince) Print() {
	for _, t := range g {
		fmt.Printf("%+v", *t)
	}
	fmt.Println()
}

func (g MoneyScline) Print() {
	for _, t := range g {
		fmt.Printf("%+v", *t)
	}
	fmt.Println()
}

func (g MoneyScline) Add(currency, num int) {
	for _, m := range g {
		if m.Currency == currency {
			m.Num += num
			break
		}
	}
}
