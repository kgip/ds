package ds

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func RemoveDuplicatedItem(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var pre, next int
	for next = 1; next < len(arr); next++ {
		if next-pre == 1 && arr[pre] != arr[next] {
			pre++
		} else if next-pre > 1 && arr[pre] != arr[next] {
			pre++
			arr[pre] = arr[next]
		}
	}
	return arr[:pre+1]
}

func TestRemoveDuplicatedItem(t *testing.T) {
	t.Log(RemoveDuplicatedItem([]int{1, 2, 2, 3, 4, 4, 4, 4, 5}))
}

func TransformNum(num string) int {
	numMap := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'a': 10, 'B': 11, 'b': 11, 'C': 12, 'c': 12, 'D': 13, 'd': 13, 'E': 14, 'e': 14, 'F': 15, 'f': 15}
	var result int
	for i := 2; i < len(num); i++ {
		result *= 16
		result += numMap[num[i]]
	}
	return result
}

func TestTransformNum(t *testing.T) {
	t.Log(TransformNum("0x43"))
}

func GetPrimeFactors(num int) []int {
	var results []int
	max := int(math.Sqrt(float64(num)))
	for i := 2; i <= max; i++ {
		if num%i == 0 {
			results = append(results, GetPrimeFactors(i)...)
			results = append(results, GetPrimeFactors(num/i)...)
			return results
		}
	}
	results = append(results, num)
	return results
}

func TestGetPrimeFactors(t *testing.T) {
	t.Log(GetPrimeFactors(424431242))
}

func TestName(t *testing.T) {
	var arr []int
	//arr = append(arr, 1, 1, 1, 1)
	//t.Log(cap(arr))
	for i := 0; i < 20; i++ {
		arr = append(arr, 1)
		t.Log(len(arr), cap(arr))
	}
}

func Calculate(expression string) float64 {
	var stack1 []float64 //存储数字
	var stack2 []byte    //存储运算符
	charNumberMap := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	var curNumber float64
	var isNegative bool
	for i := 0; i < len(expression); i++ {
		if expression[i] == ' ' { //空格跳过
			continue
		} else if expression[i] >= '0' && expression[i] <= '9' { //数字
			curNumber = curNumber*10 + float64(charNumberMap[expression[i]])
			if i >= len(expression)-1 || expression[i+1] < '0' || expression[i+1] > '9' { //下一个字符不是数字
				if isNegative {
					curNumber = -curNumber
				}
				stack1 = append(stack1, curNumber)
				curNumber = 0
				isNegative = false
			}
		} else if expression[i] == '+' || expression[i] == '-' {
			if expression[i] == '-' && (i <= 0 || expression[i-1] == '{' || expression[i-1] == '[' || expression[i-1] == '(') { // '-'是负号
				isNegative = true
				continue
			}
			if len(stack2) > 0 && (stack2[len(stack2)-1] == '+' || stack2[len(stack2)-1] == '-' || stack2[len(stack2)-1] == '*' || stack2[len(stack2)-1] == '/') { //检查符号栈顶元素
				n1 := stack1[len(stack1)-1]
				n2 := stack1[len(stack1)-2]
				if stack2[len(stack2)-1] == '+' {
					stack1[len(stack1)-2] = n1 + n2
				} else if stack2[len(stack2)-1] == '-' {
					stack1[len(stack1)-2] = n2 - n1
				} else if stack2[len(stack2)-1] == '*' {
					stack1[len(stack1)-2] = n1 * n2
				} else {
					stack1[len(stack1)-2] = n2 / n1
				}
				stack1 = stack1[:len(stack1)-1]
				stack2[len(stack2)-1] = expression[i]
			} else {
				stack2 = append(stack2, expression[i])
			}
		} else if expression[i] == '*' || expression[i] == '/' {
			if len(stack2) > 0 && (stack2[len(stack2)-1] == '*' || stack2[len(stack2)-1] == '/') { //检查符号栈顶元素
				n1 := stack1[len(stack1)-1]
				n2 := stack1[len(stack1)-2]
				if stack2[len(stack2)-1] == '*' {
					stack1[len(stack1)-2] = n1 * n2
				} else {
					stack1[len(stack1)-2] = n2 / n1
				}
				stack1 = stack1[:len(stack1)-1]
				stack2[len(stack2)-1] = expression[i]
			} else {
				stack2 = append(stack2, expression[i])
			}
		} else if expression[i] == '{' || expression[i] == '[' || expression[i] == '(' { //左括号
			stack2 = append(stack2, expression[i])
		} else { // 右括号
			for stack2[len(stack2)-1] != '(' && stack2[len(stack2)-1] != '[' && stack2[len(stack2)-1] != '{' {
				char := stack2[len(stack2)-1]
				n1 := stack1[len(stack1)-1]
				n2 := stack1[len(stack1)-2]
				if char == '+' {
					stack1[len(stack1)-2] = n1 + n2
				} else if char == '-' {
					stack1[len(stack1)-2] = n2 - n1
				} else if char == '*' {
					stack1[len(stack1)-2] = n1 * n2
				} else {
					stack1[len(stack1)-2] = n2 / n1
				}
				stack1 = stack1[:len(stack1)-1]
				stack2 = stack2[:len(stack2)-1]
			}
			stack2 = stack2[:len(stack2)-1]
		}
	}
	for len(stack2) > 0 {
		char := stack2[len(stack2)-1]
		n1 := stack1[len(stack1)-1]
		n2 := stack1[len(stack1)-2]
		if char == '+' {
			stack1[len(stack1)-2] = n1 + n2
		} else if char == '-' {
			stack1[len(stack1)-2] = n2 - n1
		} else if char == '*' {
			stack1[len(stack1)-2] = n1 * n2
		} else {
			stack1[len(stack1)-2] = n2 / n1
		}
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
	}
	return stack1[0]
}

func TestCalculate(t *testing.T) {
	t.Log(Calculate("-4 * 3 + 5 * 4/{7-[10/(4-2)]}"))
}

func Sum(n1, n2 string) string {
	charNumberMap := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	var n, s int
	var result string
	for i := 0; len(n1)-1-i >= 0 || len(n2)-1-i >= 0; i++ {
		if len(n1)-1-i >= 0 && len(n2)-1-i >= 0 {
			s = charNumberMap[n1[len(n1)-1-i]] + charNumberMap[n2[len(n2)-1-i]] + n
		} else if len(n1)-1-i >= 0 {
			s = charNumberMap[n2[len(n2)-1-i]] + n
		} else {
			s = charNumberMap[n2[len(n2)-1-i]] + n
		}
		if s >= 10 {
			n = 1
			s -= 10
		} else {
			n = 0
		}
		result = fmt.Sprintf("%d", s) + result
	}
	return result
}

func TestSum(t *testing.T) {
	t.Log(Sum("111", "21239"))
}

func GetFirstChar(str string) string {
	charCount := make([]int, 26)
	for i := 0; i < len(str); i++ {
		charCount[str[i]-'a'] += 1
	}
	for i := 0; i < len(str); i++ {
		if charCount[str[i]-'a'] == 1 {
			return fmt.Sprintf("%c", str[i])
		}
	}
	return ""
}

func TestGetFirstChar(t *testing.T) {
	t.Log(GetFirstChar("abcdadige"))
}

func GetPrimePer(n int) []int {
	var n1, n2 int
	for i := 1; i <= n/2; i++ {
		if IsPrime(i) {
			if n1 == 0 || n2 == 0 {
				n1 = i
				n2 = n - i
			} else if n2-n1 > (n-i)-i {
				n1 = i
				n2 = n - i
			}
		}
	}
	return []int{n1, n2}
}

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestGetPrimePer(t *testing.T) {
	t.Log(GetPrimePer(30))
}

func dp(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	} else {
		if m < n {
			return dp(m, m)
		} else {
			if m == n {
				return dp(m, n-1) + 1
			}
			return dp(m-n, n) + dp(m, n-1)
		}
	}
}

func TestDp(t *testing.T) {
	t.Log(dp(7, 3))
}

