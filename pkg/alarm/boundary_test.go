package alarm

import (
	"bufio"
	"os"
	"testing"
)

func TestConfigWrite(t *testing.T) {
	type args struct {
		confStrList []string
		path        string
	}
	tests := []struct {
		name string
		args args
	}{
		{"testcase1", args{[]string{"hello", "nice", "to", "meet", "you"}, "../../configs/test-alarm.conf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigWrite(tt.args.confStrList, tt.args.path)
			file, err := os.OpenFile(tt.args.path, os.O_WRONLY, 0666)
			if err != nil {
				t.Fatalf("file not saved")
			}
			scanner := bufio.NewScanner(file)
			idx := 0
			for scanner.Scan() {
				if tt.args.confStrList[idx] != scanner.Text() {
					t.Fatalf("saved file is invalid")
				}
			}
			_ = os.Remove(tt.args.path)
		})
	}
}
