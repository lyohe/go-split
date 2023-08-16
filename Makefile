all: split.out

# 実行可能ファイルの生成
split.out: config.go convert.go main.go split.go suffix.go
	go build -o split.out

# クリーンアップ
clean:
	rm -f split.out
	rm -f split.txtar

# x で始まり拡張子を持たない出力ファイルの削除
remove:
	find . -maxdepth 1 -type f -name 'x*' ! -name '*.*' -exec rm {} \;

# txtar で出力
txtar:
	find . -name .git -prune -o -type f -not -name "split.txtar" -print | xargs go run golang.org/x/exp/cmd/txtar@latest > split.txtar

.PHONY: all clean remove txtar
