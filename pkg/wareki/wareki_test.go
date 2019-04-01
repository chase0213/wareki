package wareki

import (
	"reflect"
	"testing"
)

func TestWareki_Seireki(t *testing.T) {
	type fields struct {
		name  string
		yomi  string
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Seireki
		wantErr bool
	}{
		{
			name: "it should return A.D. date",
			fields: fields{
				name:  "平成",
				yomi:  "へいせい",
				year:  30,
				month: 2,
				day:   13,
			},
			want: &Seireki{
				year:  2018,
				month: 2,
				day:   13,
			},
			wantErr: false,
		},
		{
			name: "it should return A.D. date",
			fields: fields{
				name:  "昭和",
				yomi:  "しょうわ",
				year:  64,
				month: 1,
				day:   7,
			},
			want: &Seireki{
				year:  1989,
				month: 1,
				day:   7,
			},
			wantErr: false,
		},
		{
			name: "it should return A.D. date",
			fields: fields{
				name:  "令和",
				yomi:  "れいわ",
				year:  1,
				month: 5,
				day:   1,
			},
			want: &Seireki{
				year:  2019,
				month: 5,
				day:   1,
			},
			wantErr: false,
		},
		{
			name: "it should return A.D. date",
			fields: fields{
				name:  "存在しない元号",
				yomi:  "",
				year:  123,
				month: 2,
				day:   13,
			},
			want: &Seireki{
				year:  122,
				month: 2,
				day:   13,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wareki{
				name:  tt.fields.name,
				yomi:  tt.fields.yomi,
				year:  tt.fields.year,
				month: tt.fields.month,
				day:   tt.fields.day,
			}
			got, err := w.Seireki()
			if (err != nil) != tt.wantErr {
				t.Errorf("Wareki.Seireki() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wareki.Seireki() = %v, want %v", got, tt.want)
			}
		})
	}
}
