package wareki

import "testing"

func TestWarekiToSeireki(t *testing.T) {
	w := wareki{"平成", "", 30, 2, 13}
	seireki, err := w.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", w.name, w.year, w.month, w.day)
	}

	if seireki.year != 2018 || seireki.month != 2 || seireki.day != 13 {
		t.Errorf("expected\t2018/2/13\ngot\t%d/%d/%d", seireki.year, seireki.month, seireki.day)
	}

	w = wareki{"昭和", "", 64, 1, 7}
	seireki, err = w.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", w.name, w.year, w.month, w.day)
	}

	if seireki.year != 1989 || seireki.month != 1 || seireki.day != 7 {
		t.Errorf("expected\t1989/1/7\ngot\t%d/%d/%d", seireki.year, seireki.month, seireki.day)
	}

	w = wareki{"存在しない元号", "", 123, 2, 13}
	seireki, err = w.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", w.name, w.year, w.month, w.day)
	}

	if seireki.year != 122 || seireki.month != 2 || seireki.day != 13 {
		t.Errorf("expected\122/2/13\ngot\t%d/%d/%d", seireki.year, seireki.month, seireki.day)
	}
}
