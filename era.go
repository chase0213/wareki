package wareki

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/dghubble/trie"
)

type EraSearcher struct {
	Eras []Era
	Trie trie.Trier
}

type Era struct {
	Name       string `json:"name"`
	Yomi       string `json:"yomi"`
	BeginYear  int    `json:"begin_year"`
	BeginMonth int    `json:"begin_month"`
	BeginDay   int    `json:"begin_day"`
	EndYear    int    `json:"end_year"`
	EndMonth   int    `json:"end_month"`
	EndDay     int    `json:"end_day"`
}

type rawEraData struct {
	Name    string
	Yomi    string
	BeginAt string
	EndAt   string
}

func (es *EraSearcher) Search(query string) (Era, error) {
	var era Era
	jsonEra := fmt.Sprintf("%s", es.Trie.Get(query))
	if err := json.Unmarshal([]byte(jsonEra), &era); err != nil {
		return Era{}, err
	}

	return era, nil
}

func NewEraSearcher() (*EraSearcher, error) {
	eras, err := LoadEras()
	if err != nil {
		return nil, err
	}

	trie, err := constructTrieFromEras(eras)

	es := EraSearcher{
		Eras: eras,
		Trie: trie,
	}

	return &es, nil
}

func LoadEras() ([]Era, error) {
	rawEras := []rawEraData{
		rawEraData{
			Name:    "白雉",
			Yomi:    "はくち",
			BeginAt: "650-3-22",
			EndAt:   "654-11-24",
		},
		rawEraData{
			Name:    "－",
			Yomi:    "－",
			BeginAt: "654-11-24",
			EndAt:   "686-8-14",
		},
		rawEraData{
			Name:    "朱鳥",
			Yomi:    "しゅちょう",
			BeginAt: "686-8-14",
			EndAt:   "686-10-1",
		},
		rawEraData{
			Name:    "－",
			Yomi:    "－",
			BeginAt: "686-10-1",
			EndAt:   "701-5-3",
		},
		rawEraData{
			Name:    "大宝",
			Yomi:    "たいほう",
			BeginAt: "701-5-3",
			EndAt:   "704-6-16",
		},
		rawEraData{
			Name:    "慶雲",
			Yomi:    "けいうん",
			BeginAt: "704-6-16",
			EndAt:   "708-2-7",
		},
		rawEraData{
			Name:    "和銅",
			Yomi:    "わどう",
			BeginAt: "708-2-7",
			EndAt:   "715-10-3",
		},
		rawEraData{
			Name:    "養老",
			Yomi:    "ようろう",
			BeginAt: "717-12-24",
			EndAt:   "724-3-3",
		},
		rawEraData{
			Name:    "神亀",
			Yomi:    "じんき",
			BeginAt: "724-3-3",
			EndAt:   "729-9-2",
		},
		rawEraData{
			Name:    "天平",
			Yomi:    "てんぴょう",
			BeginAt: "729-9-2",
			EndAt:   "749-5-4",
		},
		rawEraData{
			Name:    "天平感宝",
			Yomi:    "てんぴょうかんぽう",
			BeginAt: "749-5-4",
			EndAt:   "749-8-19",
		},
		rawEraData{
			Name:    "天平勝宝",
			Yomi:    "てんぴょうしょうほう",
			BeginAt: "749-8-19",
			EndAt:   "757-9-6",
		},
		rawEraData{
			Name:    "天平宝字",
			Yomi:    "てんぴょうほうじ",
			BeginAt: "757-9-6",
			EndAt:   "765-2-1",
		},
		rawEraData{
			Name:    "天平神護",
			Yomi:    "てんぴょうじんご",
			BeginAt: "765-2-1",
			EndAt:   "767-9-13",
		},
		rawEraData{
			Name:    "神護景雲",
			Yomi:    "じんごけいうん",
			BeginAt: "767-9-13",
			EndAt:   "770-10-23",
		},
		rawEraData{
			Name:    "宝亀",
			Yomi:    "ほうき",
			BeginAt: "770-10-23",
			EndAt:   "781-1-30",
		},
		rawEraData{
			Name:    "天応",
			Yomi:    "てんおう",
			BeginAt: "781-1-30",
			EndAt:   "782-9-30",
		},
		rawEraData{
			Name:    "延暦",
			Yomi:    "えんりゃく",
			BeginAt: "782-9-30",
			EndAt:   "806-6-8",
		},
		rawEraData{
			Name:    "弘仁",
			Yomi:    "こうにん",
			BeginAt: "810-10-20",
			EndAt:   "824-2-8",
		},
		rawEraData{
			Name:    "天長",
			Yomi:    "てんちょう",
			BeginAt: "824-2-8",
			EndAt:   "834-2-14",
		},
		rawEraData{
			Name:    "承和",
			Yomi:    "じょうわ",
			BeginAt: "834-2-14",
			EndAt:   "848-7-16",
		},
		rawEraData{
			Name:    "嘉祥",
			Yomi:    "かしょう",
			BeginAt: "848-7-16",
			EndAt:   "851-6-1",
		},
		rawEraData{
			Name:    "仁寿",
			Yomi:    "にんじゅ",
			BeginAt: "851-6-1",
			EndAt:   "854-12-23",
		},
		rawEraData{
			Name:    "斉衡",
			Yomi:    "さいこう",
			BeginAt: "854-12-23",
			EndAt:   "857-3-20",
		},
		rawEraData{
			Name:    "天安",
			Yomi:    "てんあん",
			BeginAt: "857-3-20",
			EndAt:   "859-5-20",
		},
		rawEraData{
			Name:    "貞観",
			Yomi:    "じょうがん",
			BeginAt: "859-5-20",
			EndAt:   "877-6-1",
		},
		rawEraData{
			Name:    "元慶",
			Yomi:    "がんぎょう",
			BeginAt: "877-6-1",
			EndAt:   "885-3-11",
		},
		rawEraData{
			Name:    "仁和",
			Yomi:    "にんな",
			BeginAt: "885-3-11",
			EndAt:   "889-5-30",
		},
		rawEraData{
			Name:    "寛平",
			Yomi:    "かんぴょう",
			BeginAt: "889-5-30",
			EndAt:   "898-5-2",
		},
		rawEraData{
			Name:    "昌泰",
			Yomi:    "しょうたい",
			BeginAt: "898-5-20",
			EndAt:   "901-8-31",
		},
		rawEraData{
			Name:    "延喜",
			Yomi:    "えんぎ",
			BeginAt: "901-8-31",
			EndAt:   "923-5-29",
		},
		rawEraData{
			Name:    "延長",
			Yomi:    "えんちょう",
			BeginAt: "923-5-29",
			EndAt:   "931-5-16",
		},
		rawEraData{
			Name:    "承平",
			Yomi:    "じょうへい",
			BeginAt: "931-5-16",
			EndAt:   "938-6-22",
		},
		rawEraData{
			Name:    "天慶",
			Yomi:    "てんぎょう",
			BeginAt: "938-6-22",
			EndAt:   "947-5-15",
		},
		rawEraData{
			Name:    "天暦",
			Yomi:    "てんりゃく",
			BeginAt: "947-5-15",
			EndAt:   "957-11-21",
		},
		rawEraData{
			Name:    "天徳",
			Yomi:    "てんとく",
			BeginAt: "957-11-21",
			EndAt:   "961-3-5",
		},
		rawEraData{
			Name:    "応和",
			Yomi:    "おうわ",
			BeginAt: "961-3-5",
			EndAt:   "964-8-19",
		},
		rawEraData{
			Name:    "康保",
			Yomi:    "こうほう",
			BeginAt: "964-8-19",
			EndAt:   "968-9-8",
		},
		rawEraData{
			Name:    "安和",
			Yomi:    "あんな",
			BeginAt: "968-9-8",
			EndAt:   "970-5-3",
		},
		rawEraData{
			Name:    "天禄",
			Yomi:    "てんろく",
			BeginAt: "970-5-3",
			EndAt:   "974-1-16",
		},
		rawEraData{
			Name:    "天延",
			Yomi:    "てんえん",
			BeginAt: "974-1-16",
			EndAt:   "976-8-11",
		},
		rawEraData{
			Name:    "貞元",
			Yomi:    "じょうげん",
			BeginAt: "976-8-11",
			EndAt:   "978-12-31",
		},
		rawEraData{
			Name:    "天元",
			Yomi:    "てんげん",
			BeginAt: "978-12-31",
			EndAt:   "983-5-29",
		},
		rawEraData{
			Name:    "永観",
			Yomi:    "えいかん",
			BeginAt: "983-5-29",
			EndAt:   "985-5-19",
		},
		rawEraData{
			Name:    "寛和",
			Yomi:    "かんな",
			BeginAt: "985-5-19",
			EndAt:   "987-5-5",
		},
		rawEraData{
			Name:    "永延",
			Yomi:    "えいえん",
			BeginAt: "987-5-5",
			EndAt:   "989-9-10",
		},
		rawEraData{
			Name:    "永祚",
			Yomi:    "えいそ",
			BeginAt: "989-9-10",
			EndAt:   "990-11-26",
		},
		rawEraData{
			Name:    "正暦",
			Yomi:    "しょうりゃく",
			BeginAt: "990-11-26",
			EndAt:   "995-3-25",
		},
		rawEraData{
			Name:    "長徳",
			Yomi:    "ちょうとく",
			BeginAt: "995-3-25",
			EndAt:   "999-2-1",
		},
		rawEraData{
			Name:    "長保",
			Yomi:    "ちょうほう",
			BeginAt: "999-2-1",
			EndAt:   "1004-8-8",
		},
		rawEraData{
			Name:    "寛弘",
			Yomi:    "かんこう",
			BeginAt: "1004-8-8",
			EndAt:   "1013-2-8",
		},
		rawEraData{
			Name:    "長和",
			Yomi:    "ちょうわ",
			BeginAt: "1013-2-8",
			EndAt:   "1017-5-21",
		},
		rawEraData{
			Name:    "寛仁",
			Yomi:    "かんにん",
			BeginAt: "1017-5-21",
			EndAt:   "1021-3-17",
		},
		rawEraData{
			Name:    "治安",
			Yomi:    "じあん",
			BeginAt: "1021-3-17",
			EndAt:   "1024-8-19",
		},
		rawEraData{
			Name:    "万寿",
			Yomi:    "まんじゅ",
			BeginAt: "1024-8-19",
			EndAt:   "1028-8-18",
		},
		rawEraData{
			Name:    "長元",
			Yomi:    "ちょうげん",
			BeginAt: "1028-8-18",
			EndAt:   "1037-5-9",
		},
		rawEraData{
			Name:    "長暦",
			Yomi:    "ちょうりゃく",
			BeginAt: "1037-5-9",
			EndAt:   "1040-12-16",
		},
		rawEraData{
			Name:    "長久",
			Yomi:    "ちょうきゅう",
			BeginAt: "1040-12-16",
			EndAt:   "1044-12-16",
		},
		rawEraData{
			Name:    "寛徳",
			Yomi:    "かんとく",
			BeginAt: "1044-12-16",
			EndAt:   "1046-5-22",
		},
		rawEraData{
			Name:    "永承",
			Yomi:    "えいしょう",
			BeginAt: "1046-5-22",
			EndAt:   "1053-2-2",
		},
		rawEraData{
			Name:    "天喜",
			Yomi:    "てんき",
			BeginAt: "1053-2-2",
			EndAt:   "1058-9-19",
		},
		rawEraData{
			Name:    "康平",
			Yomi:    "こうへい",
			BeginAt: "1058-9-19",
			EndAt:   "1065-9-4",
		},
		rawEraData{
			Name:    "治暦",
			Yomi:    "じりゃく",
			BeginAt: "1065-9-4",
			EndAt:   "1069-5-6",
		},
		rawEraData{
			Name:    "延久",
			Yomi:    "えんきゅう",
			BeginAt: "1069-5-6",
			EndAt:   "1074-9-16",
		},
		rawEraData{
			Name:    "承保",
			Yomi:    "じょうほう",
			BeginAt: "1074-9-16",
			EndAt:   "1077-12-5",
		},
		rawEraData{
			Name:    "承暦",
			Yomi:    "じょうりゃく",
			BeginAt: "1077-12-5",
			EndAt:   "1081-3-22",
		},
		rawEraData{
			Name:    "永保",
			Yomi:    "えいほう",
			BeginAt: "1081-3-22",
			EndAt:   "1084-3-15",
		},
		rawEraData{
			Name:    "応徳",
			Yomi:    "おうとく",
			BeginAt: "1084-3-15",
			EndAt:   "1087-5-11",
		},
		rawEraData{
			Name:    "寛治",
			Yomi:    "かんじ",
			BeginAt: "1087-5-11",
			EndAt:   "1095-1-23",
		},
		rawEraData{
			Name:    "嘉保",
			Yomi:    "かほう",
			BeginAt: "1095-1-23",
			EndAt:   "1097-1-3",
		},
		rawEraData{
			Name:    "永長",
			Yomi:    "えいちょう",
			BeginAt: "1097-1-3",
			EndAt:   "1097-12-27",
		},
		rawEraData{
			Name:    "承徳",
			Yomi:    "じょうとく",
			BeginAt: "1097-12-27",
			EndAt:   "1099-9-15",
		},
		rawEraData{
			Name:    "康和",
			Yomi:    "こうわ",
			BeginAt: "1099-9-15",
			EndAt:   "1104-3-8",
		},
		rawEraData{
			Name:    "長治",
			Yomi:    "ちょうじ",
			BeginAt: "1104-3-8",
			EndAt:   "1106-5-13",
		},
		rawEraData{
			Name:    "嘉承",
			Yomi:    "かしょう",
			BeginAt: "1106-5-13",
			EndAt:   "1108-9-9",
		},
		rawEraData{
			Name:    "天仁",
			Yomi:    "てんにん",
			BeginAt: "1108-9-9",
			EndAt:   "1110-7-31",
		},
		rawEraData{
			Name:    "天永",
			Yomi:    "てんえい",
			BeginAt: "1110-7-31",
			EndAt:   "1113-8-25",
		},
		rawEraData{
			Name:    "永久",
			Yomi:    "えいきゅう",
			BeginAt: "1113-8-25",
			EndAt:   "1118-4-25",
		},
		rawEraData{
			Name:    "元永",
			Yomi:    "げんえい",
			BeginAt: "1118-4-25",
			EndAt:   "1120-5-9",
		},
		rawEraData{
			Name:    "保安",
			Yomi:    "ほうあん",
			BeginAt: "1120-5-9",
			EndAt:   "1124-5-18",
		},
		rawEraData{
			Name:    "天治",
			Yomi:    "てんじ",
			BeginAt: "1124-5-18",
			EndAt:   "1126-2-15",
		},
		rawEraData{
			Name:    "大治",
			Yomi:    "だいじ",
			BeginAt: "1126-2-15",
			EndAt:   "1131-2-28",
		},
		rawEraData{
			Name:    "天承",
			Yomi:    "てんしょう",
			BeginAt: "1131-2-28",
			EndAt:   "1132-9-21",
		},
		rawEraData{
			Name:    "長承",
			Yomi:    "ちょうしょう",
			BeginAt: "1132-9-21",
			EndAt:   "1135-6-10",
		},
		rawEraData{
			Name:    "保延",
			Yomi:    "ほうえん",
			BeginAt: "1135-6-10",
			EndAt:   "1141-8-13",
		},
		rawEraData{
			Name:    "永治",
			Yomi:    "えいじ",
			BeginAt: "1141-8-13",
			EndAt:   "1142-5-25",
		},
		rawEraData{
			Name:    "康治",
			Yomi:    "こうじ",
			BeginAt: "1142-5-25",
			EndAt:   "1144-3-28",
		},
		rawEraData{
			Name:    "天養",
			Yomi:    "てんよう",
			BeginAt: "1144-3-28",
			EndAt:   "1145-8-12",
		},
		rawEraData{
			Name:    "久安",
			Yomi:    "きゅうあん",
			BeginAt: "1145-8-12",
			EndAt:   "1151-2-14",
		},
		rawEraData{
			Name:    "仁平",
			Yomi:    "にんぺい",
			BeginAt: "1151-2-14",
			EndAt:   "1154-12-4",
		},
		rawEraData{
			Name:    "久寿",
			Yomi:    "きゅうじゅ",
			BeginAt: "1154-12-4",
			EndAt:   "1156-5-18",
		},
		rawEraData{
			Name:    "保元",
			Yomi:    "ほうげん",
			BeginAt: "1156-5-18",
			EndAt:   "1159-5-9",
		},
		rawEraData{
			Name:    "平治",
			Yomi:    "へいじ",
			BeginAt: "1159-5-9",
			EndAt:   "1160-2-18",
		},
		rawEraData{
			Name:    "永暦",
			Yomi:    "えいりゃく",
			BeginAt: "1160-2-18",
			EndAt:   "1161-9-24",
		},
		rawEraData{
			Name:    "応保",
			Yomi:    "おうほう",
			BeginAt: "1161-9-24",
			EndAt:   "1163-5-4",
		},
		rawEraData{
			Name:    "長寛",
			Yomi:    "ちょうかん",
			BeginAt: "1163-5-4",
			EndAt:   "1165-7-14",
		},
		rawEraData{
			Name:    "永万",
			Yomi:    "えいまん",
			BeginAt: "1165-7-14",
			EndAt:   "1166-9-23",
		},
		rawEraData{
			Name:    "仁安",
			Yomi:    "にんあん",
			BeginAt: "1166-9-23",
			EndAt:   "1169-5-6",
		},
		rawEraData{
			Name:    "嘉応",
			Yomi:    "かおう",
			BeginAt: "1169-5-6",
			EndAt:   "1171-5-27",
		},
		rawEraData{
			Name:    "承安",
			Yomi:    "しょうあん",
			BeginAt: "1171-5-27",
			EndAt:   "1175-8-16",
		},
		rawEraData{
			Name:    "安元",
			Yomi:    "あんげん",
			BeginAt: "1175-8-16",
			EndAt:   "1177-8-29",
		},
		rawEraData{
			Name:    "治承",
			Yomi:    "じしょう",
			BeginAt: "1177-8-29",
			EndAt:   "1181-8-25",
		},
		rawEraData{
			Name:    "養和",
			Yomi:    "ようわ",
			BeginAt: "1181-8-25",
			EndAt:   "1182-6-29",
		},
		rawEraData{
			Name:    "寿永",
			Yomi:    "じゅえい",
			BeginAt: "1182-6-29",
			EndAt:   "1184-5-27",
		},
		rawEraData{
			Name:    "元暦",
			Yomi:    "げんりゃく",
			BeginAt: "1184-5-27",
			EndAt:   "1185-9-9",
		},
		rawEraData{
			Name:    "建久",
			Yomi:    "けんきゅう",
			BeginAt: "1190-5-16",
			EndAt:   "1199-5-23",
		},
		rawEraData{
			Name:    "正治",
			Yomi:    "しょうじ",
			BeginAt: "1199-5-23",
			EndAt:   "1201-3-19",
		},
		rawEraData{
			Name:    "建仁",
			Yomi:    "けんにん",
			BeginAt: "1201-3-19",
			EndAt:   "1204-3-23",
		},
		rawEraData{
			Name:    "元久",
			Yomi:    "げんきゅう",
			BeginAt: "1204-3-23",
			EndAt:   "1206-6-5",
		},
		rawEraData{
			Name:    "建永",
			Yomi:    "けんえい",
			BeginAt: "1206-6-5",
			EndAt:   "1207-11-16",
		},
		rawEraData{
			Name:    "承元",
			Yomi:    "じょうげん",
			BeginAt: "1207-11-16",
			EndAt:   "1211-4-23",
		},
		rawEraData{
			Name:    "建暦",
			Yomi:    "けんりゃく",
			BeginAt: "1211-4-23",
			EndAt:   "1214-1-18",
		},
		rawEraData{
			Name:    "建保",
			Yomi:    "けんぽう",
			BeginAt: "1214-1-18",
			EndAt:   "1219-5-27",
		},
		rawEraData{
			Name:    "承久",
			Yomi:    "じょうきゅう",
			BeginAt: "1219-5-27",
			EndAt:   "1222-5-25",
		},
		rawEraData{
			Name:    "貞応",
			Yomi:    "じょうおう",
			BeginAt: "1222-5-25",
			EndAt:   "1224-12-31",
		},
		rawEraData{
			Name:    "元仁",
			Yomi:    "げんにん",
			BeginAt: "1224-12-31",
			EndAt:   "1225-5-28",
		},
		rawEraData{
			Name:    "嘉禄",
			Yomi:    "かろく",
			BeginAt: "1225-5-28",
			EndAt:   "1228-1-18",
		},
		rawEraData{
			Name:    "安貞",
			Yomi:    "あんてい",
			BeginAt: "1228-1-18",
			EndAt:   "1229-3-31",
		},
		rawEraData{
			Name:    "寛喜",
			Yomi:    "かんき",
			BeginAt: "1229-3-31",
			EndAt:   "1232-4-23",
		},
		rawEraData{
			Name:    "貞永",
			Yomi:    "じょうえい",
			BeginAt: "1232-4-23",
			EndAt:   "1233-5-25",
		},
		rawEraData{
			Name:    "天福",
			Yomi:    "てんぷく",
			BeginAt: "1233-5-25",
			EndAt:   "1234-11-27",
		},
		rawEraData{
			Name:    "文暦",
			Yomi:    "ぶんりゃく",
			BeginAt: "1234-11-27",
			EndAt:   "1235-11-1",
		},
		rawEraData{
			Name:    "嘉禎",
			Yomi:    "かてい",
			BeginAt: "1235-11-1",
			EndAt:   "1238-12-30",
		},
		rawEraData{
			Name:    "暦仁",
			Yomi:    "りゃくにん",
			BeginAt: "1238-12-30",
			EndAt:   "1239-3-13",
		},
		rawEraData{
			Name:    "延応",
			Yomi:    "えんおう",
			BeginAt: "1239-3-13",
			EndAt:   "1240-8-5",
		},
		rawEraData{
			Name:    "仁治",
			Yomi:    "にんじ",
			BeginAt: "1240-8-5",
			EndAt:   "1243-3-18",
		},
		rawEraData{
			Name:    "寛元",
			Yomi:    "かんげん",
			BeginAt: "1243-3-18",
			EndAt:   "1247-4-5",
		},
		rawEraData{
			Name:    "宝治",
			Yomi:    "ほうじ",
			BeginAt: "1247-4-5",
			EndAt:   "1249-5-2",
		},
		rawEraData{
			Name:    "建長",
			Yomi:    "けんちょう",
			BeginAt: "1249-5-2",
			EndAt:   "1256-10-24",
		},
		rawEraData{
			Name:    "康元",
			Yomi:    "こうげん",
			BeginAt: "1256-10-24",
			EndAt:   "1257-3-31",
		},
		rawEraData{
			Name:    "正嘉",
			Yomi:    "しょうか",
			BeginAt: "1257-3-31",
			EndAt:   "1259-4-20",
		},
		rawEraData{
			Name:    "正元",
			Yomi:    "しょうげん",
			BeginAt: "1259-4-20",
			EndAt:   "1260-5-24",
		},
		rawEraData{
			Name:    "文応",
			Yomi:    "ぶんおう",
			BeginAt: "1260-5-24",
			EndAt:   "1261-3-22",
		},
		rawEraData{
			Name:    "弘長",
			Yomi:    "こうちょう",
			BeginAt: "1261-3-22",
			EndAt:   "1264-3-27",
		},
		rawEraData{
			Name:    "文永",
			Yomi:    "ぶんえい",
			BeginAt: "1264-3-27",
			EndAt:   "1275-5-22",
		},
		rawEraData{
			Name:    "建治",
			Yomi:    "けんじ",
			BeginAt: "1275-5-22",
			EndAt:   "1278-3-23",
		},
		rawEraData{
			Name:    "弘安",
			Yomi:    "こうあん",
			BeginAt: "1278-3-23",
			EndAt:   "1288-5-29",
		},
		rawEraData{
			Name:    "正応",
			Yomi:    "しょうおう",
			BeginAt: "1288-5-29",
			EndAt:   "1293-9-6",
		},
		rawEraData{
			Name:    "永仁",
			Yomi:    "えいにん",
			BeginAt: "1293-9-6",
			EndAt:   "1299-5-25",
		},
		rawEraData{
			Name:    "正安",
			Yomi:    "しょうあん",
			BeginAt: "1299-5-25",
			EndAt:   "1302-12-10",
		},
		rawEraData{
			Name:    "乾元",
			Yomi:    "けんげん",
			BeginAt: "1302-12-10",
			EndAt:   "1303-9-16",
		},
		rawEraData{
			Name:    "嘉元",
			Yomi:    "かげん",
			BeginAt: "1303-9-16",
			EndAt:   "1307-1-18",
		},
		rawEraData{
			Name:    "徳治",
			Yomi:    "とくじ",
			BeginAt: "1307-1-18",
			EndAt:   "1308-11-22",
		},
		rawEraData{
			Name:    "延慶",
			Yomi:    "えんきょう",
			BeginAt: "1308-11-22",
			EndAt:   "1311-5-17",
		},
		rawEraData{
			Name:    "応長",
			Yomi:    "おうちょう",
			BeginAt: "1311-5-17",
			EndAt:   "1312-4-27",
		},
		rawEraData{
			Name:    "正和",
			Yomi:    "しょうわ",
			BeginAt: "1312-4-27",
			EndAt:   "1317-3-16",
		},
		rawEraData{
			Name:    "文保",
			Yomi:    "ぶんぽう",
			BeginAt: "1317-3-16",
			EndAt:   "1319-5-18",
		},
		rawEraData{
			Name:    "元応",
			Yomi:    "げんおう",
			BeginAt: "1319-5-18",
			EndAt:   "1321-3-22",
		},
		rawEraData{
			Name:    "元亨",
			Yomi:    "げんこう",
			BeginAt: "1321-3-22",
			EndAt:   "1324-12-25",
		},
		rawEraData{
			Name:    "正中",
			Yomi:    "しょうちゅう",
			BeginAt: "1324-12-25",
			EndAt:   "1326-5-28",
		},
		rawEraData{
			Name:    "嘉暦",
			Yomi:    "かりゃく",
			BeginAt: "1326-5-28",
			EndAt:   "1329-9-22",
		},
		rawEraData{
			Name:    "元徳",
			Yomi:    "げんとく",
			BeginAt: "1329-9-22",
			EndAt:   "1332-5-23",
		},
		rawEraData{
			Name:    "興国",
			Yomi:    "こうこく",
			BeginAt: "1340-5-25",
			EndAt:   "1347-1-20",
		},
		rawEraData{
			Name:    "正平",
			Yomi:    "しょうへい",
			BeginAt: "1347-1-20",
			EndAt:   "1370-8-16",
		},
		rawEraData{
			Name:    "建徳",
			Yomi:    "けんとく",
			BeginAt: "1370-8-16",
			EndAt:   "1372-5-1",
		},
		rawEraData{
			Name:    "文中",
			Yomi:    "ぶんちゅう",
			BeginAt: "1372-5-1",
			EndAt:   "1375-6-26",
		},
		rawEraData{
			Name:    "天授",
			Yomi:    "てんじゅ",
			BeginAt: "1375-6-26",
			EndAt:   "1381-3-6",
		},
		rawEraData{
			Name:    "弘和",
			Yomi:    "こうわ",
			BeginAt: "1381-3-6",
			EndAt:   "1384-5-18",
		},
		rawEraData{
			Name:    "元中",
			Yomi:    "げんちゅう",
			BeginAt: "1384-5-18",
			EndAt:   "1392-11-19",
		},
		rawEraData{
			Name:    "暦応",
			Yomi:    "りゃくおう",
			BeginAt: "1338-10-11",
			EndAt:   "1342-6-1",
		},
		rawEraData{
			Name:    "康永",
			Yomi:    "こうえい",
			BeginAt: "1342-6-1",
			EndAt:   "1345-11-15",
		},
		rawEraData{
			Name:    "貞和",
			Yomi:    "じょうわ",
			BeginAt: "1345-11-15",
			EndAt:   "1350-4-4",
		},
		rawEraData{
			Name:    "観応",
			Yomi:    "かんのう",
			BeginAt: "1350-4-4",
			EndAt:   "1352-11-4",
		},
		rawEraData{
			Name:    "文和",
			Yomi:    "ぶんな",
			BeginAt: "1352-11-4",
			EndAt:   "1356-4-29",
		},
		rawEraData{
			Name:    "延文",
			Yomi:    "えんぶん",
			BeginAt: "1356-4-29",
			EndAt:   "1361-5-4",
		},
		rawEraData{
			Name:    "康安",
			Yomi:    "こうあん",
			BeginAt: "1361-5-4",
			EndAt:   "1362-10-11",
		},
		rawEraData{
			Name:    "貞治",
			Yomi:    "じょうじ",
			BeginAt: "1362-10-11",
			EndAt:   "1368-3-7",
		},
		rawEraData{
			Name:    "応安",
			Yomi:    "おうあん",
			BeginAt: "1368-3-7",
			EndAt:   "1375-3-29",
		},
		rawEraData{
			Name:    "永和",
			Yomi:    "えいわ",
			BeginAt: "1375-3-29",
			EndAt:   "1379-4-9",
		},
		rawEraData{
			Name:    "康暦",
			Yomi:    "こうりゃく",
			BeginAt: "1379-4-9",
			EndAt:   "1381-3-20",
		},
		rawEraData{
			Name:    "永徳",
			Yomi:    "えいとく",
			BeginAt: "1381-3-20",
			EndAt:   "1384-3-19",
		},
		rawEraData{
			Name:    "至徳",
			Yomi:    "しとく",
			BeginAt: "1384-3-19",
			EndAt:   "1387-10-5",
		},
		rawEraData{
			Name:    "嘉慶",
			Yomi:    "かきょう",
			BeginAt: "1387-10-5",
			EndAt:   "1389-3-7",
		},
		rawEraData{
			Name:    "康応",
			Yomi:    "こうおう",
			BeginAt: "1389-3-7",
			EndAt:   "1390-4-12",
		},
		rawEraData{
			Name:    "明徳",
			Yomi:    "めいとく",
			BeginAt: "1390-4-12",
			EndAt:   "1394-8-2",
		},
		rawEraData{
			Name:    "正長",
			Yomi:    "しょうちょう",
			BeginAt: "1428-6-10",
			EndAt:   "1429-10-3",
		},
		rawEraData{
			Name:    "永享",
			Yomi:    "えいきょう",
			BeginAt: "1429-10-3",
			EndAt:   "1441-3-10",
		},
		rawEraData{
			Name:    "嘉吉",
			Yomi:    "かきつ",
			BeginAt: "1441-3-10",
			EndAt:   "1444-2-23",
		},
		rawEraData{
			Name:    "文安",
			Yomi:    "ぶんあん",
			BeginAt: "1444-2-23",
			EndAt:   "1449-8-16",
		},
		rawEraData{
			Name:    "宝徳",
			Yomi:    "ほうとく",
			BeginAt: "1449-8-16",
			EndAt:   "1452-8-10",
		},
		rawEraData{
			Name:    "享徳",
			Yomi:    "きょうとく",
			BeginAt: "1452-8-10",
			EndAt:   "1455-9-6",
		},
		rawEraData{
			Name:    "康正",
			Yomi:    "こうしょう",
			BeginAt: "1455-9-6",
			EndAt:   "1457-10-16",
		},
		rawEraData{
			Name:    "長禄",
			Yomi:    "ちょうろく",
			BeginAt: "1457-10-16",
			EndAt:   "1461-2-1",
		},
		rawEraData{
			Name:    "寛正",
			Yomi:    "かんしょう",
			BeginAt: "1461-2-1",
			EndAt:   "1466-3-14",
		},
		rawEraData{
			Name:    "文正",
			Yomi:    "ぶんしょう",
			BeginAt: "1466-3-14",
			EndAt:   "1467-4-9",
		},
		rawEraData{
			Name:    "文明",
			Yomi:    "ぶんめい",
			BeginAt: "1469-6-8",
			EndAt:   "1487-8-9",
		},
		rawEraData{
			Name:    "長享",
			Yomi:    "ちょうきょう",
			BeginAt: "1487-8-9",
			EndAt:   "1489-9-16",
		},
		rawEraData{
			Name:    "延徳",
			Yomi:    "えんとく",
			BeginAt: "1489-9-16",
			EndAt:   "1492-8-12",
		},
		rawEraData{
			Name:    "明応",
			Yomi:    "めいおう",
			BeginAt: "1492-8-12",
			EndAt:   "1501-3-18",
		},
		rawEraData{
			Name:    "文亀",
			Yomi:    "ぶんき",
			BeginAt: "1501-3-18",
			EndAt:   "1504-3-16",
		},
		rawEraData{
			Name:    "永正",
			Yomi:    "えいしょう",
			BeginAt: "1504-3-16",
			EndAt:   "1521-9-23",
		},
		rawEraData{
			Name:    "大永",
			Yomi:    "だいえい",
			BeginAt: "1521-9-23",
			EndAt:   "1528-9-3",
		},
		rawEraData{
			Name:    "享禄",
			Yomi:    "きょうろく",
			BeginAt: "1528-9-3",
			EndAt:   "1532-8-29",
		},
		rawEraData{
			Name:    "天文",
			Yomi:    "てんぶん",
			BeginAt: "1532-8-29",
			EndAt:   "1555-11-7",
		},
		rawEraData{
			Name:    "弘治",
			Yomi:    "こうじ",
			BeginAt: "1555-11-7",
			EndAt:   "1558-3-18",
		},
		rawEraData{
			Name:    "永禄",
			Yomi:    "えいろく",
			BeginAt: "1558-3-18",
			EndAt:   "1570-5-27",
		},
		rawEraData{
			Name:    "元亀",
			Yomi:    "げんき",
			BeginAt: "1570-5-27",
			EndAt:   "1573-8-25",
		},
		rawEraData{
			Name:    "文禄",
			Yomi:    "ぶんろく",
			BeginAt: "1593-1-10",
			EndAt:   "1596-12-16",
		},
		rawEraData{
			Name:    "慶長",
			Yomi:    "けいちょう",
			BeginAt: "1596-12-16",
			EndAt:   "1615-9-5",
		},
		rawEraData{
			Name:    "寛永",
			Yomi:    "かんえい",
			BeginAt: "1624-4-17",
			EndAt:   "1645-1-13",
		},
		rawEraData{
			Name:    "正保",
			Yomi:    "しょうほう",
			BeginAt: "1645-1-13",
			EndAt:   "1648-4-7",
		},
		rawEraData{
			Name:    "慶安",
			Yomi:    "けいあん",
			BeginAt: "1648-4-7",
			EndAt:   "1652-10-20",
		},
		rawEraData{
			Name:    "承応",
			Yomi:    "じょうおう",
			BeginAt: "1652-10-20",
			EndAt:   "1655-5-18",
		},
		rawEraData{
			Name:    "明暦",
			Yomi:    "めいれき",
			BeginAt: "1655-5-18",
			EndAt:   "1658-8-21",
		},
		rawEraData{
			Name:    "万治",
			Yomi:    "まんじ",
			BeginAt: "1658-8-21",
			EndAt:   "1661-5-23",
		},
		rawEraData{
			Name:    "寛文",
			Yomi:    "かんぶん",
			BeginAt: "1661-5-23",
			EndAt:   "1673-10-30",
		},
		rawEraData{
			Name:    "延宝",
			Yomi:    "えんぽう",
			BeginAt: "1673-10-30",
			EndAt:   "1681-11-9",
		},
		rawEraData{
			Name:    "天和",
			Yomi:    "てんな",
			BeginAt: "1681-11-9",
			EndAt:   "1684-4-5",
		},
		rawEraData{
			Name:    "貞享",
			Yomi:    "じょうきょう",
			BeginAt: "1684-4-5",
			EndAt:   "1688-10-23",
		},
		rawEraData{
			Name:    "元禄",
			Yomi:    "げんろく",
			BeginAt: "1688-10-23",
			EndAt:   "1704-4-16",
		},
		rawEraData{
			Name:    "宝永",
			Yomi:    "ほうえい",
			BeginAt: "1704-4-16",
			EndAt:   "1711-6-11",
		},
		rawEraData{
			Name:    "正徳",
			Yomi:    "しょうとく",
			BeginAt: "1711-6-11",
			EndAt:   "1716-8-9",
		},
		rawEraData{
			Name:    "享保",
			Yomi:    "きょうほう",
			BeginAt: "1716-8-9",
			EndAt:   "1736-6-7",
		},
		rawEraData{
			Name:    "元文",
			Yomi:    "げんぶん",
			BeginAt: "1736-6-7",
			EndAt:   "1741-4-12",
		},
		rawEraData{
			Name:    "寛保",
			Yomi:    "かんぽう",
			BeginAt: "1741-4-12",
			EndAt:   "1744-4-3",
		},
		rawEraData{
			Name:    "延享",
			Yomi:    "えんきょう",
			BeginAt: "1744-4-3",
			EndAt:   "1748-8-5",
		},
		rawEraData{
			Name:    "寛延",
			Yomi:    "かんえん",
			BeginAt: "1748-8-5",
			EndAt:   "1751-12-14",
		},
		rawEraData{
			Name:    "宝暦",
			Yomi:    "ほうれき",
			BeginAt: "1751-12-14",
			EndAt:   "1764-6-30",
		},
		rawEraData{
			Name:    "明和",
			Yomi:    "めいわ",
			BeginAt: "1764-6-30",
			EndAt:   "1772-12-10",
		},
		rawEraData{
			Name:    "安永",
			Yomi:    "あんえい",
			BeginAt: "1772-12-10",
			EndAt:   "1781-4-25",
		},
		rawEraData{
			Name:    "天明",
			Yomi:    "てんめい",
			BeginAt: "1781-4-25",
			EndAt:   "1789-2-19",
		},
		rawEraData{
			Name:    "寛政",
			Yomi:    "かんせい",
			BeginAt: "1789-2-19",
			EndAt:   "1801-3-19",
		},
		rawEraData{
			Name:    "享和",
			Yomi:    "きょうわ",
			BeginAt: "1801-3-19",
			EndAt:   "1804-3-22",
		},
		rawEraData{
			Name:    "文化",
			Yomi:    "ぶんか",
			BeginAt: "1804-3-22",
			EndAt:   "1818-5-26",
		},
		rawEraData{
			Name:    "文政",
			Yomi:    "ぶんせい",
			BeginAt: "1818-5-26",
			EndAt:   "1831-1-23",
		},
		rawEraData{
			Name:    "天保",
			Yomi:    "てんぽう",
			BeginAt: "1831-1-23",
			EndAt:   "1845-1-9",
		},
		rawEraData{
			Name:    "弘化",
			Yomi:    "こうか",
			BeginAt: "1845-1-9",
			EndAt:   "1848-4-1",
		},
		rawEraData{
			Name:    "嘉永",
			Yomi:    "かえい",
			BeginAt: "1848-4-1",
			EndAt:   "1855-1-15",
		},
		rawEraData{
			Name:    "安政",
			Yomi:    "あんせい",
			BeginAt: "1855-1-15",
			EndAt:   "1860-4-8",
		},
		rawEraData{
			Name:    "万延",
			Yomi:    "まんえん",
			BeginAt: "1860-4-8",
			EndAt:   "1861-3-29",
		},
		rawEraData{
			Name:    "文久",
			Yomi:    "ぶんきゅう",
			BeginAt: "1861-3-29",
			EndAt:   "1864-3-27",
		},
		rawEraData{
			Name:    "元治",
			Yomi:    "げんじ",
			BeginAt: "1864-3-27",
			EndAt:   "1865-5-1",
		},
		rawEraData{
			Name:    "慶応",
			Yomi:    "けいおう",
			BeginAt: "1865-5-1",
			EndAt:   "1868-10-23",
		},
		rawEraData{
			Name:    "明治",
			Yomi:    "めいじ",
			BeginAt: "1868-1-25",
			EndAt:   "1912-7-29",
		},
		rawEraData{
			Name:    "大正",
			Yomi:    "たいしょう",
			BeginAt: "1912-7-30",
			EndAt:   "1926-12-24",
		},
		rawEraData{
			Name:    "昭和",
			Yomi:    "しょうわ",
			BeginAt: "1926-12-25",
			EndAt:   "1989-1-7",
		},
		rawEraData{
			Name:    "平成",
			Yomi:    "へいせい",
			BeginAt: "1989-1-8",
			EndAt:   "",
		},
	}

	eras, err := parseEra(rawEras)
	if err != nil {
		return nil, err
	}

	return eras, nil
}