func GetBitCount(n int) int {
	var count int
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func TestGetBitCount(t *testing.T) {
	t.Log(GetBitCount(5))
}

func Massage(nums []int) int {
	var preTotal, total int
	for i := 0; i < len(nums); i++ {
		maxTotal := int(math.Max(float64(total), float64(preTotal+nums[i])))
		preTotal = total
		total = maxTotal
	}
	return total
}

func TestMassage(t *testing.T) {
	t.Log(Massage([]int{2, 7, 9, 3, 1}))
}

func WaysToStep(n int) int {
	var n1, n2, n3 = 1, 0, 0
	for i := 1; i <= n; i++ {
		curr := 0
		if i-1 >= 0 {
			curr = n1
		}
		if i-2 >= 0 {
			curr += n2
		}
		if i-3 >= 0 {
			curr += n3
		}
		n3 = n2
		n2 = n1
		n1 = curr % 1000000007
	}
	return n1
}

func WaysToStep2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return (WaysToStep2(n-1) + WaysToStep2(n-2) + WaysToStep2(n-3)) % 1000000007
}

func TestWaysToStep(t *testing.T) {
	t.Log(WaysToStep(5))
}

func ReverseBits(num int) int {
	var pre, next, preZeroIndex, maxLength = 0, 0, -1, 0
	for next = 0; next < 32; next++ {
		if num%2 == 0 { //当前位为0
			if preZeroIndex >= 0 { //已经遇到0
				maxLength = int(math.Max(float64(maxLength), float64(next-pre)))
				pre = preZeroIndex + 1
			}
			preZeroIndex = next
		}
		num >>= 1
	}
	return int(math.Max(float64(maxLength), float64(next-pre)))
}

func DpReverseBits(num int) int {
	var maxLength, length, zeroIndex int
	for i := 1; i <= 32; i++ {
		if num%2 == 0 {
			if zeroIndex > 0 {
				maxLength = int(math.Max(float64(maxLength), float64(length)))
				length = i - zeroIndex
			} else {
				length = i
			}
			zeroIndex = i
		} else {
			length++
		}
		num >>= 1
	}
	return maxLength
}

func TestReverseBits(t *testing.T) {
	t.Log(DpReverseBits(1775))
}

func GetFirstMaxGCRadioRange(dna string, n int) string {
	var left, right, currMaxLeft, currMaxRight int
	var GCCount, currMaxGCCount int
	if len(dna) < n {
		return ""
	} else if len(dna) == n {
		return dna
	}
	for right = 0; right < n; right++ {
		if dna[right] == 'G' || dna[right] == 'C' {
			GCCount++
		}
	}
	currMaxLeft, currMaxRight, currMaxGCCount = left, right, GCCount
	for right = n; right < len(dna); right++ {
		if dna[right] == 'G' || dna[right] == 'C' {
			GCCount++
		}
		if dna[left] == 'G' || dna[left] == 'C' {
			GCCount--
		}
		left++
		if GCCount > currMaxGCCount {
			currMaxLeft, currMaxRight, currMaxGCCount = left, right, GCCount
		}
	}
	return dna[currMaxLeft : currMaxRight+1]
}

func TestGetFirstMaxGCRadioRange(t *testing.T) {
	t.Log(GetFirstMaxGCRadioRange("ACGTCGGAT", 3))
}

func ChangeSong(n int, options string) (int, []int) {
	var index = 1
	var songIndex = make([]int, 4)
	//初始化显示的歌曲
	for i := 1; i <= 4; i++ {
		if i <= n {
			songIndex[i-1] = i
		}
	}
	for i := 0; i < len(options); i++ {
		if n <= 4 {
			if options[i] == 'U' {
				if index == 1 {
					index = n
				} else {
					index -= 1
				}
			} else {
				if index == n {
					index = 1
				} else {
					index += 1
				}
			}
			songIndex = songIndex[:n]
		} else {
			if options[i] == 'U' { //向上
				//检查光标是否在显示的第一首歌
				if index == songIndex[0] {
					if index == 1 {
						index = n //跳到最后一首歌
						for j := 0; j < len(songIndex); j++ {
							songIndex[j] = n + j - 3
						}
					} else {
						index -= 1
						for j := 0; j < len(songIndex); j++ {
							songIndex[j] -= 1
						}
					}
				} else {
					index -= 1
				}
			} else { //向下
				//检查光标是否在显示的最后一首歌
				if index == songIndex[3] {
					if index == n {
						index = 1
						for j := index; j <= len(songIndex); j++ {
							songIndex[j-1] = j
						}
					} else {
						index += 1
						for j := 0; j < len(songIndex); j++ {
							songIndex[j] += 1
						}
					}
				} else {
					index += 1
				}
			}
		}
	}
	return index, songIndex
}

func TestChangeSong(t *testing.T) {
	t.Log(ChangeSong(10, "UUUUUDDDDD"))
}

func GetMaxLengthPublicSubString(str1, str2 string) string {
	var maxPublicSubString string
	if str1 == "" || str2 == "" {
		return maxPublicSubString
	}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			var k = 0
			for ; k <= i && k <= j; k++ {
				if str1[i-k] != str2[j-k] {
					if k > len(maxPublicSubString) {
						maxPublicSubString = str1[i-k+1 : i+1]
					}
					break
				}
			}
			if (k > i || k > j) && k > len(maxPublicSubString) {
				maxPublicSubString = str1[i-k+1 : i+1]
			}
		}
	}
	return maxPublicSubString
}

func TestGetMaxLengthPublicSubString(t *testing.T) {
	t.Log(GetMaxLengthPublicSubString("abcdesdfafe", "hifeaabciea"))
}

//reset -> reset what
//reset board -> board fault
//board add -> where to add
//board delete -> no board at all
//reboot backplane -> impossible
//backplane abort -> install first
func ResolveCommand(input string) string {
	commands := map[string]string{
		"reset":            "reset what",
		"reset board":      "board fault",
		"board add":        "where to add",
		"board delete":     "no board at all",
		"reboot backplane": "impossible",
		"backplane abort":  "install first",
	}
	inputWords := strings.Split(input, " ")
	var result string
	for command, execResult := range commands {
		commandWords := strings.Split(command, " ")
		if len(inputWords) == len(commandWords) {
			for i := 0; i < len(inputWords); i++ {
				if !strings.HasPrefix(commandWords[i], inputWords[i]) {
					break
				}
				if i == len(inputWords)-1 {
					if result != "" {
						return "unknown command"
					}
					result = execResult
				}
			}
		}
	}
	if result == "" {
		result = "unknown command"
	}
	return result
}

func TestResolveCommand(t *testing.T) {
	for _, input := range []string{"r", "r b", "r c", "b ad"} {
		t.Log(ResolveCommand(input))
	}
}

func CanCalculate(numbers []float64, result float64) bool {
	if len(numbers) == 1 {
		return numbers[0] == result
	}
	var canCalculate bool
	for i := 0; i < len(numbers); i++ {
		copiedNumbers := make([]float64, len(numbers))
		copy(copiedNumbers, numbers)
		leftNumbers := append(copiedNumbers[0:i], copiedNumbers[i+1:]...)
		canCalculate = canCalculate || CanCalculate(leftNumbers, result-numbers[i]) ||
			CanCalculate(leftNumbers, numbers[i]-result) ||
			CanCalculate(leftNumbers, numbers[i]+result) ||
			CanCalculate(leftNumbers, result/numbers[i]) ||
			CanCalculate(leftNumbers, numbers[i]*result)
		if result != 0 {
			canCalculate = canCalculate || CanCalculate(leftNumbers, numbers[i]/result)
		}
		if !canCalculate {
			for j := 0; j < len(leftNumbers); j++ {
				copiedLeftNumbers := make([]float64, len(leftNumbers))
				copy(copiedLeftNumbers, leftNumbers)
				subLeftNumbers := append(copiedLeftNumbers[0:j], copiedLeftNumbers[j+1:]...)
				canCalculate = CanCalculate(subLeftNumbers, result/(leftNumbers[j]+numbers[i])) ||
					CanCalculate(subLeftNumbers, (leftNumbers[j]+numbers[i])*result) ||
					CanCalculate(subLeftNumbers, result-numbers[i]*leftNumbers[j]) ||
					CanCalculate(subLeftNumbers, result-numbers[i]/leftNumbers[j]) ||
					CanCalculate(subLeftNumbers, numbers[i]*leftNumbers[j]-result) ||
					CanCalculate(subLeftNumbers, numbers[i]/leftNumbers[j]-result) ||
					CanCalculate(subLeftNumbers, numbers[i]*leftNumbers[j]+result) ||
					CanCalculate(subLeftNumbers, numbers[i]/leftNumbers[j]+result)
				if numbers[i]-leftNumbers[j] != 0 {
					canCalculate = canCalculate || CanCalculate(subLeftNumbers, (numbers[i]-leftNumbers[j])*result) ||
						CanCalculate(subLeftNumbers, result/(numbers[i]-leftNumbers[j]))
				}
				if result != 0 {
					canCalculate = canCalculate || CanCalculate(subLeftNumbers, (leftNumbers[j]+numbers[i])/result)
					if numbers[i] != leftNumbers[j] {
						canCalculate = canCalculate || CanCalculate(subLeftNumbers, (numbers[i]-leftNumbers[j])/result)
					}
				}
			}
		}
		if canCalculate {
			return true
		}
	}
	return canCalculate
}

