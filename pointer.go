package slice

type Record struct {
	N1 int
}

func toRecords(pp []Record) []Record {
	records := make([]Record, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord(p))
	}
	return records
}

func toRecordsPointer(pp []*Record) []*Record {
	records := make([]*Record, len(pp))
	for i, p := range pp {
		records[i] = toRecordPointer(p)
	}
	return records
}

func toRecord(p Record) Record {
	return Record{
		N1: p.N1,
	}
}

func toRecordPointer(p *Record) *Record {
	return &Record{
		N1: p.N1,
	}
}

type Record8 struct {
	N1 int
	N2 int
	N3 int
	N4 int
	N5 int
	N6 int
	N7 int
	N8 int
}

func toRecord8s(pp []Record8) []Record8 {
	records := make([]Record8, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord8(p))
	}
	return records
}

func toRecord8sPointer(pp []*Record8) []*Record8 {
	records := make([]*Record8, len(pp))
	for i, p := range pp {
		records[i] = toRecord8Pointer(p)
	}
	return records
}

func toRecord8(p Record8) Record8 {
	return Record8{
		N1: p.N1,
		N2: p.N2,
		N3: p.N3,
		N4: p.N4,
		N5: p.N5,
		N6: p.N6,
		N7: p.N7,
		N8: p.N8,
	}
}

func toRecord8Pointer(p *Record8) *Record8 {
	return &Record8{
		N1: p.N1,
		N2: p.N2,
		N3: p.N3,
		N4: p.N4,
		N5: p.N5,
		N6: p.N6,
		N7: p.N7,
		N8: p.N8,
	}
}

type Record16 struct {
	N1  int
	N2  int
	N3  int
	N4  int
	N5  int
	N6  int
	N7  int
	N8  int
	N9  int
	N10 int
	N11 int
	N12 int
	N13 int
	N14 int
	N15 int
	N16 int
}

func toRecord16s(pp []Record16) []Record16 {
	records := make([]Record16, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord16(p))
	}
	return records
}

func toRecord16sPointer(pp []*Record16) []*Record16 {
	records := make([]*Record16, len(pp))
	for i, p := range pp {
		records[i] = toRecord16Pointer(p)
	}
	return records
}

func toRecord16(p Record16) Record16 {
	return Record16{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
	}
}

func toRecord16Pointer(p *Record16) *Record16 {
	return &Record16{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
	}
}

type Record24 struct {
	N1  int
	N2  int
	N3  int
	N4  int
	N5  int
	N6  int
	N7  int
	N8  int
	N9  int
	N10 int
	N11 int
	N12 int
	N13 int
	N14 int
	N15 int
	N16 int
	N17 int
	N18 int
	N19 int
	N20 int
	N21 int
	N22 int
	N23 int
	N24 int
}

func toRecord24s(pp []Record24) []Record24 {
	records := make([]Record24, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord24(p))
	}
	return records
}

func toRecord24sPointer(pp []*Record24) []*Record24 {
	records := make([]*Record24, len(pp))
	for i, p := range pp {
		records[i] = toRecord24Pointer(p)
	}
	return records
}

func toRecord24(p Record24) Record24 {
	return Record24{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
	}
}

func toRecord24Pointer(p *Record24) *Record24 {
	return &Record24{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
	}
}

type Record32 struct {
	N1  int
	N2  int
	N3  int
	N4  int
	N5  int
	N6  int
	N7  int
	N8  int
	N9  int
	N10 int
	N11 int
	N12 int
	N13 int
	N14 int
	N15 int
	N16 int
	N17 int
	N18 int
	N19 int
	N20 int
	N21 int
	N22 int
	N23 int
	N24 int
	N25 int
	N26 int
	N27 int
	N28 int
	N29 int
	N30 int
	N31 int
	N32 int
}

func toRecord32s(pp []Record32) []Record32 {
	records := make([]Record32, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord32(p))
	}
	return records
}

func toRecord32sPointer(pp []*Record32) []*Record32 {
	records := make([]*Record32, len(pp))
	for i, p := range pp {
		records[i] = toRecord32Pointer(p)
	}
	return records
}

func toRecord32(p Record32) Record32 {
	return Record32{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
		N25: p.N25,
		N26: p.N26,
		N27: p.N27,
		N28: p.N28,
		N29: p.N29,
		N30: p.N30,
		N31: p.N31,
		N32: p.N32,
	}
}

func toRecord32Pointer(p *Record32) *Record32 {
	return &Record32{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
		N25: p.N25,
		N26: p.N26,
		N27: p.N27,
		N28: p.N28,
		N29: p.N29,
		N30: p.N30,
		N31: p.N31,
		N32: p.N32,
	}
}

