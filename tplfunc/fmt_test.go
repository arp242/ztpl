package tplfunc

import (
	"testing"
	"time"
)

func TestFmt(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		t.Fatal(err)
	}
	ti := time.Date(2021, 4, 5, 14, 49, 20, 666, tz)
	var tip *time.Time

	tests := []struct {
		tpl  string
		want string
		data any
	}{
		{`{{json (map "k" 5 "k2" "q")}}`, `{"k":5,"k2":"q"}`, nil},
		{`{{json_pretty (map "k" 5 "k2" "q")}}`, "{\n    \"k\": 5,\n    \"k2\": \"q\"\n}", nil},

		{`{{number 999 0x2009}}`, "999", nil},
		{`{{number 1000 0x2009}}`, "1 000", nil},
		{`{{number 20000 ','}}`, "20,000", nil},
		{`{{number 300000 '.'}}`, "300.000", nil},
		{`{{number 4987654 '\''}}`, "4'987'654", nil},
		{`{{number 4987654 0x00}}`, "4987654", nil},
		{`{{number 4987654 0x01}}`, "4987654", nil},
		{`{{number 1.1 ','}}`, "1.1", nil},
		{`{{number 9999.11111 ','}}`, "9,999.11111", nil},
		{`{{number 9999.11111 '.'}}`, "9.999,11111", nil},

		{`{{large_number 1000 ','}}`, "1,000", nil},
		{`{{large_number 10_001 ','}}`, "10k", nil},
		{`{{large_number 94_301.123 ','}}`, "94k", nil},
		{`{{large_number 1_090_000 ','}}`, "1.1M", nil},
		{`{{large_number 1_090_000 '.'}}`, "1,1M", nil},
		{`{{large_number 101_090_000 ','}}`, "101M", nil},

		{`{{time . ""}}`, "2021-04-05", ti},
		{`{{time . "2006"}}`, "2021", ti},
		{`{{time . "ansic"}}`, "Mon Apr  5 14:49:20 2021", ti},
		{`{{time . ""}}`, "", time.Time{}},
		{`{{time . ""}}`, "", &time.Time{}},
		{`{{time . ""}}`, "", tip},

		{`{{size 14}}`, "14.0", nil},
		{`{{size 1444}}`, "1.4K", nil},
		{`{{size 1444 "from=k"}}`, "1.4M", nil},

		{`{{slug "hello, world"}}`, "hello-world", nil},
		{`{{slug "hello, world!"}}`, "hello-world", nil},
		{`{{slug "hëllø, wörłd"}}`, "hëllø-wörłd", nil},
		{`{{"አማርኛ" | slug}}`, "አማርኛ", nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have := mktpl(t, "", tt.tpl, tt.data)
			if have != tt.want {
				t.Errorf("\nhave: %s\nwant: %s\n", have, tt.want)
			}
		})
	}
}