func TestCanCalculate(t *testing.T) {
	for _, numbers := range [][]float64{{7, 2, 9, 1}} {
		t.Log(CanCalculate(numbers, 24))
	}
}

func ScoreSort(scores [][]interface{}, sortType int) [][]interface{} {
	if len(scores) == 2 && (sortType == 0 && scores[0][1].(int) < scores[1][1].(int) || sortType == 1 && scores[0][1].(int) > scores[1][1].(int)) { //只有两个元素，进行排序
		scores[0], scores[1] = scores[1], scores[0]
	} else if len(scores) > 2 { //有多个元素，对数组进行拆分
		mid := len(scores) / 2
		scores1 := ScoreSort(scores[:mid], sortType)
		scores2 := ScoreSort(scores[mid:], sortType)
		scores = make([][]interface{}, len(scores))
		var index1, index2 int
		for i := 0; i < len(scores); i++ {
			if index1 < len(scores1) && index2 < len(scores2) {
				if sortType == 1 { //降序
					if scores1[index1][1].(int) >= scores2[index2][1].(int) {
						scores[i] = scores1[index1]
						index1++
					} else {
						scores[i] = scores2[index2]
						index2++
					}
				} else {
					if scores1[index1][1].(int) <= scores2[index2][1].(int) {
						scores[i] = scores1[index1]
						index1++
					} else {
						scores[i] = scores2[index2]
						index2++
					}
				}
			} else {
				if index1 >= len(scores1) && index2 < len(scores2) {
					scores[i] = scores2[index2]
					index2++
				} else if index1 < len(scores1) && index2 >= len(scores2) {
					scores[i] = scores1[index1]
					index1++
				}
			}
		}
	}
	return scores
}

func ResolveArgs(command string) ([]string, int) {
	var result []string
	var arg string
	var flag bool
	for i := 0; i < len(command); i++ {
		if command[i] == ' ' && !flag {
			result = append(result, arg)
			arg = ""
		} else if command[i] == '"' && !flag {
			flag = true
		} else if command[i] == '"' && flag {
			flag = false
		} else {
			arg += fmt.Sprintf("%c", command[i])
		}
	}
	result = append(result, arg)
	return result, len(result)
}

func TestResolveArgs(t *testing.T) {
	t.Log(ResolveArgs(`xcopy /s "c:\\11 ddf feef" d:\\f`))
}

func GetOddNumberRange(m int) string {
	var result []string
	m3 := m * m * m
	//找到初始奇数
	for i := 1; i <= m3; i = i + 2 {
		if m3 == m*i+m*(m-1) {
			for j := 0; j < m; j++ {
				result = append(result, fmt.Sprintf("%d", i+2*j))
			}
			break
		}
	}
	return strings.Join(result, " ")
}

func TestGetOddNumberRange(t *testing.T) {
	t.Log(GetOddNumberRange(4))
}

func TrainDepartureList(trainList []int, currTrainList []int) [][]int {
	var results [][]int
	if len(trainList) > 0 {
		for i := 0; i < len(trainList); i++ {
			copiedTrainList := make([]int, len(trainList))
			copy(copiedTrainList, trainList)
			subTrainList := append(copiedTrainList[:i], copiedTrainList[i+1:]...)
			//出站
			results = TrainDepartureList(subTrainList, currTrainList)
			for j := 0; j < len(results); j++ {
				results[j] = append([]int{trainList[i]}, results[j]...)
			}
			//不出站
			currTrainList = append(currTrainList, trainList[i])
			results = append(results, TrainDepartureList(subTrainList, currTrainList)...)
		}
	} else {
		var result []int
		for i := len(currTrainList) - 1; i >= 0; i-- {
			result = append(result, currTrainList[i])
		}
		results = [][]int{result}
	}
	return results
}

func TestTrainDepartureList(t *testing.T) {
	t.Log(TrainDepartureList([]int{1, 2, 3}, nil))
}

func ContainSubStr(str, substr string) bool {
	var bitmap int32
	for i := 0; i < len(str); i++ {
		bitmap = bitmap | 1<<(str[i]-'a')
	}
	for i := 0; i < len(substr); i++ {
		if bitmap&1<<(substr[i]-'a') == 0 {
			return false
		}
	}
	return true
}

func TestContainSubStr(t *testing.T) {
	t.Log(ContainSubStr("abc", "ab"))
}

func Gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return Gcd(b, a%b)
	}
}

func TestGcd(t *testing.T) {
	t.Log(Gcd(22, 5))
}

func GetRepresentation(fraction string) string {
	split := strings.Split(fraction, "/")
	numerator, _ := strconv.Atoi(split[0])
	denominator, _ := strconv.Atoi(split[1])
	gcd := Gcd(denominator, numerator)
	numerator /= gcd
	denominator /= gcd
	//寻找比当前分数小的最大的埃及数
	for i := 2; i <= denominator; i++ {
		if denominator-numerator*i < 0 {
			return fmt.Sprintf("1/%d + %s", i, GetRepresentation(fmt.Sprintf("%d/%d", numerator*i-denominator, denominator*i)))
		} else if numerator == 1 && denominator == i {
			return fmt.Sprintf("1/%d", denominator)
		}
	}
	return ""
}

func TestGetRepresentation(t *testing.T) {
	t.Log(GetRepresentation("7/9"))
}

func MatrixTransform(matrix [][]int, x1, y1, x2, y2 int, x, y int, xn, yn int) {
	var result = make([]int, 5)
	if len(matrix) > 9 || len(matrix[0]) > 9 {
		result[0] = -1
	}
	if x1 > len(matrix)-1 || x1 < 0 || x2 > len(matrix)-1 || x2 < 0 || y1 > len(matrix[0])-1 || y1 < 0 || y2 > len(matrix[0])-1 || y2 < 0 {
		result[1] = -1
	} else {
		matrix[x1][y1], matrix[x2][y2] = matrix[x2][y2], matrix[x1][y1]
	}
	if x < 0 || x > len(matrix)-1 {
		result[2] = -1
	}
	if y < 0 || y > len(matrix[0])-1 {
		result[3] = -1
	}
	if xn < 0 || xn > len(matrix)-1 || yn < 0 || yn > len(matrix[0])-1 {
		result[4] = -1
	}
}

func GetMaxPalindrome(str string) string {
	var maxLeft, maxRight int
	for i := 0; i < len(str); i++ {
		for left, right := i-1, i+1; left >= 0 && right < len(str) && str[left] == str[right]; {
			if right-left > maxRight-maxLeft {
				maxLeft, maxRight = left, right
			}
			left--
			right++
		}
		for left, right := i, i+1; left >= 0 && right < len(str) && str[left] == str[right]; {
			if right-left > maxRight-maxLeft {
				maxLeft, maxRight = left, right
			}
			left--
			right++
		}
	}
	return str[maxLeft : maxRight+1]
}

func TestGetMaxPalindrome(t *testing.T) {
	t.Log(GetMaxPalindrome("cdabbacc"))
}

func GetLongest1Sequence(n int) int {
	var length, maxLength int
	for ; n > 0; n >>= 1 {
		if n%2 > 0 {
			length++
			if length > maxLength {
				maxLength = length
			}
		} else {
			length = 0
		}
	}
	return maxLength
}

func TestGetLongest1Sequence(t *testing.T) {
	t.Log(GetLongest1Sequence(10))
}

//dp[i][j] = max(dp[i-1][j], dp[i][j-i]*i)
func IntegerBreak(n int) int {
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
		if i == 0 {
			for j := 0; j < len(dp[i]); j++ {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if j >= i {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-i]*i)))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][n]
}

