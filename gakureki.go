package main

type GakurekiOption struct {
	ElementarySchool int
	JuniorHighSchool int
	HighSchool       int
	University       int
}

type Gakureki []GakurekiRecord

type GakurekiRecord struct {
	BeginYear  int
	BeginMonth int
	BeginDay   int
	EndYear    int
	EndMonth   int
	EndDay     int
	Active     bool
}

func (s *Seireki) Gakureki(options GakurekiOption) {

}
