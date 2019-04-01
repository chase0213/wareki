package wareki

import (
	"reflect"
	"testing"
)

func TestSeireki_Wareki(t *testing.T) {
	type fields struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Wareki
		wantErr bool
	}{
		{
			name: "it should return 平成 date",
			fields: fields{
				year:  2018,
				month: 3,
				day:   31,
			},
			want: &Wareki{
				name:  "平成",
				yomi:  "へいせい",
				year:  30,
				month: 3,
				day:   31,
			},
			wantErr: false,
		},
		{
			name: "it should return 昭和 date",
			fields: fields{
				year:  1989,
				month: 1,
				day:   7,
			},
			want: &Wareki{
				name:  "昭和",
				yomi:  "しょうわ",
				year:  64,
				month: 1,
				day:   7,
			},
			wantErr: false,
		},
		{
			name: "it should return 平成 date",
			fields: fields{
				year:  1989,
				month: 1,
				day:   8,
			},
			want: &Wareki{
				name:  "平成",
				yomi:  "へいせい",
				year:  1,
				month: 1,
				day:   8,
			},
			wantErr: false,
		},
		{
			name: "it should return 平成 date",
			fields: fields{
				year:  2019,
				month: 4,
				day:   30,
			},
			want: &Wareki{
				name:  "平成",
				yomi:  "へいせい",
				year:  31,
				month: 4,
				day:   30,
			},
			wantErr: false,
		},
		{
			name: "it should return 令和 date",
			fields: fields{
				year:  2019,
				month: 5,
				day:   1,
			},
			want: &Wareki{
				name:  "令和",
				yomi:  "れいわ",
				year:  1,
				month: 5,
				day:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Seireki{
				year:  tt.fields.year,
				month: tt.fields.month,
				day:   tt.fields.day,
			}
			got, err := s.Wareki()
			if (err != nil) != tt.wantErr {
				t.Errorf("Seireki.Wareki() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seireki.Wareki() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCompareDates(t *testing.T) {
	date1 := Seireki{2018, 1, 13}
	date2 := Seireki{2017, 4, 8}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != 1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 1", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}

	date1 = Seireki{2018, 1, 13}
	date2 = Seireki{2018, 1, 14}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != -1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return -1", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}

	date1 = Seireki{2018, 1, 13}
	date2 = Seireki{2018, 1, 13}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != 0 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 0", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}
}