func TestIntegerBreak(t *testing.T) {
	t.Log(IntegerBreak(10))
}

//dp[j] = dp[j] || dp[j-a[i]]
func CanGroupSumEqual(arr []int) bool {
	var sum, sum3 int
	var newArr []int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i]%3 == 0 {
			sum3 += arr[i]
		}
		if arr[i]%3 != 0 && arr[i]%5 != 0 {
			newArr = append(newArr, arr[i])
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum/2 - sum3
	var dp = make([][]bool, len(newArr)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, int(math.Abs(float64(target)))+1)
	}
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j]
			if j-newArr[i-1] >= 0 && j-newArr[i-1] < len(dp[0]) {
				dp[i][j] = dp[i][j] || dp[i-1][j-newArr[i-1]]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestCanGroupSumEqual(t *testing.T) {
	t.Log(CanGroupSumEqual([]int{3, 5, 8}))
}

func StatisticsVote(candidates []string, votes []string) map[string]int {
	results := make(map[string]int)
	for i := 0; i < len(candidates); i++ {
		results[candidates[i]] = 0
	}
	for i := 0; i < len(votes); i++ {
		if _, ok := results[votes[i]]; ok {
			results[votes[i]]++
		} else {
			results["invalid"]++
		}
	}
	return results
}

func TestStatisticsVote(t *testing.T) {
	t.Log(StatisticsVote([]string{"A", "B", "C", "D"}, []string{"A", "B", "D", "D", "A", "E"}))
}

//中文读取金额
// xxxx亿/xxxx万/xxxx元.x角x分
//1001000
var NumberToChineseMap = map[int]string{1: "壹", 2: "贰", 3: "叁", 4: "肆", 5: "伍", 6: "陆", 7: "柒", 8: "捌", 9: "玖"}

func ReadChineseYuan(money float64) string {
	result := ReadNumber(int(money))
	result += "元"
	n := int(money*10) % 10
	m := int(money*100) % 10
	if n == 0 && m == 0 {
		result += "整"
	} else {
		if n == 0 {
			result += fmt.Sprintf("零%s分", NumberToChineseMap[m])
		} else {
			result += fmt.Sprintf("%s角", NumberToChineseMap[n])
			if m > 0 {
				result += fmt.Sprintf("%s分", NumberToChineseMap[m])
			}
		}
	}
	return result
}

func ReadNumber(n int) string {
	var result string
	if n >= 100000000 {
		result = ReadNumber(n/100000000) + "亿"
		//读取剩余部分
		subResult := ReadNumber(n % 100000000)
		if n%100000000 < 10000000 {
			result = result + "零" + subResult
		} else {
			result += subResult
		}
	} else if n >= 10000 {
		result = ReadNumber(n/10000) + "万"
		//读取剩余部分
		subResult := ReadNumber(n % 10000)
		if n%10000 < 1000 {
			result = result + "零" + subResult
		} else {
			result += subResult
		}
	} else {
		list := []string{"", "拾", "佰", "仟"}
		for i := 3; i >= 0; i-- {
			m := n / int(math.Pow10(i))
			if m > 0 {
				result += fmt.Sprintf("%s%s", NumberToChineseMap[m], list[i])
			} else if i > 0 && i < 3 && result != "" && !strings.HasSuffix(result, "零") {
				result += "零"
			}
			n %= int(math.Pow10(i))
		}
	}
	return result
}

func TestReadChineseYuan(t *testing.T) {
	t.Log(ReadChineseYuan(102840141010.14))
}

func AddNumberPrefixAndSuffix(str string) string {
	var result string
	var isNumber bool
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' && !isNumber {
			isNumber = true
			result += fmt.Sprintf("*%c", str[i])
		} else if (str[i] < '0' || str[i] > '9') && isNumber {
			isNumber = false
			result += fmt.Sprintf("*%c", str[i])
		} else {
			result += fmt.Sprintf("%c", str[i])
		}
		if isNumber && i == len(str)-1 {
			result += "*"
		}
	}
	return result
}

func TestAddNumberPrefixAndSuffix(t *testing.T) {
	t.Log(AddNumberPrefixAndSuffix("aifeoh23fjei3243jifeaoi0333"))
}

func GetAvg(arr []int) (n int, avg float64) {
	var n1, n2, sum int
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			n1++
		} else if arr[i] > 0 {
			n2++
			sum += arr[i]
		}
	}
	if n2 > 0 {
		avg, _ = strconv.ParseFloat(fmt.Sprintf("%0.1f", float64(sum)/float64(n2)), 64)
	}
	return n1, avg
}

func TestGetAvg(t *testing.T) {
	t.Log(GetAvg([]int{-1, -2, 1, 2, 3, 1}))
}

func GetSelfNumber(n int) int {
	var count int
	for i := 0; i <= n; i++ {
		if strings.HasSuffix(fmt.Sprintf("%d", i*i), fmt.Sprintf("%d", i)) {
			count++
		}
	}
	return count
}

func TestGetSelfNumber(t *testing.T) {
	t.Log(GetSelfNumber(6))
}

func StatisticsLetterCount(str string) string {
	var list [][]int
	//统计字符串个数
	for i := 0; i < len(str); i++ {
		var isExist bool
		for j := 0; j < len(list); j++ {
			if list[j][0] == int(str[i]) {
				list[j][1]++
				isExist = true
			}
		}
		if !isExist {
			list = append(list, []int{int(str[i]), 1})
		}
	}
	//排序
	sort.Slice(list, func(i, j int) bool {
		return list[j][1] < list[i][1] || list[j][1] == list[i][1] && list[j][0] > list[i][0]
	})
	//组合结果
	var result string
	for i := 0; i < len(list); i++ {
		result += fmt.Sprintf("%c", list[i][0])
	}
	return result
}

func TestStatisticLetterCount(t *testing.T) {
	t.Log(StatisticsLetterCount("abcdddace"))
}

func MaxSteps(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var maxSteps int
	var dp = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if maxSteps < dp[i] {
			maxSteps = dp[i]
		}
	}
	return maxSteps
}

func TestMaxSteps(t *testing.T) {
	t.Log(MaxSteps([]int{2, 5, 1, 5, 4, 5}))
}

func GetNumber(n float64) float64 {
	var x1, x2 float64 = 0, n
	var result float64
	var mid, y float64
	for {
		mid = (x1 + x2) / 2
		y = mid*mid*mid - n
		if math.Abs(y) < 0.0000001 {
			result = mid
			break
		}
		if y > 0 {
			x2 = mid
		} else {
			x1 = mid
		}
	}
	result, _ = strconv.ParseFloat(fmt.Sprintf("%0.1f", result), 64)
	return result
}

