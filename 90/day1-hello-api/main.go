package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response 是统一返回结构（Day 5 会系统讲）
// 但从 Day 1 就养成好习惯
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// helloHandler 是一个 HTTP Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 GET 请求
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: 6666,
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回 JSON
	_ = json.NewEncoder(w).Encode(resp)
	// 等价于 data = json.Marshal(resp)   w.Write(data)

}

func main() {
	// 1️⃣ 创建路由器
	mux := http.NewServeMux()

	// 2️⃣ 注册路由
	mux.HandleFunc("/hello", helloHandler)

	// 3️⃣ 启动 HTTP 服务
	addr := ":3000"
	log.Println("🚀 Server started at http://localhost" + addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