func constructTrieFromEras(eras []Era) (*trie.RuneTrie, error) {
	trie := trie.NewRuneTrie()
	for _, era := range eras {
		jsonEra, _ := json.Marshal(era)
		trie.Put(era.Name, jsonEra)
	}
	return trie, nil
}

func parseEra(rawData []rawEraData) ([]Era, error) {
	eras := make([]Era, 0, len(rawData))

	for _, data := range rawData {
		era := Era{}
		era.Name = data.Name
		era.Yomi = data.Yomi

		era.BeginYear = -1
		era.BeginMonth = -1
		era.BeginDay = -1
		era.EndYear = -1
		era.EndMonth = -1
		era.EndDay = -1

		var err error
		beginAtSlice := strings.Split(data.BeginAt, "-")
		beginYear, err := strconv.Atoi(beginAtSlice[0])
		if err != nil {
			eras = append(eras, era)
			continue
		}
		beginMonth, err := strconv.Atoi(beginAtSlice[1])
		if err != nil {
			eras = append(eras, era)
			continue
		}
		beginDay, err := strconv.Atoi(beginAtSlice[2])
		if err != nil {
			eras = append(eras, era)
			continue
		}

		era.BeginYear = beginYear
		era.BeginMonth = beginMonth
		era.BeginDay = beginDay

		endAtSlice := strings.Split(data.EndAt, "-")
		endYear, err := strconv.Atoi(endAtSlice[0])
		if err != nil {
			eras = append(eras, era)
			continue
		}
		endMonth, err := strconv.Atoi(endAtSlice[1])
		if err != nil {
			eras = append(eras, era)
			continue
		}
		endDay, err := strconv.Atoi(endAtSlice[2])
		if err != nil {
			eras = append(eras, era)
			continue
		}

		era.EndYear = endYear
		era.EndMonth = endMonth
		era.EndDay = endDay

		eras = append(eras, era)
	}

	return eras, nil
}