func TestGetNumber(t *testing.T) {
	t.Log(GetNumber(389272))
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func TreeMinDeep(root *Node) int {
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		return TreeMinDeep(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return TreeMinDeep(root.Right) + 1
	} else {
		return int(math.Min(float64(TreeMinDeep(root.Left)), float64(TreeMinDeep(root.Right)))) + 1
	}
}

func ReversePolishNotation(exp []string) int {
	var stack []int
	for i := 0; i < len(exp); i++ {
		if exp[i] == "+" {
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if exp[i] == "-" {
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if exp[i] == "*" {
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if exp[i] == "/" {
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			n, _ := strconv.ParseInt(exp[i], 10, 64)
			stack = append(stack, int(n))
		}
	}
	return stack[0]
}

func TestNaReversePolishNotation(t *testing.T) {
	t.Log(ReversePolishNotation([]string{"20", "10", "+", "30", "*"}))
}

// (x1, y1) (x2, y2)
// (y2 - y1) / (x2 - x1)
func ThroughMostPoints(points [][]int) int {
	var mostPoints int
	for i := 0; i < len(points); i++ {
		kMap := make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			dy := points[i][1] - points[j][1]
			dx := points[i][0] - points[j][0]
			if dy == 0 {
				kMap["0"] += 1
				if mostPoints < kMap["0"] {
					mostPoints = kMap["0"]
				}
			} else if dx == 0 {
				kMap["00"] += 1
				if mostPoints < kMap["00"] {
					mostPoints = kMap["00"]
				}
			} else {
				var gcd int
				if math.Abs(float64(dx)) >= math.Abs(float64(dy)) {
					gcd = Gcd(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
				} else {
					gcd = Gcd(int(math.Abs(float64(dy))), int(math.Abs(float64(dx))))
				}
				dy /= gcd
				dx /= gcd
				if dy*dx < 0 {
					key := fmt.Sprintf("-%d/%d", int(math.Abs(float64(dy))), int(math.Abs(float64(dx))))
					kMap[key] += 1
					if mostPoints < kMap[key] {
						mostPoints = kMap[key]
					}
				} else {
					key := fmt.Sprintf("%d/%d", int(math.Abs(float64(dy))), int(math.Abs(float64(dx))))
					kMap[key] += 1
					if mostPoints < kMap[key] {
						mostPoints = kMap[key]
					}
				}
			}
		}
	}
	return mostPoints + 1
}

func TestThroughMostPoints(t *testing.T) {
	t.Log(ThroughMostPoints([][]int{{0, 0}, {0, 1}}))
	t.Log(ThroughMostPoints([][]int{{2, 3}, {3, 3}, {-5, 3}}))
}

//a -> b -> c -> d -> e -> f -> nil

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func LinkSort(root *LinkNode) *LinkNode {
	//分割
	if root.Next == nil {
		return root
	}
	var pre, next = root, root.Next
	for ; next != nil && next.Next != nil; {
		pre = pre.Next
		next = next.Next
		if next != nil {
			next = next.Next
		}
	}
	nextRoot := pre.Next
	pre.Next = nil
	root = LinkSort(root)         //前半部分连表排序
	nextRoot = LinkSort(nextRoot) //后半部分连表排序
	//归并
	var newRoot, curr, preLink, nextLink *LinkNode
	preLink = root
	nextLink = nextRoot
	for {
		if preLink.Val <= nextLink.Val {
			if newRoot == nil {
				newRoot = preLink
				curr = preLink
			} else {
				curr.Next = preLink
				curr = curr.Next
			}
			preLink = preLink.Next
			if preLink == nil {
				curr.Next = nextLink
				break
			}
		} else {
			if newRoot == nil {
				newRoot = nextLink
				curr = nextLink
			} else {
				curr.Next = nextLink
				curr = curr.Next
			}
			nextLink = nextLink.Next
			if nextLink == nil {
				curr.Next = preLink
				break
			}
		}
	}
	return newRoot
}

func TestLinkSort(t *testing.T) {
	root := &LinkNode{Val: 1, Next: &LinkNode{Val: 3, Next: &LinkNode{Val: 2, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 1}}}}}
	root = LinkSort(root)
	var curr = root
	for ; curr != nil; curr = curr.Next {
		t.Log(curr.Val)
	}
}

//1->3->2->4->2->nil
func InsertSort(root *LinkNode) *LinkNode {
	var endNode, minNode, minNodePre, curr, pre *LinkNode
	endNode = &LinkNode{Next: root}
	root = endNode
	for endNode.Next != nil {
		curr = endNode.Next
		pre = endNode
		minNode = curr
		minNodePre = pre
		for curr != nil {
			if minNode.Val > curr.Val {
				minNode = curr
				minNodePre = pre
			}
			pre = curr
			curr = curr.Next
		}
		//插入元素
		minNodePre.Next = minNode.Next
		minNode.Next = endNode.Next
		endNode.Next = minNode
		endNode = endNode.Next
	}
	return root.Next
}

func TestInsertSort(t *testing.T) {
	root := &LinkNode{Val: 1, Next: &LinkNode{Val: 3, Next: &LinkNode{Val: 2, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 1}}}}}
	root = InsertSort(root)
	var curr = root
	for ; curr != nil; curr = curr.Next {
		t.Log(curr.Val)
	}
}

//1->2->3->4->5->6->nil
func TransformLink(root *LinkNode) *LinkNode {
	var list []*LinkNode
	for curr := root; curr != nil; curr = curr.Next {
		list = append(list, curr)
	}
	for i := 0; i < len(list)/2; i++ {
		if len(list)%2 == 0 && i == len(list)/2-1 {
			break
		}
		list[len(list)-i-1].Next = list[i].Next
		list[len(list)-i-2].Next = nil
		list[i].Next = list[len(list)-i-1]
	}
	return root
}

func TestTransformLink(t *testing.T) {
	root := &LinkNode{Val: 1, Next: &LinkNode{Val: 3, Next: &LinkNode{Val: 2, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 1}}}}}
	root = TransformLink(root)
	var curr = root
	for ; curr != nil; curr = curr.Next {
		t.Log(curr.Val)
	}
}

// s + c = l
/**
s = 0     2s - s = s
c - (c - s % c) = s % c
s > c
delta = s % c = s - n * c
s = delta + n * c
s <= c
s = delta
*/

func GetCircleNode(root *LinkNode) *LinkNode {
	var node = root
	var nodeSet = make(map[*LinkNode]bool)
	for ; node != nil && !nodeSet[node]; node = node.Next {
		nodeSet[node] = true
	}
	return node
}

func TestGetCircleNode(t *testing.T) {
	node2 := &LinkNode{Val: 2}
	node3 := &LinkNode{Val: 3, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 5, Next: node2}}}
	node2.Next = node3
	t.Log(GetCircleNode(&LinkNode{Val: 1, Next: node2}))
}

func HasCircle(root *LinkNode) bool {
	if root == nil || root.Next == nil {
		return false
	}
	var slow, fast = root, root.Next
	for ; fast != nil && slow != fast; {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow == fast
}

func TestHasCircle(t *testing.T) {
	node2 := &LinkNode{Val: 2}
	node3 := &LinkNode{Val: 3, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 5, Next: node2}}}
	node2.Next = node3
	t.Log(HasCircle(&LinkNode{Val: 1, Next: node2}))

	t.Log(HasCircle(&LinkNode{Val: 1, Next: &LinkNode{Val: 2, Next: &LinkNode{Val: 3, Next: &LinkNode{Val: 4, Next: &LinkNode{Val: 5}}}}}))
}

//dp[i][j] = dp[i][j-1] && contains(list[:i], str[j-1:j]) || dp[i-1][j]
func CanCombine(str string, word []string) bool {
	var dp = make([][]bool, len(word)+1)
	for i := 0; i < len(word)+1; i++ {
		dp[i] = make([]bool, len(str)+1)
		dp[i][0] = true
	}
	var wordSet = make(map[string]bool)
	for i := 0; i < len(word); i++ {
		wordSet[word[i]] = true
	}
	for i := 1; i < len(word)+1; i++ {
		for j := 1; j < len(str)+1; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			} else {
				for k := 1; k <= j; k++ {
					if wordSet[str[j-k:j]] && dp[i][j-k] {
						dp[i][j] = true
						break
					}
				}
			}
		}
	}
	return dp[len(word)][len(str)]
}

func TestCanCombine(t *testing.T) {
	t.Log(CanCombine("nowcode", []string{"now", "code"}))
}

func GetSingleNumber(arr []int) int {
	var singleNumber int
	for i := 0; i < 32; i++ {
		var bitCount int
		for j := 0; j < len(arr); j++ {
			bitCount += (arr[j] >> i) & 1
		}
		if bitCount%3 != 0 { //该位数孤立的数为1
			singleNumber = singleNumber | (1 << i)
		}
	}
	return singleNumber
}

func TestGetSingleNumber(t *testing.T) {
	t.Log(GetSingleNumber([]int{0, 0, 0, 5}))
}

//arr[i-1]  arr[i]
//arr[i-1] > arr[i]
//cost[i] = 1, cost[i-1]+1 cost[i-2]+1 ... util cost[i-n] > cost[i-n-1]
//arr[i-1] <= arr[i]
//cost[i] = cost[i-1]+1
//arr[i-1] == arr[i]
//cost[i] = cost[i-1]
func DivideCandy(scores []int) int {
	var count, preCandyCount int
	for i := 0; i < len(scores); i++ {
		if i == 0 {
			count = 1
			preCandyCount = 1
		} else {
			if scores[i] == scores[i-1] {
				count += preCandyCount
			} else if scores[i] < scores[i-1] {
				count += 1
				if preCandyCount >= 1 {
					for j := 1; j <= i && scores[i-j] >= scores[i-j+1]; j++ {
						count++
					}
				}
				preCandyCount = 1
			} else { // scores[i] > scores[i-1]
				count += preCandyCount + 1
				preCandyCount++
			}
		}
	}
	return count
}

