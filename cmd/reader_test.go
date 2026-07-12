package cmd

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	// テスト用のダミーJSONをただの文字列として用意する
	expected := `[
		{
			"Date": "2026-05-01",
			"Count": 1,
			"Contents": [{"Detail": "CLIツールを作った"}]
		}
	]`

	r := strings.NewReader(expected)

	records, err := Read(r)
	if err != nil {
		t.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("レコード数は1件であるべきですが、%d件でした", len(records))
	}

	if records[0].Date != "2026-05-01" {
		t.Errorf("日付が正しく読み込めていません: %s", records[0].Date)
	}
}
