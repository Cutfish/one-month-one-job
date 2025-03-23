### http
任何实现了 ServeHTTP 方法的对象都可以作为 HTTP 的 Handler。
```Go
type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello World!"))
}

func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
```

**panic和error**
panic适用于不能恢复的错误（recover可以截取），error是可以恢复的错误，一个是完全错误，一个业务错误

**type...和...type**
...type表示可变参数，type...就是对切片展开
```Go
func (m *Map) Add(keys ...string) {
	fmt.Println(keys) // keys 是一个 []string 切片
}
func main() {
	var myMap Map
	myMap.Add("apple", "banana", "cherry")
	// 输出: ["apple" "banana" "cherry"]

    fruits := []string{"apple", "banana", "cherry"}
    myMap.Add(fruits...) // 需要使用 "..." 展开切片 不展开就报错

}
```
**Context**
`ctx, cancel := context.WithCancel(context.Background())`
context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel。