func TestDivideCandy(t *testing.T) {
	t.Log(DivideCandy([]int{3, 2, 1}))
}

func FindFirstStation(oils []int, costs []int) int {
	var currOil, curr int
	for i := 0; i < len(oils); i++ {
		currOil = oils[i]
		curr = i
		for step := 1; step <= len(oils); step++ {
			if currOil >= costs[curr] {
				currOil = currOil - costs[curr]
				curr = (i + step) % len(oils)
				if curr == i {
					return i
				}
				currOil += oils[curr]
			} else {
				break
			}
		}
	}
	return -1
}

func NewFindFirstStation(oils []int, costs []int) int {
	var start, currOil, total int
	for i := 0; i < len(oils); i++ {
		total += oils[i] - costs[i]
		currOil += oils[i] - costs[i]
		if currOil < 0 {
			currOil = 0
			start = i + 1
		}
	}
	if total < 0 || start == len(oils) {
		return -1
	}
	return start
}

func TestFindFirstStation(t *testing.T) {
	t.Log(FindFirstStation([]int{2, 3, 1}, []int{3, 1, 2}))
	t.Log(NewFindFirstStation([]int{2, 3, 1}, []int{3, 1, 2}))
}

type GraphNode struct {
	Val int
	Nbs []*GraphNode
}

func MarshalGraph(node *GraphNode) string {
	var result = "{"
	var queue []*GraphNode
	var nodeSet = make(map[*GraphNode]bool)
	queue = append(queue, node)
	for len(queue) > 0 {
		nodeSet[queue[0]] = true
		if result == "{" {
			result = fmt.Sprintf("%d", queue[0].Val)
		} else {
			result += fmt.Sprintf("#%d", queue[0].Val)
		}
		for i := 0; i < len(queue[0].Nbs); i++ {
			result += fmt.Sprintf(",%d", queue[0].Nbs[i].Val)
			if !nodeSet[queue[0].Nbs[i]] {
				queue = append(queue, queue[0].Nbs[i])
			}
		}
		queue = queue[1:]
	}
	result += "}"
	return result
}

//dd/efed
//d d e f e d
//dd e f e d
func GetSubStrList(str string) [][]string {
	var result [][]string
	for i := 0; i < len(str); i++ {
		ls := str[:i+1]
		rs := str[i+1:]
		var flag bool
		//判断ls是否
		for j := 0; j < len(ls)/2; j++ {
			if ls[j] != ls[len(ls)-j-1] {
				flag = true
				break
			}
		}
		if !flag {
			list := GetSubStrList(rs)
			if len(list) > 0 {
				for j := 0; j < len(list); j++ {
					result = append(result, append([]string{ls}, list[j]...))
				}
			} else {
				result = append(result, []string{ls})
			}
		}
	}
	return result
}

func TestGetSubStrList(t *testing.T) {
	t.Log(GetSubStrList("ddeefese"))
}

func GetMinSubStrList(str string) []string {
	var minResult []string
	var minLs string
	for i := 0; i < len(str); i++ {
		ls := str[:i+1]
		rs := str[i+1:]
		var flag bool
		//判断ls是否
		for j := 0; j < len(ls)/2; j++ {
			if ls[j] != ls[len(ls)-j-1] {
				flag = true
				break
			}
		}
		if !flag {
			list := GetSubStrList(rs)
			if len(list) > 0 {
				for j := 0; j < len(list); j++ {
					if len(minResult) == 0 || len(list[j]) < len(minResult) {
						minResult = list[j]
						minLs = ls
					}
				}
			} else {
				return []string{ls}
			}
		}
	}
	return append([]string{minLs}, minResult...)
}

func TestGetMinSubStrList(t *testing.T) {
	t.Log(GetMinSubStrList("ddeefese"))
}

func FlagO(arr [][]byte) {
	var flags [][]bool //记录所有被标记的O
	var flag func(arr [][]byte, i, j int)
	flag = func(arr [][]byte, i, j int) { //标记函数
		flags[i][j] = true
		//左边
		if j > 0 && arr[i][j-1] == 'O' && !flags[i][j-1] {
			flag(arr, i, j-1)
		}
		//上边
		if i > 0 && arr[i-1][j] == 'O' && !flags[i-1][j] {
			flag(arr, i-1, j)
		}
		//右边
		if j < len(arr[0])-1 && arr[i][j+1] == 'O' && !flags[i][j+1] {
			flag(arr, i, j+1)
		}
		//下边
		if i < len(arr)-1 && arr[i+1][j] == 'O' && !flags[i+1][j] {
			flag(arr, i+1, j)
		}
	}
	//左边界查找O
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == 'O' && !flags[i][0] {
			//深度优先标记O
			flag(arr, i, 0)
		}
	}
	//右边界查找O
	for i := 0; i < len(arr); i++ {
		if arr[i][len(arr[0])-1] == 'O' && !flags[i][len(arr[0])-1] {
			//深度优先标记O
			flag(arr, i, len(arr[0])-1)
		}
	}
	//上边界查找O
	for i := 1; i < len(arr[0])-1; i++ {
		if arr[0][i] == 'O' && !flags[0][i] {
			//深度优先标记O
			flag(arr, 0, i)
		}
	}
	//下边界查找O
	for i := 1; i < len(arr[0])-1; i++ {
		if arr[len(arr)-1][i] == 'O' && !flags[len(arr)-1][i] {
			//深度优先标记O
			flag(arr, len(arr)-1, i)
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 'O' && !flags[i][j] {
				arr[i][j] = 'X'
			}
		}
	}
}

func GetPathsOfVal(root *Node, val int) [][]*Node {
	var results [][]*Node
	if root.Left == nil && root.Right == nil && root.Val == val {
		results = append(results, []*Node{root})
		return results
	}
	var paths [][]*Node
	if root.Left != nil {
		paths = append(paths, GetPathsOfVal(root.Left, val-root.Val)...)
	}
	if root.Right != nil {
		paths = append(paths, GetPathsOfVal(root.Right, val-root.Val)...)
	}
	if len(paths) > 0 {
		for i := 0; i < len(paths); i++ {
			results = append(results, append([]*Node{root}, paths[i]...))
		}
	}
	return results
}

func IsBalanceBTree(root *Node) (height int) {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lh := IsBalanceBTree(root.Left)
	if lh < 0 {
		return lh
	}
	rh := IsBalanceBTree(root.Right)
	if rh < 0 {
		return rh
	}
	if math.Abs(float64(lh)-float64(rh)) <= 1 {
		return int(math.Max(float64(lh), float64(rh)))
	} else {
		return -1
	}
}

func CreateBalanceBTreeByArrayList(list []int) *Node {
	if len(list) <= 0 {
		return nil
	}
	mid := len(list) / 2
	root := &Node{Val: list[mid]}
	if mid > 0 {
		root.Left = CreateBalanceBTreeByArrayList(list[:mid])
	}
	if mid < len(list)-1 {
		root.Right = CreateBalanceBTreeByArrayList(list[mid+1:])
	}
	return root
}

func CreateBSTByLinkList(root *LinkNode) *Node {
	if root.Next == nil {
		return &Node{Val: root.Val}
	}
	var s, f = root, root.Next
	var preS *LinkNode
	//找出中点
	for f != nil {
		preS = s
		s = s.Next
		f = f.Next
		if f != nil {
			f = f.Next
		}
	}
	TRoot := &Node{Val: s.Val}
	if preS != nil {
		preS.Next = nil
	}
	TRoot.Left = CreateBSTByLinkList(root)
	TRoot.Right = CreateBSTByLinkList(s.Next)
	return TRoot
}

func GetTree(preList, midList []int) *Node {
	if len(preList) <= 0 || len(midList) <= 0 {
		return nil
	}
	root := &Node{Val: preList[0]}
	for i := 0; i < len(midList); i++ {
		if midList[i] == preList[0] {
			root.Left = GetTree(preList[1:i+1], midList[0:i])
			root.Right = GetTree(preList[i+1:], midList[i+1:])
		}
	}
	return root
}

