package main

import (
	"testing"
)

func TestHelloWorlf(t *testing.T) {
	t.Run("retunrs hello world", func(t *testing.T) {
		// テストを必ず失敗することを確認する
		// 下記でテストが失敗すらしなかったら、テスト環境がおかしい
		t.Fail()
	})

}
