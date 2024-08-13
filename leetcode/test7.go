package main

type ReqMsg struct {
	Id, Uid int // 必须大写开头导出变量
	Str     string
}

// func main() {
// 	fmt.Println("=== main test7 begin == ")

// 	var buf bytes.Buffer
// 	encoder := gob.NewEncoder(&buf)

// 	req := ReqMsg{
// 		Id:  1,
// 		Uid: 10001,
// 		Str: "lord10001",
// 	}
// 	err := encoder.Encode(req)
// 	if err != nil {
// 		fmt.Println("test7 Encode err", err)
// 		return
// 	}

// 	r := &ReqMsg{}
// 	decoder := gob.NewDecoder(&buf)
// 	err = decoder.Decode(r)
// 	if err != nil {
// 		fmt.Println("test7 encode err", err)
// 		return
// 	}
// 	fmt.Println("decode sucess=", r)

// 	fmt.Println("=== main test7 end == ")
// }