func ZList(root *Node) []int {
	var stack1, stack2 []*Node
	var result []int
	if root == nil {
		return result
	}
	stack1 = append(stack1, root)
	for len(stack1) > 0 || len(stack2) > 0 {
		for len(stack1) > 0 {
			top := stack1[len(stack1)-1]
			result = append(result, top.Val)
			stack2 = append(stack2, top.Left, top.Right)
			stack1 = stack1[:len(stack1)-1]
		}
		for len(stack2) > 0 {
			top := stack2[len(stack2)-1]
			result = append(result, top.Val)
			stack1 = append(stack1, top.Left, top.Right)
			stack2 = stack2[:len(stack2)-1]
		}
	}
	return result
}

func IsMirrorTree(root *Node) bool {
	if root == nil {
		return true
	}
	var isMirrorTree func(root1, root2 *Node) bool
	isMirrorTree = func(root1, root2 *Node) bool {
		if root1 == nil && root2 == nil {
			return true
		} else if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val == root2.Val {
			return isMirrorTree(root1.Left, root2.Right) && isMirrorTree(root1.Right, root2.Left)
		}
		return false
	}
	return isMirrorTree(root.Left, root.Right)
}

func GetIPs(seq string, level int) []string {
	var res []string
	if len(seq) == 0 {
		return nil
	}
	if level < 4 {
		//从seq中取1,2,3个字符作为ip分段
		if seq[0] == '0' {
			subRes := GetIPs(seq[1:], level+1)
			if len(subRes) > 0 {
				for i := 0; i < len(subRes); i++ {
					res = append(res, "0."+subRes[i])
				}
			}
		} else {
			for i := 0; i < 3 && i < len(seq); i++ {
				parseInt, _ := strconv.ParseInt(seq[0:i+1], 10, 64)
				if parseInt <= 255 {
					subRes := GetIPs(seq[i+1:], level+1)
					if len(subRes) > 0 {
						for j := 0; j < len(subRes); j++ {
							res = append(res, seq[0:i+1]+"."+subRes[j])
						}
					}
				} else {
					break
				}
			}
		}
	} else if level == 4 && len(seq) <= 3 && (len(seq) == 1 || seq[0] != '0') {
		parseInt, _ := strconv.ParseInt(seq, 10, 64)
		if parseInt <= 255 {
			res = []string{seq}
		}
	}
	return res
}

func TestGetIPs(t *testing.T) {
	t.Log(GetIPs("1111", 1))
}

func ReverseLink(head *LinkNode, m, n int) *LinkNode {
	if n-m == 0 {
		return head
	}
	var preLNode, lNode, rNode *LinkNode
	//找到m, n位置节点
	var pre, curr = head, head
	for i := 0; i < n && curr != nil; i++ {
		if i == m-1 {
			preLNode = pre
			lNode = curr
		}
		curr = curr.Next
	}
	if curr == nil {
		return nil
	}
	rNode = curr
	//翻转连表
	var preReverseNode, currReverseNode, nextReverseNode *LinkNode = nil, lNode, lNode.Next
	for ; currReverseNode != rNode.Next; {
		currReverseNode.Next = preReverseNode
		preReverseNode = currReverseNode
		currReverseNode = nextReverseNode
		nextReverseNode = nextReverseNode.Next
	}
	preLNode.Next = rNode
	lNode.Next = rNode.Next
	return head
}

func MergeSortedLink(head1, head2 *LinkNode) *LinkNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}
	var head, curr *LinkNode
	for node1, node2 := head1, head2; node1 != nil || node2 != nil; {
		if node1 == nil {
			curr.Next = node2
			break
		} else if node2 == nil {
			curr.Next = node1
			break
		} else {
			if node1.Val <= node2.Val {
				curr.Next = node1
				node1 = node1.Next
			} else {
				curr.Next = node2
				node2 = node2.Next
			}
		}
	}
	return head
}

//1 2 4 3 2 5 -> 1 2 2 4 3 5
func MoveLinkNode(head *LinkNode, x int) *LinkNode {
	if head == nil {
		return nil
	}
	var lHead, rHead *LinkNode
	//从链表中找出大于或等于x元素的第一的节点
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val >= x {
			rHead = curr
			if curr == head {
				lHead = nil
			} else {
				lHead = head
			}
			break
		}
	}
	if rHead == nil {
		return head
	}
	//从rHead链表中找出所有比x小的元素，并从链表中删除，然后插入到lHead链表中
	var pre, tail *LinkNode
	for curr := rHead; curr != nil; {
		if curr.Val < x {
			pre.Next = curr.Next
			if lHead == nil {
				lHead = curr
				tail = curr
			} else {
				tail.Next = curr
				tail = tail.Next
			}
			curr = curr.Next
		} else {
			pre = curr
			curr = curr.Next
		}
	}
	tail.Next = rHead
	return lHead
}

func DelRepeatItem(link string) string {
	var head, tail *LinkNode
	var curr string
	for i := 0; i < len(link); i++ {
		if link[i] >= '0' && link[i] <= '9' {
			curr += fmt.Sprintf("%c", link[i])
		} else if link[i] == ',' || link[i] == '}' {
			parseInt, err := strconv.ParseInt(curr, 10, 64)
			if err == nil {
				if head == nil {
					head = &LinkNode{Val: int(parseInt)}
					tail = head
				} else {
					tail.Next = &LinkNode{Val: int(parseInt)}
					tail = tail.Next
				}
			}
			curr = ""
		}
	}
	if head.Next == nil {
		return fmt.Sprintf("{%d}", head.Val)
	}
	var l, r = head, head.Next
	for ; r != nil; {
		if l.Val != r.Val {
			if l.Next != r {
				l.Next = r
				l = r
				r = r.Next
			} else {
				l = l.Next
				r = r.Next
			}
		} else {
			r = r.Next
			if r == nil {
				l.Next = nil
				break
			}
		}
	}
	var res = ""
	for currNode := head; currNode != nil; currNode = currNode.Next {
		res += fmt.Sprintf("%d", currNode.Val)
		if currNode.Next != nil {
			res += ","
		}
	}
	return "{" + res + "}"
}

func TestDelRepeatItem(t *testing.T) {
	t.Log(DelRepeatItem("{1,1,2,2,3,4}"))
}

func DelRepeatItem2(link string) string {
	var head, tail *LinkNode
	var curr string
	for i := 0; i < len(link); i++ {
		if link[i] >= '0' && link[i] <= '9' {
			curr += fmt.Sprintf("%c", link[i])
		} else if link[i] == ',' || link[i] == '}' {
			parseInt, err := strconv.ParseInt(curr, 10, 64)
			if err == nil {
				if head == nil {
					head = &LinkNode{Val: int(parseInt)}
					tail = head
				} else {
					tail.Next = &LinkNode{Val: int(parseInt)}
					tail = tail.Next
				}
			}
			curr = ""
		}
	}
	if head.Next == nil {
		return fmt.Sprintf("{%d}", head.Val)
	}
	var l, r = head, head.Next
	var pre *LinkNode
	for ; r != nil; {
		if l.Val != r.Val {
			if l.Next != r {
				if pre == nil {
					head = r
				} else {
					pre.Next = r
				}
				l = r
				r = r.Next
			} else {
				pre = l
				l = r
				r = r.Next
			}
		} else {
			r = r.Next
			if r == nil {
				if pre == nil {
					return "{}"
				} else {
					pre.Next = nil
				}
			}
		}
	}
	var res = ""
	for currNode := head; currNode != nil; currNode = currNode.Next {
		res += fmt.Sprintf("%d", currNode.Val)
		if currNode.Next != nil {
			res += ","
		}
	}
	return "{" + res + "}"
}

func TestDelRepeatItem2(t *testing.T) {
	t.Log(DelRepeatItem2("{1,1,2,2,3,4,5,5}"))
}

func GetSubSets(input string) string {
	if len(input) <= 2 {
		return "[[]]"
	}
	var list []int
	inputList := strings.Split(input[1:len(input)-1], ",")
	for _, s := range inputList {
		i, _ := strconv.ParseInt(s, 10, 64)
		list = append(list, int(i))
	}
	var results [][]int
	sort.Ints(list)
	for bitmap := 0; bitmap <= 1<<len(list)-1; bitmap++ {
		var set []int
		bits := bitmap
		for i := 0; i < len(list); i++ {
			if bits&1 > 0 {
				set = append(set, list[i])
			}
			bits >>= 1
		}
		results = append(results, set)
	}
	var res string
	for i := 0; i < len(results); i++ {
		var set string
		for j := 0; j < len(results[i]); j++ {
			if j < len(results[i])-1 {
				set += fmt.Sprintf("%d,", results[i][j])
			} else {
				set += fmt.Sprintf("%d", results[i][j])
			}
		}
		set = "[" + set + "]"
		res += set
		if i < len(results)-1 {
			res += ","
		}
	}
	return "[" + res + "]"
}

