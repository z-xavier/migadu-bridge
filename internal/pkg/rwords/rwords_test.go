package rwords

import "testing"

func TestGetRWordsFromEmbed(t *testing.T) {
	type args struct {
		capitalize    bool
		includeNumber bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "embed",
			args: args{
				capitalize:    false,
				includeNumber: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRWords(tt.args.capitalize, tt.args.includeNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GetRWords() got = %v", got)
		})
	}
}
