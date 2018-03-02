package wareki

import "testing"

func TestSeirekiToWareki(t *testing.T) {
	var seireki Seireki

	seireki.Year = 2018
	seireki.Month = 3
	seireki.Day = 31

	wareki, err := seireki.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", seireki.Year, seireki.Month, seireki.Day)
	}

	if wareki.Name != "平成" || wareki.Yomi != "へいせい" || wareki.Year != 30 || wareki.Month != 3 || wareki.Day != 31 {
		t.Errorf("expected\t平成（へいせい）30年3月31日\ngot\t%s（%s）%d年%d月%d日", wareki.Name, wareki.Yomi, wareki.Year, wareki.Month, wareki.Day)
	}

	seireki.Year = 1989
	seireki.Month = 1
	seireki.Day = 7

	wareki, err = seireki.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", seireki.Year, seireki.Month, seireki.Day)
	}

	if wareki.Name != "昭和" || wareki.Yomi != "しょうわ" || wareki.Year != 64 || wareki.Month != 1 || wareki.Day != 7 {
		t.Errorf("expected\t昭和（しょうわ）1年1月7日\ngot\t%s（%s）%d年%d月%d日", wareki.Name, wareki.Yomi, wareki.Year, wareki.Month, wareki.Day)
	}

	seireki.Year = 1989
	seireki.Month = 1
	seireki.Day = 8

	wareki, err = seireki.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", seireki.Year, seireki.Month, seireki.Day)
	}

	if wareki.Name != "平成" || wareki.Yomi != "へいせい" || wareki.Year != 1 || wareki.Month != 1 || wareki.Day != 8 {
		t.Errorf("expected\t平成（へいせい）1年1月8日\ngot\t%s（%s）%d年%d月%d日", wareki.Name, wareki.Yomi, wareki.Year, wareki.Month, wareki.Day)
	}

	seireki.Year = 123
	seireki.Month = 1
	seireki.Day = 8

	wareki, err = seireki.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", seireki.Year, seireki.Month, seireki.Day)
	}

	if wareki.Name != "【元号不明】" || wareki.Yomi != "" || wareki.Year != 123 || wareki.Month != 1 || wareki.Day != 8 {
		t.Errorf("expected\t平成（へいせい）1年1月8日\ngot\t%s（%s）%d年%d月%d日", wareki.Name, wareki.Yomi, wareki.Year, wareki.Month, wareki.Day)
	}

}

func TestCompareDates(t *testing.T) {
	var date1 Seireki
	var date2 Seireki

	date1.Year = 2018
	date1.Month = 1
	date1.Day = 13

	date2.Year = 2017
	date2.Month = 4
	date2.Day = 8

	if CompareDate(date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day) != 1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 1", date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day)
	}

	date1.Year = 2018
	date1.Month = 1
	date1.Day = 13

	date2.Year = 2018
	date2.Month = 1
	date2.Day = 14

	if CompareDate(date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day) != -1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return -1", date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day)
	}

	date1.Year = 2018
	date1.Month = 1
	date1.Day = 13

	date2.Year = 2018
	date2.Month = 1
	date2.Day = 13

	if CompareDate(date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day) != 0 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 0", date1.Year, date1.Month, date1.Day, date2.Year, date2.Month, date2.Day)
	}
}
