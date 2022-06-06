package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmoucse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println(m)
	//	map[course:golang name:ccmoucse quality:notbad site:imooc]

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil

	fmt.Println(m, m2, m3)
	// map[course:golang name:ccmoucse quality:notbad site:imooc] map[] map[]

	fmt.Println("Traversing map") // map遍历

	for k, v := range m {
		fmt.Println(k, v)
	}
	// name ccmoucse
	// course golang
	// site imooc
	// quality notbad

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	//golang

	caurseName := m["caurse"]
	fmt.Println(caurseName)
	//	输出为空，拿到的是空值（zero value）,go语言不怕我们变量不初始化

	//	判断key是否在map里存在与否(通过使用额外一个变量去接收key的存在状态，就能通过这个变量判断key在map中的存在与否）
	courseName2, ok := m["course"]
	fmt.Println(courseName2, ok)
	//golang true

	ccurseName, ok := m["ccurse"]
	fmt.Println(ccurseName, ok)
	//	false

	// 先判断key是否存在，存在就将值输出
	if ccc, ok := m["course"]; ok {
		fmt.Println(ccc)
	} else {
		fmt.Println("key dose not exist")
	}
	//	golang

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	//ccmoucse true
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	//false

	if name, ok = m["name"]; ok {
		delete(m, name)
	} else {
		fmt.Println("no key")
	}

}
