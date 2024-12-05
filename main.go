package main

import (
	"fmt"
	"sort"
)

func main() {
	//测试net/http::RoundTrip的hook效果
	//req, err := http.NewRequestWithContext(context.Background(), "GET", "http://www.baidu.com", nil)
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Set("kotelbuild", "true")
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//defer resp.Body.Close()

	//测试service::CreateUser的hook效果
	//u, _ := service.CreateUser(context.Background(), 1, "test")
	//
	////model.UserModel注入了新字段age
	//fmt.Println(fmt.Sprintf("UserModel: %+v", u))

	//替换排序算法实现
	nums := []int{5, 2, 7, 1, 8, 3, 4, 6}
	// 使用sort包进行排序
	sort.Ints(nums)
	fmt.Println("Sorted nums:", nums)

}
