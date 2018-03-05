package wareki

import "testing"

func TestSeirekiToWareki(t *testing.T) {
	s := seireki{2018, 3, 31}
	wareki, err := s.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", s.year, s.month, s.day)
	}

	if wareki.name != "平成" || wareki.yomi != "へいせい" || wareki.year != 30 || wareki.month != 3 || wareki.day != 31 {
		t.Errorf("expected\t平成（へいせい）30年3月31日\ngot\t%s（%s）%d年%d月%d日", wareki.name, wareki.yomi, wareki.year, wareki.month, wareki.day)
	}

	s = seireki{1989, 1, 7}
	wareki, err = s.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", s.year, s.month, s.day)
	}

	if wareki.name != "昭和" || wareki.yomi != "しょうわ" || wareki.year != 64 || wareki.month != 1 || wareki.day != 7 {
		t.Errorf("expected\t昭和（しょうわ）1年1月7日\ngot\t%s（%s）%d年%d月%d日", wareki.name, wareki.yomi, wareki.year, wareki.month, wareki.day)
	}

	s = seireki{1989, 1, 8}
	wareki, err = s.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", s.year, s.month, s.day)
	}

	if wareki.name != "平成" || wareki.yomi != "へいせい" || wareki.year != 1 || wareki.month != 1 || wareki.day != 8 {
		t.Errorf("expected\t平成（へいせい）1年1月8日\ngot\t%s（%s）%d年%d月%d日", wareki.name, wareki.yomi, wareki.year, wareki.month, wareki.day)
	}

	s = seireki{123, 1, 8}
	wareki, err = s.Wareki()
	if err != nil {
		t.Errorf("%d/%d/%d should be parsed successfully", s.year, s.month, s.day)
	}

	if wareki.name != "【元号不明】" || wareki.yomi != "" || wareki.year != 123 || wareki.month != 1 || wareki.day != 8 {
		t.Errorf("expected\t平成（へいせい）1年1月8日\ngot\t%s（%s）%d年%d月%d日", wareki.name, wareki.yomi, wareki.year, wareki.month, wareki.day)
	}

}

func TestCompareDates(t *testing.T) {
	date1 := seireki{2018, 1, 13}
	date2 := seireki{2017, 4, 8}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != 1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 1", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}

	date1 = seireki{2018, 1, 13}
	date2 = seireki{2018, 1, 14}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != -1 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return -1", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}

	date1 = seireki{2018, 1, 13}
	date2 = seireki{2018, 1, 13}

	if CompareDate(date1.year, date1.month, date1.day, date2.year, date2.month, date2.day) != 0 {
		t.Errorf("comparison of %d/%d/%d with %d/%d/%d should return 0", date1.year, date1.month, date1.day, date2.year, date2.month, date2.day)
	}
}
