package stringhelper

import (
	"strconv"
	"fmt"
)
func Test(){
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 255
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	// 255
	fmt.Println(strconv.ParseInt("9", 10, 4))
	// 7 strconv.ParseInt: parsing "9": value out of range
	//解析字符串为10进制的数字9
	fmt.Println(strconv.Atoi("9"))
	//解析数字9为字符串的"9"
	fmt.Println(strconv.Itoa(9))
	// 7 strconv.ParseInt: parsing "9": value out of range
	fmt.Println(strconv.ParseInt("9", 10, 4))
	// 7 strconv.ParseInt: parsing "9": value out of range
 /**
 **FormatFloat 将浮点数 f 转换为字符串形式
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64），结果会据此进行舍入。
	//
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	//
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	// 参考格式化输入输出中的旗标和精度说明
	func FormatFloat(f float64, fmt byte, prec, bitSize int) string

	// 将字符串解析为浮点数，使用 IEEE754 规范进行舍入。
	// bigSize 取值有 32 和 64 两种，表示转换结果的精度。
	// 如果有语法错误，则 err.Error = ErrSyntax
	// 如果结果超出范围，则返回 ±Inf，err.Error = ErrRange
	func ParseFloat(s string, bitSize int) (float64, error)
  */
	s := "0.12345678901234567890"

	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)                // 0.12345679104328156
	fmt.Println(float32(f), err)       // 0.12345679

	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err)                // 0.12345678901234568
}
