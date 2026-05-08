package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// NASのルートパス
const NASRoot = "/mnt/d/nas"

// ファイル情報の構造体
type FileInfo struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	IsDir   bool      `json:"isDir"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
}

// ファイル一覧を返すハンドラー
func ListFiles(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからパスを取得
	path := r.URL.Query().Get("path")
	if path == "" {
		path = NASRoot
	} else {
		path = filepath.Join(NASRoot, path)
	}

	// ディレクトリを読み込む
	entries, err := os.ReadDir(path)
	if err != nil {
		http.Error(w, "ディレクトリが読み込めません", http.StatusInternalServerError)
		return
	}

	// ファイル情報を収集
	var files []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, FileInfo{
			Name:    entry.Name(),
			Path:    filepath.Join(path, entry.Name()),
			IsDir:   entry.IsDir(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
		})
	}

	// JSONで返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

