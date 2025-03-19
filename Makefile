.PHONY: build install clean help formal psychedelic cosmic digital death run all

# デフォルトターゲット
all: build

# ビルドコマンド
build:
	go build -o bunnyspeak

# インストールコマンド
install: build
	mkdir -p $(HOME)/bin
	cp bunnyspeak $(HOME)/bin/
	@echo "bunnyspeak を $(HOME)/bin にインストールしました"
	@echo "PATH環境変数に $(HOME)/bin が含まれていることを確認してください"

# クリーンアップコマンド
clean:
	rm -f bunnyspeak

# 各モード用のコマンド
formal:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		./bunnyspeak convert "$$text" --mode=formal; \
	else \
		./bunnyspeak convert "$(TEXT)" --mode=formal; \
	fi

psychedelic:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		./bunnyspeak convert "$$text" --mode=psychedelic; \
	else \
		./bunnyspeak convert "$(TEXT)" --mode=psychedelic; \
	fi

cosmic:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		./bunnyspeak convert "$$text" --mode=cosmic_absurd; \
	else \
		./bunnyspeak convert "$(TEXT)" --mode=cosmic_absurd; \
	fi

digital:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		./bunnyspeak convert "$$text" --mode=digital_decay; \
	else \
		./bunnyspeak convert "$(TEXT)" --mode=digital_decay; \
	fi

death:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		./bunnyspeak convert "$$text" --mode=death_circular; \
	else \
		./bunnyspeak convert "$(TEXT)" --mode=death_circular; \
	fi

# 指定したファイルの内容を変換
file-convert:
	@if [ -z "$(FILE)" ]; then \
		echo "ファイルパスを指定してください: make file-convert FILE=path/to/file MODE=モード名"; \
	elif [ -z "$(MODE)" ]; then \
		./bunnyspeak convert "$$(cat $(FILE))" --mode=formal; \
	else \
		./bunnyspeak convert "$$(cat $(FILE))" --mode=$(MODE); \
	fi

# Dockerイメージのビルドとコマンド実行
docker-build:
	docker build -t bunnyspeak .

docker-run:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		docker run bunnyspeak convert "$$text" $(if $(MODE),--mode=$(MODE),); \
	else \
		docker run bunnyspeak convert "$(TEXT)" $(if $(MODE),--mode=$(MODE),); \
	fi

# すべてのモードでテキストを変換
all-modes:
	@if [ -z "$(TEXT)" ]; then \
		read -p "テキストを入力してください: " text; \
		echo "【フォーマルモード】"; \
		./bunnyspeak convert "$$text" --mode=formal; \
		echo "\n【サイケデリックモード】"; \
		./bunnyspeak convert "$$text" --mode=psychedelic; \
		echo "\n【コズミック・アブサードモード】"; \
		./bunnyspeak convert "$$text" --mode=cosmic_absurd; \
		echo "\n【デジタル崩壊モード】"; \
		./bunnyspeak convert "$$text" --mode=digital_decay; \
		echo "\n【デスサーキュラーモード】"; \
		./bunnyspeak convert "$$text" --mode=death_circular; \
	else \
		echo "【フォーマルモード】"; \
		./bunnyspeak convert "$(TEXT)" --mode=formal; \
		echo "\n【サイケデリックモード】"; \
		./bunnyspeak convert "$(TEXT)" --mode=psychedelic; \
		echo "\n【コズミック・アブサードモード】"; \
		./bunnyspeak convert "$(TEXT)" --mode=cosmic_absurd; \
		echo "\n【デジタル崩壊モード】"; \
		./bunnyspeak convert "$(TEXT)" --mode=digital_decay; \
		echo "\n【デスサーキュラーモード】"; \
		./bunnyspeak convert "$(TEXT)" --mode=death_circular; \
	fi

# ヘルプコマンド
help:
	@echo "使用可能なコマンド:"
	@echo "  make build             - BunnySpeakをビルドします"
	@echo "  make install           - BunnySpeakをビルドしてインストールします"
	@echo "  make clean             - ビルドファイルを削除します"
	@echo "  make formal [TEXT=文字列]  - フォーマルバニーモードでテキストを変換します"
	@echo "  make psychedelic [TEXT=文字列] - サイケデリックバニーモードでテキストを変換します"
	@echo "  make cosmic [TEXT=文字列]  - コズミック・アブサードモードでテキストを変換します"
	@echo "  make digital [TEXT=文字列] - デジタル崩壊モードでテキストを変換します"
	@echo "  make death [TEXT=文字列]   - デスサーキュラーモードでテキストを変換します"
	@echo "  make file-convert FILE=path/to/file [MODE=モード名] - ファイルの内容を変換します"
	@echo "  make all-modes [TEXT=文字列] - すべてのモードでテキストを変換します"
	@echo "  make docker-build      - Dockerイメージをビルドします"
	@echo "  make docker-run [TEXT=文字列] [MODE=モード名] - Dockerでコマンドを実行します"
	@echo "  make help              - このヘルプを表示します"
	@echo ""
	@echo "使用例:"
	@echo "  make formal TEXT=\"こんにちは世界\""
	@echo "  make digital TEXT=\"システムエラー発生\""
	@echo "  make file-convert FILE=input.txt MODE=death"
	@echo "  make all-modes TEXT=\"全モードでテスト\"" 
	@echo "  cat input.txt | xargs -I {} make formal TEXT=\"{}\""