func TestGetSubSets(t *testing.T) {
	t.Log(GetSubSets("[1,2]"))
}

//XYDxGZZisiegYX
//XYZ
func GetMinSubSeq(S, T string) string {
	if len(S) < len(T) {
		return ""
	}
	var TSet = make(map[byte]bool)
	for i := 0; i < len(T); i++ {
		TSet[T[i]] = false
	}
	var l, r = 0, math.MaxInt64
	for i := 0; i < len(S); i++ {
		if _, ok := TSet[S[i]]; ok {
			for k := range TSet {
				TSet[k] = false
			}
			var j, count int
			for j = i; j < len(S)-i && count < len(TSet); j++ {
				if v, ok := TSet[S[j]]; ok && !v {
					TSet[S[j]] = true
					count++
				}
			}
			if count >= len(TSet) && r-l > j-i {
				l = i
				r = j - 1
			}
		}
	}
	if r == math.MaxInt64 {
		return ""
	}
	return S[l : r+1]
}

func TestGetMinSubSeq(t *testing.T) {
	t.Log(GetMinSubSeq("XYMTZXIHIGY", "XYZ"))
}

//y = x^2 - C
func Sqrt(c float64) float64 {
	if c < 0 {
		return -1
	}
	if c == 0 {
		return 0
	}
	var l, r float64 = 0, c
	var mid = (l + r) / 2
	for y := mid*mid - c; math.Abs(y) > 0.0001; y = mid*mid - c {
		if y > 0 {
			r = mid
		} else if y < 0 {
			l = mid
		} else {
			return mid
		}
		mid = (l + r) / 2
	}
	return mid
}

func TestSqrt(t *testing.T) {
	t.Log(Sqrt(2))
}

func TotalPathCount(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return TotalPathCount(m-1, n) + TotalPathCount(m, n-1)
}

func TestTotalPathCount(t *testing.T) {
	t.Log(TotalPathCount(2, 2))
}

func MergeInterval(input string) string {
	if len(input) <= 2 {
		return "[]"
	}
	input = input[1 : len(input)-1]
	var intervalList [][]int
	var l, r int
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			l = i
		} else if input[i] == ']' {
			r = i
			splits := strings.Split(input[l+1:r], ",")
			ln, _ := strconv.ParseInt(splits[0], 10, 64)
			rn, _ := strconv.ParseInt(splits[1], 10, 64)
			intervalList = append(intervalList, []int{int(ln), int(rn)})
		}
	}
	var output string
	sort.Slice(intervalList, func(i, j int) bool {
		return intervalList[j][0] > intervalList[i][0]
	})
	var result [][]int
	l, r = intervalList[0][0], intervalList[0][1]
	for i := 1; i < len(intervalList); i++ {
		if intervalList[i][0] >= l && intervalList[i][0] <= r && intervalList[i][1] >= r {
			r = intervalList[i][1]
		} else if intervalList[i][0] > r {
			result = append(result, []int{l, r})
			l = intervalList[i][0]
			r = intervalList[i][1]
		}
	}
	result = append(result, []int{l, r})
	for i := 0; i < len(result); i++ {
		output += fmt.Sprintf("[%d,%d]", result[i][0], result[i][1])
		if i < len(result)-1 {
			output += ","
		}
	}
	return "[" + output + "]"
}

func TestMergeInterval(t *testing.T) {
	t.Log(MergeInterval("[[0,3],[2,4],[4,5],[6,8],[8,8],[9,11]]"))
}

//1 2 3
//4 5 6
//7 8 9
func ListMatrix(input string) string {
	if input == "[]" {
		return "[]"
	}
	var matrix [][]int
	var row []int
	var l, r int
	input = input[1 : len(input)-1]
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			l = i
		} else if input[i] == ']' {
			r = i
			split := strings.Split(input[l+1:r], ",")
			row = nil
			for j := 0; j < len(split); j++ {
				n, _ := strconv.ParseInt(split[j], 10, 64)
				row = append(row, int(n))
			}
			matrix = append(matrix, row)
		}
	}
	var result []int
	var x, y int
	for m, n := len(matrix), len(matrix[0]); m > 0 && n > 0; {
		//遍历上边
		for i := y; i <= y+n-1; i++ {
			result = append(result, matrix[x][i])
		}
		//遍历右边
		if m > 1 {
			for i := x + 1; i <= x+m-1; i++ {
				result = append(result, matrix[i][y+n-1])
			}
		} else {
			break
		}
		//遍历下边
		if n > 1 {
			for i := y + n - 2; i >= y; i-- {
				result = append(result, matrix[x+m-1][i])
			}
		} else {
			break
		}
		//遍历左边
		for i := x + m - 2; i >= x+1; i-- {
			result = append(result, matrix[i][y])
		}
		m -= 2
		n -= 2
		x++
		y++
	}
	var output string
	for i := 0; i < len(result); i++ {
		output += fmt.Sprintf("%d", result[i])
		if i < len(result)-1 {
			output += ","
		}
	}
	return "[" + output + "]"
}

func TestListMatrix(t *testing.T) {
	t.Log(ListMatrix("[[1,2,3],[4,5,6],[7,8,9]]"))
}

func Permutations(input string) string {
	if len(input) <= 2 {
		return "[]"
	}
	arr := strings.Split(input[1:len(input)-1], ",")
	var f func(arr []string) [][]string
	f = func(arr []string) (res [][]string) {
		//选择一个元素
		for i := 0; i < len(arr); i++ {
			if arr[i] != "" {
				n := arr[i]
				arr[i] = ""
				subRes := f(arr)
				if subRes == nil {
					res = append(res, []string{n})
				} else {
					for j := 0; j < len(subRes); j++ {
						res = append(res, append(subRes[j], n))
					}
				}
				arr[i] = n
			}
		}
		return
	}
	res := f(arr)
	var output string
	for i := 0; i < len(res); i++ {
		output += fmt.Sprintf("[%s]", strings.Join(res[i], ","))
		if i < len(res)-1 {
			output += ","
		}
	}
	return "[" + output + "]"
}

func TestPermutations(t *testing.T) {
	t.Log(Permutations("[1,2,3]"))
}

func RepeatPermutations(input string) string {
	if len(input) <= 2 {
		return "[]"
	}
	arr := strings.Split(input[1:len(input)-1], ",")
	var f func(arr []string) (res [][]string)
	f = func(arr []string) (res [][]string) {
		var set = make(map[string]struct{})
		//选择一个不重复的元素
		for i := 0; i < len(arr); i++ {
			_, ok := set[arr[i]]
			if arr[i] != "" && !ok {
				n := arr[i]
				arr[i] = ""
				subRes := f(arr)
				if subRes == nil {
					res = append(res, []string{n})
				} else {
					for j := 0; j < len(subRes); j++ {
						res = append(res, append(subRes[j], n))
					}
				}
				arr[i] = n
				set[n] = struct{}{}
			}
		}
		return
	}
	res := f(arr)
	var output string
	for i := 0; i < len(res); i++ {
		output += fmt.Sprintf("[%s]", strings.Join(res[i], ","))
		if i < len(res)-1 {
			output += ","
		}
	}
	return "[" + output + "]"
}

func TestRepeatPermutations(t *testing.T) {
	t.Log(RepeatPermutations("[1,2,3,3]"))
}

func IsMatch(s, p string) bool {
	var dp [][]bool
	var flag bool
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			flag = true
			break
		}
	}
	for i := 0; i < len(s)+1; i++ {
		row := make([]bool, len(p)+1)
		if i == 0 {
			row[0] = true
			if !flag {
				for j := 1; j < len(row)+1; j++ {
					row[j] = true
				}
			}
		}
		dp = append(dp, row)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '*' || p[j-1] == '?') || dp[i-1][j] && p[j-1] == '*' || dp[i][j-1] && p[j-1] == '*'
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestIsMatch(t *testing.T) {
	t.Log(IsMatch("aa", "a*"))
}