type Record64 struct {
	N1  int
	N2  int
	N3  int
	N4  int
	N5  int
	N6  int
	N7  int
	N8  int
	N9  int
	N10 int
	N11 int
	N12 int
	N13 int
	N14 int
	N15 int
	N16 int
	N17 int
	N18 int
	N19 int
	N20 int
	N21 int
	N22 int
	N23 int
	N24 int
	N25 int
	N26 int
	N27 int
	N28 int
	N29 int
	N30 int
	N31 int
	N32 int
	N33 int
	N34 int
	N35 int
	N36 int
	N37 int
	N38 int
	N39 int
	N40 int
	N41 int
	N42 int
	N43 int
	N44 int
	N45 int
	N46 int
	N47 int
	N48 int
	N49 int
	N50 int
	N51 int
	N52 int
	N53 int
	N54 int
	N55 int
	N56 int
	N57 int
	N58 int
	N59 int
	N60 int
	N61 int
	N62 int
	N63 int
	N64 int
}

func toRecord64s(pp []Record64) []Record64 {
	records := make([]Record64, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord64(p))
	}
	return records
}

func toRecord64sPointer(pp []*Record64) []*Record64 {
	records := make([]*Record64, len(pp))
	for i, p := range pp {
		records[i] = toRecord64Pointer(p)
	}
	return records
}

func toRecord64(p Record64) Record64 {
	return Record64{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
		N25: p.N25,
		N26: p.N26,
		N27: p.N27,
		N28: p.N28,
		N29: p.N29,
		N30: p.N30,
		N31: p.N31,
		N32: p.N32,
		N33: p.N33,
		N34: p.N34,
		N35: p.N35,
		N36: p.N36,
		N37: p.N37,
		N38: p.N38,
		N39: p.N39,
		N40: p.N40,
		N41: p.N41,
		N42: p.N42,
		N43: p.N43,
		N44: p.N44,
		N45: p.N45,
		N46: p.N46,
		N47: p.N47,
		N48: p.N48,
		N49: p.N49,
		N50: p.N50,
		N51: p.N51,
		N52: p.N52,
		N53: p.N53,
		N54: p.N54,
		N55: p.N55,
		N56: p.N56,
		N57: p.N57,
		N58: p.N58,
		N59: p.N59,
		N60: p.N60,
		N61: p.N61,
		N62: p.N62,
		N63: p.N63,
		N64: p.N64,
	}
}

func toRecord64Pointer(p *Record64) *Record64 {
	return &Record64{
		N1:  p.N1,
		N2:  p.N2,
		N3:  p.N3,
		N4:  p.N4,
		N5:  p.N5,
		N6:  p.N6,
		N7:  p.N7,
		N8:  p.N8,
		N9:  p.N9,
		N10: p.N10,
		N11: p.N11,
		N12: p.N12,
		N13: p.N13,
		N14: p.N14,
		N15: p.N15,
		N16: p.N16,
		N17: p.N17,
		N18: p.N18,
		N19: p.N19,
		N20: p.N20,
		N21: p.N21,
		N22: p.N22,
		N23: p.N23,
		N24: p.N24,
		N25: p.N25,
		N26: p.N26,
		N27: p.N27,
		N28: p.N28,
		N29: p.N29,
		N30: p.N30,
		N31: p.N31,
		N32: p.N32,
		N33: p.N33,
		N34: p.N34,
		N35: p.N35,
		N36: p.N36,
		N37: p.N37,
		N38: p.N38,
		N39: p.N39,
		N40: p.N40,
		N41: p.N41,
		N42: p.N42,
		N43: p.N43,
		N44: p.N44,
		N45: p.N45,
		N46: p.N46,
		N47: p.N47,
		N48: p.N48,
		N49: p.N49,
		N50: p.N50,
		N51: p.N51,
		N52: p.N52,
		N53: p.N53,
		N54: p.N54,
		N55: p.N55,
		N56: p.N56,
		N57: p.N57,
		N58: p.N58,
		N59: p.N59,
		N60: p.N60,
		N61: p.N61,
		N62: p.N62,
		N63: p.N63,
		N64: p.N64,
	}
}

