all: split.out

# 実行可能ファイルの生成
split.out: main.go split.go suffix.go usage.go convert.go
	go build -o split.out

# クリーンアップ
clean:
	rm -f split.out

# x で始まり拡張子を持たない出力ファイルの削除
remove:
	find . -maxdepth 1 -type f -name 'x*' ! -name '*.*' -exec rm {} \;

.PHONY: all clean remove
