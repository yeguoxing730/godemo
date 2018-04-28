package _type
import (
	"fmt"
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"reflect"
)

func  ArrayTest(){
	var arrs [5]int = [...]int{11,12,13,14,15};
	fmt.Print(arrs)
	for  i :=0; i<len(arrs);i++  {
		fmt.Println(arrs[i])
	}
	for index,key := range  arrs {
		fmt.Println(key)
		fmt.Println("key:"+string(key))
		fmt.Printf("arr[%d]=%d\n", index, key)
		//fmt.Println("index:"+string(index))
		fmt.Printf("key: %d,index:%d",key,index)
	}
	fmt.Printf("|%0+- #[1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 8, 4, 32, 64)
	fmt.Printf("|%f|%8.4f|%8.f|%.4f|%.f|\n", 3.2, 3.2, 3.2, 3.2, 3.2)
	fmt.Printf("|%.3f|%.3g|\n", 12.345678, 12.345678)
	fmt.Printf("|%.2f|\n", 12.345678+12.345678i)
	modify(arrs)
	fmt.Println("In main(), arr values:", arrs)
}
func MapTest(){
	var b map[int]string  = make(map[int]string);
	b[1] = "first field"
	b[2] = "second field"
	b[3] = "third"
	fmt.Println(b)
	b[2] ="change field"
	for value,key := range b{
		fmt.Println(key,value)
	}
	delete(b, 3);
	var val,ok = b[2];
	if ok {
		fmt.Println(ok)
		fmt.Println(val)
		fmt.Println("拥有数据：" )
	}
	fmt.Println(b)
}
func modify(arr [5]int){
	arr[0] = 10
	fmt.Println("In modify(), arr values:", arr)
}

type Person struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}

func TestStructToJSON()  {
	var p = Person{
		Name:"liyang",
		Age:29,
	}
	fmt.Println(p)
	jsonBytes,err := json.Marshal(p)
	if err != nil{
		fmt.Println("转换json对象时出错")
		panic(err)
	}
	fmt.Println("转换的json对象为："+string(jsonBytes))
}

func TestJSONToStruct()  {
	var jsonStr = `{"name":"liyang","age":29}`
	var p Person;
	var err = json.Unmarshal([]byte(jsonStr),&p)
	if err != nil{
		fmt.Println("convert error")
	}
	fmt.Println(p)
}
func JSONToMap(){
	var jsonStr = `{"name":"liyang","age":29}`
	var mapres = make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr),&mapres)
	if err != nil{
		fmt.Println("convert fatal")
	}
	fmt.Println(mapres)
}
func MapToJson()  {
	var mapVar = make(map[string]interface{});
	mapVar["name"] = "sss"
	mapVar["age"] = 12
	mapVar["address"] = "user address"
	jsonBytes,err := json.Marshal(mapVar)
	if err != nil{
		fmt.Println("convert fatal")
	}
	fmt.Println("map to json result:"+string(jsonBytes))

}
func MapToStruct(){
	var mapVar = make(map[string]interface{});
	mapVar["name"] = "小米"
	mapVar["age"] = 12
	mapVar["address"] = "user address"
	var person Person
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(mapVar, &person); err != nil {
		fmt.Println("convert fatal")
	}
	fmt.Println("map2struct后得到的 struct 内容为:%v", person)
}
func StructToMap(obj interface{}) map[string]interface{}{
	t:= reflect.TypeOf(obj)
	v:= reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}