type Record128 struct {
	N1   int
	N2   int
	N3   int
	N4   int
	N5   int
	N6   int
	N7   int
	N8   int
	N9   int
	N10  int
	N11  int
	N12  int
	N13  int
	N14  int
	N15  int
	N16  int
	N17  int
	N18  int
	N19  int
	N20  int
	N21  int
	N22  int
	N23  int
	N24  int
	N25  int
	N26  int
	N27  int
	N28  int
	N29  int
	N30  int
	N31  int
	N32  int
	N33  int
	N34  int
	N35  int
	N36  int
	N37  int
	N38  int
	N39  int
	N40  int
	N41  int
	N42  int
	N43  int
	N44  int
	N45  int
	N46  int
	N47  int
	N48  int
	N49  int
	N50  int
	N51  int
	N52  int
	N53  int
	N54  int
	N55  int
	N56  int
	N57  int
	N58  int
	N59  int
	N60  int
	N61  int
	N62  int
	N63  int
	N64  int
	N65  int
	N66  int
	N67  int
	N68  int
	N69  int
	N70  int
	N71  int
	N72  int
	N73  int
	N74  int
	N75  int
	N76  int
	N77  int
	N78  int
	N79  int
	N80  int
	N81  int
	N82  int
	N83  int
	N84  int
	N85  int
	N86  int
	N87  int
	N88  int
	N89  int
	N90  int
	N91  int
	N92  int
	N93  int
	N94  int
	N95  int
	N96  int
	N97  int
	N98  int
	N99  int
	N100 int
	N101 int
	N102 int
	N103 int
	N104 int
	N105 int
	N106 int
	N107 int
	N108 int
	N109 int
	N110 int
	N111 int
	N112 int
	N113 int
	N114 int
	N115 int
	N116 int
	N117 int
	N118 int
	N119 int
	N120 int
	N121 int
	N122 int
	N123 int
	N124 int
	N125 int
	N126 int
	N127 int
	N128 int
}

func toRecord128s(pp []Record128) []Record128 {
	records := make([]Record128, 0, len(pp))
	for _, p := range pp {
		records = append(records, toRecord128(p))
	}
	return records
}

func toRecord128sPointer(pp []*Record128) []*Record128 {
	records := make([]*Record128, len(pp))
	for i, p := range pp {
		records[i] = toRecord128Pointer(p)
	}
	return records
}

func toRecord128(p Record128) Record128 {
	return Record128{
		N1:   p.N1,
		N2:   p.N2,
		N3:   p.N3,
		N4:   p.N4,
		N5:   p.N5,
		N6:   p.N6,
		N7:   p.N7,
		N8:   p.N8,
		N9:   p.N9,
		N10:  p.N10,
		N11:  p.N11,
		N12:  p.N12,
		N13:  p.N13,
		N14:  p.N14,
		N15:  p.N15,
		N16:  p.N16,
		N17:  p.N17,
		N18:  p.N18,
		N19:  p.N19,
		N20:  p.N20,
		N21:  p.N21,
		N22:  p.N22,
		N23:  p.N23,
		N24:  p.N24,
		N25:  p.N25,
		N26:  p.N26,
		N27:  p.N27,
		N28:  p.N28,
		N29:  p.N29,
		N30:  p.N30,
		N31:  p.N31,
		N32:  p.N32,
		N33:  p.N33,
		N34:  p.N34,
		N35:  p.N35,
		N36:  p.N36,
		N37:  p.N37,
		N38:  p.N38,
		N39:  p.N39,
		N40:  p.N40,
		N41:  p.N41,
		N42:  p.N42,
		N43:  p.N43,
		N44:  p.N44,
		N45:  p.N45,
		N46:  p.N46,
		N47:  p.N47,
		N48:  p.N48,
		N49:  p.N49,
		N50:  p.N50,
		N51:  p.N51,
		N52:  p.N52,
		N53:  p.N53,
		N54:  p.N54,
		N55:  p.N55,
		N56:  p.N56,
		N57:  p.N57,
		N58:  p.N58,
		N59:  p.N59,
		N60:  p.N60,
		N61:  p.N61,
		N62:  p.N62,
		N63:  p.N63,
		N64:  p.N64,
		N65:  p.N65,
		N66:  p.N66,
		N67:  p.N67,
		N68:  p.N68,
		N69:  p.N69,
		N70:  p.N70,
		N71:  p.N71,
		N72:  p.N72,
		N73:  p.N73,
		N74:  p.N74,
		N75:  p.N75,
		N76:  p.N76,
		N77:  p.N77,
		N78:  p.N78,
		N79:  p.N79,
		N80:  p.N80,
		N81:  p.N81,
		N82:  p.N82,
		N83:  p.N83,
		N84:  p.N84,
		N85:  p.N85,
		N86:  p.N86,
		N87:  p.N87,
		N88:  p.N88,
		N89:  p.N89,
		N90:  p.N90,
		N91:  p.N91,
		N92:  p.N92,
		N93:  p.N93,
		N94:  p.N94,
		N95:  p.N95,
		N96:  p.N96,
		N97:  p.N97,
		N98:  p.N98,
		N99:  p.N99,
		N100: p.N100,
		N101: p.N101,
		N102: p.N102,
		N103: p.N103,
		N104: p.N104,
		N105: p.N105,
		N106: p.N106,
		N107: p.N107,
		N108: p.N108,
		N109: p.N109,
		N110: p.N110,
		N111: p.N111,
		N112: p.N112,
		N113: p.N113,
		N114: p.N114,
		N115: p.N115,
		N116: p.N116,
		N117: p.N117,
		N118: p.N118,
		N119: p.N119,
		N120: p.N120,
		N121: p.N121,
		N122: p.N122,
		N123: p.N123,
		N124: p.N124,
		N125: p.N125,
		N126: p.N126,
		N127: p.N127,
		N128: p.N128,
	}
}

