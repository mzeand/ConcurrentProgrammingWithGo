package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var filelist []string

func fileSearch(dir string, filename string, wg *sync.WaitGroup) {
	// 関数に渡されたディレクトリ内の全てのファイルを取得
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		// 各ファイルをディレクトリに結合してフルパスを作成
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			// 一致すればコンソールにパスを出力
			filelist = append(filelist, fpath)
		}
		if file.IsDir() {
			// ディレクトリの場合、新たなゴルーチンを開始する前にWaitGroupにカウントを追加
			wg.Add(1)
			// 再帰的にfileSearchを呼び出す
			go fileSearch(fpath, filename, wg)
		}
	}
	// 現在のゴルーチンの作業が完了したことを通知
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	filelist = make([]string, 0)
	go fileSearch(os.Args[1], os.Args[2], &wg)
	wg.Wait()
	sort.Strings(filelist)
	for _, file := range filelist {
		fmt.Println(file)
	}
}
