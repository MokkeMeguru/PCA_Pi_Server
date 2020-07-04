package sound

import (
	"os"
	"reflect"
	"testing"
)

func TestFindSounds(t *testing.T) {
	tests := []struct {
		name string
		saveFolder string
		want []string
	}{
		{name: "case1",
		saveFolder: "../../assets/",
		want:[]string{"test.png"},
		},
		{name: "case2",
		saveFolder: "../../assets/",
		want:[]string{"サウンド_1.wav"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.OpenFile(tt.saveFolder + tt.want[0], 
				os.O_WRONLY|os.O_CREATE, 0666)

			if got := FindSounds(tt.saveFolder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSounds() = %v, want %v", got, tt.want)
			}
			os.RemoveAll(tt.saveFolder + tt.want[0])
		})
	}
}
