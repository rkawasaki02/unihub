package main

import (
	"fmt"
	"net/http"
	"unihub/backend/api"
)

func main() {
	// ルーティング設定
	http.HandleFunc("/api/files", api.ListFiles)

	// サーバー起動
	port := ":8080"
	fmt.Println("unihubサーバー起動中... port" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("サーバー起動失敗:", err)
	}
}