func toRecord128Pointer(p *Record128) *Record128 {
	return &Record128{
		N1:   p.N1,
		N2:   p.N2,
		N3:   p.N3,
		N4:   p.N4,
		N5:   p.N5,
		N6:   p.N6,
		N7:   p.N7,
		N8:   p.N8,
		N9:   p.N9,
		N10:  p.N10,
		N11:  p.N11,
		N12:  p.N12,
		N13:  p.N13,
		N14:  p.N14,
		N15:  p.N15,
		N16:  p.N16,
		N17:  p.N17,
		N18:  p.N18,
		N19:  p.N19,
		N20:  p.N20,
		N21:  p.N21,
		N22:  p.N22,
		N23:  p.N23,
		N24:  p.N24,
		N25:  p.N25,
		N26:  p.N26,
		N27:  p.N27,
		N28:  p.N28,
		N29:  p.N29,
		N30:  p.N30,
		N31:  p.N31,
		N32:  p.N32,
		N33:  p.N33,
		N34:  p.N34,
		N35:  p.N35,
		N36:  p.N36,
		N37:  p.N37,
		N38:  p.N38,
		N39:  p.N39,
		N40:  p.N40,
		N41:  p.N41,
		N42:  p.N42,
		N43:  p.N43,
		N44:  p.N44,
		N45:  p.N45,
		N46:  p.N46,
		N47:  p.N47,
		N48:  p.N48,
		N49:  p.N49,
		N50:  p.N50,
		N51:  p.N51,
		N52:  p.N52,
		N53:  p.N53,
		N54:  p.N54,
		N55:  p.N55,
		N56:  p.N56,
		N57:  p.N57,
		N58:  p.N58,
		N59:  p.N59,
		N60:  p.N60,
		N61:  p.N61,
		N62:  p.N62,
		N63:  p.N63,
		N64:  p.N64,
		N65:  p.N65,
		N66:  p.N66,
		N67:  p.N67,
		N68:  p.N68,
		N69:  p.N69,
		N70:  p.N70,
		N71:  p.N71,
		N72:  p.N72,
		N73:  p.N73,
		N74:  p.N74,
		N75:  p.N75,
		N76:  p.N76,
		N77:  p.N77,
		N78:  p.N78,
		N79:  p.N79,
		N80:  p.N80,
		N81:  p.N81,
		N82:  p.N82,
		N83:  p.N83,
		N84:  p.N84,
		N85:  p.N85,
		N86:  p.N86,
		N87:  p.N87,
		N88:  p.N88,
		N89:  p.N89,
		N90:  p.N90,
		N91:  p.N91,
		N92:  p.N92,
		N93:  p.N93,
		N94:  p.N94,
		N95:  p.N95,
		N96:  p.N96,
		N97:  p.N97,
		N98:  p.N98,
		N99:  p.N99,
		N100: p.N100,
		N101: p.N101,
		N102: p.N102,
		N103: p.N103,
		N104: p.N104,
		N105: p.N105,
		N106: p.N106,
		N107: p.N107,
		N108: p.N108,
		N109: p.N109,
		N110: p.N110,
		N111: p.N111,
		N112: p.N112,
		N113: p.N113,
		N114: p.N114,
		N115: p.N115,
		N116: p.N116,
		N117: p.N117,
		N118: p.N118,
		N119: p.N119,
		N120: p.N120,
		N121: p.N121,
		N122: p.N122,
		N123: p.N123,
		N124: p.N124,
		N125: p.N125,
		N126: p.N126,
		N127: p.N127,
		N128: p.N128,
	}
}
