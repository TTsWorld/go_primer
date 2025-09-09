package main

import (
	"context"
	"os"
)

func getNovel(ctx context.Context) (*NovelRequest, error) {
	req := NovelRequest{
		ResourceId: 10000,
		AnswerId:   10000,
		Content:    getNovelTex(),
		Question:   "有什么好看的穿越小说推荐？",
	}
	return &req, nil
}

func getNovelTex() string {
	//阅读当前目录下的 novel.txt 文件中的内容
	content, err := os.ReadFile("novel.txt")
	if err != nil {
		return ""
	}
	return string(content)

}
