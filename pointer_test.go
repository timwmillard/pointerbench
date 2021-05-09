// Pointer Test
// go test -bench=. -benchmem -size=10000

package slice

import (
	"flag"
	"math/rand"
	"testing"
)

var size = flag.Int("size", 100, "size of test slice")

// Record -----

func BenchmarkToRecords(b *testing.B) {

	records := randomRecords(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecords(records)
	}
}

func BenchmarkToRecordsPointer(b *testing.B) {

	records := randomRecordsPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecordsPointer(records)
	}
}

func randomRecord() Record {
	return Record{
		N1: rand.Int(),
	}
}

func randomRecords(n int) []Record {
	records := make([]Record, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord()
	}
	return records
}

func randomRecordPointer() *Record {
	return &Record{
		N1: rand.Int(),
	}
}

func randomRecordsPointer(n int) []*Record {
	records := make([]*Record, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecordPointer()
	}
	return records
}

// Record8 -----

func BenchmarkToRecord8s(b *testing.B) {

	records := randomRecord8s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord8s(records)
	}
}

func BenchmarkToRecord8sPointer(b *testing.B) {

	records := randomRecord8sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord8sPointer(records)
	}
}

func randomRecord8() Record8 {
	return Record8{
		N1: rand.Int(),
		N2: rand.Int(),
		N3: rand.Int(),
		N4: rand.Int(),
		N5: rand.Int(),
		N6: rand.Int(),
		N7: rand.Int(),
		N8: rand.Int(),
	}
}

func randomRecord8s(n int) []Record8 {
	records := make([]Record8, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord8()
	}
	return records
}

func randomRecord8Pointer() *Record8 {
	return &Record8{
		N1: rand.Int(),
		N2: rand.Int(),
		N3: rand.Int(),
		N4: rand.Int(),
		N5: rand.Int(),
		N6: rand.Int(),
		N7: rand.Int(),
		N8: rand.Int(),
	}
}

func randomRecord8sPointer(n int) []*Record8 {
	records := make([]*Record8, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord8Pointer()
	}
	return records
}

// Record16 -----

func BenchmarkToRecord16s(b *testing.B) {

	records := randomRecord16s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord16s(records)
	}
}

func BenchmarkToRecord16sPointer(b *testing.B) {

	records := randomRecord16sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord16sPointer(records)
	}
}

func randomRecord16() Record16 {
	return Record16{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
	}
}

func randomRecord16s(n int) []Record16 {
	records := make([]Record16, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord16()
	}
	return records
}

func randomRecord16Pointer() *Record16 {
	return &Record16{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
	}
}

func randomRecord16sPointer(n int) []*Record16 {
	records := make([]*Record16, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord16Pointer()
	}
	return records
}

// Record24 -----

func BenchmarkToRecord24s(b *testing.B) {

	records := randomRecord24s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord24s(records)
	}
}

func BenchmarkToRecord24sPointer(b *testing.B) {

	records := randomRecord24sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord24sPointer(records)
	}
}

func randomRecord24() Record24 {
	return Record24{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
	}
}

func randomRecord24s(n int) []Record24 {
	records := make([]Record24, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord24()
	}
	return records
}

func randomRecord24Pointer() *Record24 {
	return &Record24{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
	}
}

func randomRecord24sPointer(n int) []*Record24 {
	records := make([]*Record24, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord24Pointer()
	}
	return records
}

// Record32 -----

func BenchmarkToRecord32s(b *testing.B) {

	records := randomRecord32s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord32s(records)
	}
}

func BenchmarkToRecord32sPointer(b *testing.B) {

	records := randomRecord32sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord32sPointer(records)
	}
}

func randomRecord32() Record32 {
	return Record32{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
		N25: rand.Int(),
		N26: rand.Int(),
		N27: rand.Int(),
		N28: rand.Int(),
		N29: rand.Int(),
		N30: rand.Int(),
		N31: rand.Int(),
		N32: rand.Int(),
	}
}

func randomRecord32s(n int) []Record32 {
	records := make([]Record32, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord32()
	}
	return records
}

func randomRecord32Pointer() *Record32 {
	return &Record32{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
		N25: rand.Int(),
		N26: rand.Int(),
		N27: rand.Int(),
		N28: rand.Int(),
		N29: rand.Int(),
		N30: rand.Int(),
		N31: rand.Int(),
		N32: rand.Int(),
	}
}

func randomRecord32sPointer(n int) []*Record32 {
	records := make([]*Record32, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord32Pointer()
	}
	return records
}

// Record64 -----

func BenchmarkToRecord64s(b *testing.B) {

	records := randomRecord64s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord64s(records)
	}
}

func BenchmarkToRecord64sPointer(b *testing.B) {

	records := randomRecord64sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord64sPointer(records)
	}
}

func randomRecord64() Record64 {
	return Record64{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
		N25: rand.Int(),
		N26: rand.Int(),
		N27: rand.Int(),
		N28: rand.Int(),
		N29: rand.Int(),
		N30: rand.Int(),
		N31: rand.Int(),
		N32: rand.Int(),
		N33: rand.Int(),
		N34: rand.Int(),
		N35: rand.Int(),
		N36: rand.Int(),
		N37: rand.Int(),
		N38: rand.Int(),
		N39: rand.Int(),
		N40: rand.Int(),
		N41: rand.Int(),
		N42: rand.Int(),
		N43: rand.Int(),
		N44: rand.Int(),
		N45: rand.Int(),
		N46: rand.Int(),
		N47: rand.Int(),
		N48: rand.Int(),
		N49: rand.Int(),
		N50: rand.Int(),
		N51: rand.Int(),
		N52: rand.Int(),
		N53: rand.Int(),
		N54: rand.Int(),
		N55: rand.Int(),
		N56: rand.Int(),
		N57: rand.Int(),
		N58: rand.Int(),
		N59: rand.Int(),
		N60: rand.Int(),
		N61: rand.Int(),
		N62: rand.Int(),
		N63: rand.Int(),
		N64: rand.Int(),
	}
}

func randomRecord64s(n int) []Record64 {
	records := make([]Record64, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord64()
	}
	return records
}

func randomRecord64Pointer() *Record64 {
	return &Record64{
		N1:  rand.Int(),
		N2:  rand.Int(),
		N3:  rand.Int(),
		N4:  rand.Int(),
		N5:  rand.Int(),
		N6:  rand.Int(),
		N7:  rand.Int(),
		N8:  rand.Int(),
		N9:  rand.Int(),
		N10: rand.Int(),
		N11: rand.Int(),
		N12: rand.Int(),
		N13: rand.Int(),
		N14: rand.Int(),
		N15: rand.Int(),
		N16: rand.Int(),
		N17: rand.Int(),
		N18: rand.Int(),
		N19: rand.Int(),
		N20: rand.Int(),
		N21: rand.Int(),
		N22: rand.Int(),
		N23: rand.Int(),
		N24: rand.Int(),
		N25: rand.Int(),
		N26: rand.Int(),
		N27: rand.Int(),
		N28: rand.Int(),
		N29: rand.Int(),
		N30: rand.Int(),
		N31: rand.Int(),
		N32: rand.Int(),
		N33: rand.Int(),
		N34: rand.Int(),
		N35: rand.Int(),
		N36: rand.Int(),
		N37: rand.Int(),
		N38: rand.Int(),
		N39: rand.Int(),
		N40: rand.Int(),
		N41: rand.Int(),
		N42: rand.Int(),
		N43: rand.Int(),
		N44: rand.Int(),
		N45: rand.Int(),
		N46: rand.Int(),
		N47: rand.Int(),
		N48: rand.Int(),
		N49: rand.Int(),
		N50: rand.Int(),
		N51: rand.Int(),
		N52: rand.Int(),
		N53: rand.Int(),
		N54: rand.Int(),
		N55: rand.Int(),
		N56: rand.Int(),
		N57: rand.Int(),
		N58: rand.Int(),
		N59: rand.Int(),
		N60: rand.Int(),
		N61: rand.Int(),
		N62: rand.Int(),
		N63: rand.Int(),
		N64: rand.Int(),
	}
}

func randomRecord64sPointer(n int) []*Record64 {
	records := make([]*Record64, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord64Pointer()
	}
	return records
}

// Record128 -----

func BenchmarkToRecord128s(b *testing.B) {

	records := randomRecord128s(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord128s(records)
	}
}

func BenchmarkToRecord128sPointer(b *testing.B) {

	records := randomRecord128sPointer(*size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toRecord128sPointer(records)
	}
}

func randomRecord128() Record128 {
	return Record128{
		N1:   rand.Int(),
		N2:   rand.Int(),
		N3:   rand.Int(),
		N4:   rand.Int(),
		N5:   rand.Int(),
		N6:   rand.Int(),
		N7:   rand.Int(),
		N8:   rand.Int(),
		N9:   rand.Int(),
		N10:  rand.Int(),
		N11:  rand.Int(),
		N12:  rand.Int(),
		N13:  rand.Int(),
		N14:  rand.Int(),
		N15:  rand.Int(),
		N16:  rand.Int(),
		N17:  rand.Int(),
		N18:  rand.Int(),
		N19:  rand.Int(),
		N20:  rand.Int(),
		N21:  rand.Int(),
		N22:  rand.Int(),
		N23:  rand.Int(),
		N24:  rand.Int(),
		N25:  rand.Int(),
		N26:  rand.Int(),
		N27:  rand.Int(),
		N28:  rand.Int(),
		N29:  rand.Int(),
		N30:  rand.Int(),
		N31:  rand.Int(),
		N32:  rand.Int(),
		N33:  rand.Int(),
		N34:  rand.Int(),
		N35:  rand.Int(),
		N36:  rand.Int(),
		N37:  rand.Int(),
		N38:  rand.Int(),
		N39:  rand.Int(),
		N40:  rand.Int(),
		N41:  rand.Int(),
		N42:  rand.Int(),
		N43:  rand.Int(),
		N44:  rand.Int(),
		N45:  rand.Int(),
		N46:  rand.Int(),
		N47:  rand.Int(),
		N48:  rand.Int(),
		N49:  rand.Int(),
		N50:  rand.Int(),
		N51:  rand.Int(),
		N52:  rand.Int(),
		N53:  rand.Int(),
		N54:  rand.Int(),
		N55:  rand.Int(),
		N56:  rand.Int(),
		N57:  rand.Int(),
		N58:  rand.Int(),
		N59:  rand.Int(),
		N60:  rand.Int(),
		N61:  rand.Int(),
		N62:  rand.Int(),
		N63:  rand.Int(),
		N64:  rand.Int(),
		N65:  rand.Int(),
		N66:  rand.Int(),
		N67:  rand.Int(),
		N68:  rand.Int(),
		N69:  rand.Int(),
		N70:  rand.Int(),
		N71:  rand.Int(),
		N72:  rand.Int(),
		N73:  rand.Int(),
		N74:  rand.Int(),
		N75:  rand.Int(),
		N76:  rand.Int(),
		N77:  rand.Int(),
		N78:  rand.Int(),
		N79:  rand.Int(),
		N80:  rand.Int(),
		N81:  rand.Int(),
		N82:  rand.Int(),
		N83:  rand.Int(),
		N84:  rand.Int(),
		N85:  rand.Int(),
		N86:  rand.Int(),
		N87:  rand.Int(),
		N88:  rand.Int(),
		N89:  rand.Int(),
		N90:  rand.Int(),
		N91:  rand.Int(),
		N92:  rand.Int(),
		N93:  rand.Int(),
		N94:  rand.Int(),
		N95:  rand.Int(),
		N96:  rand.Int(),
		N97:  rand.Int(),
		N98:  rand.Int(),
		N99:  rand.Int(),
		N100: rand.Int(),
		N101: rand.Int(),
		N102: rand.Int(),
		N103: rand.Int(),
		N104: rand.Int(),
		N105: rand.Int(),
		N106: rand.Int(),
		N107: rand.Int(),
		N108: rand.Int(),
		N109: rand.Int(),
		N110: rand.Int(),
		N111: rand.Int(),
		N112: rand.Int(),
		N113: rand.Int(),
		N114: rand.Int(),
		N115: rand.Int(),
		N116: rand.Int(),
		N117: rand.Int(),
		N118: rand.Int(),
		N119: rand.Int(),
		N120: rand.Int(),
		N121: rand.Int(),
		N122: rand.Int(),
		N123: rand.Int(),
		N124: rand.Int(),
		N125: rand.Int(),
		N126: rand.Int(),
		N127: rand.Int(),
		N128: rand.Int(),
	}
}

func randomRecord128s(n int) []Record128 {
	records := make([]Record128, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord128()
	}
	return records
}

func randomRecord128Pointer() *Record128 {
	return &Record128{
		N1:   rand.Int(),
		N2:   rand.Int(),
		N3:   rand.Int(),
		N4:   rand.Int(),
		N5:   rand.Int(),
		N6:   rand.Int(),
		N7:   rand.Int(),
		N8:   rand.Int(),
		N9:   rand.Int(),
		N10:  rand.Int(),
		N11:  rand.Int(),
		N12:  rand.Int(),
		N13:  rand.Int(),
		N14:  rand.Int(),
		N15:  rand.Int(),
		N16:  rand.Int(),
		N17:  rand.Int(),
		N18:  rand.Int(),
		N19:  rand.Int(),
		N20:  rand.Int(),
		N21:  rand.Int(),
		N22:  rand.Int(),
		N23:  rand.Int(),
		N24:  rand.Int(),
		N25:  rand.Int(),
		N26:  rand.Int(),
		N27:  rand.Int(),
		N28:  rand.Int(),
		N29:  rand.Int(),
		N30:  rand.Int(),
		N31:  rand.Int(),
		N32:  rand.Int(),
		N33:  rand.Int(),
		N34:  rand.Int(),
		N35:  rand.Int(),
		N36:  rand.Int(),
		N37:  rand.Int(),
		N38:  rand.Int(),
		N39:  rand.Int(),
		N40:  rand.Int(),
		N41:  rand.Int(),
		N42:  rand.Int(),
		N43:  rand.Int(),
		N44:  rand.Int(),
		N45:  rand.Int(),
		N46:  rand.Int(),
		N47:  rand.Int(),
		N48:  rand.Int(),
		N49:  rand.Int(),
		N50:  rand.Int(),
		N51:  rand.Int(),
		N52:  rand.Int(),
		N53:  rand.Int(),
		N54:  rand.Int(),
		N55:  rand.Int(),
		N56:  rand.Int(),
		N57:  rand.Int(),
		N58:  rand.Int(),
		N59:  rand.Int(),
		N60:  rand.Int(),
		N61:  rand.Int(),
		N62:  rand.Int(),
		N63:  rand.Int(),
		N64:  rand.Int(),
		N65:  rand.Int(),
		N66:  rand.Int(),
		N67:  rand.Int(),
		N68:  rand.Int(),
		N69:  rand.Int(),
		N70:  rand.Int(),
		N71:  rand.Int(),
		N72:  rand.Int(),
		N73:  rand.Int(),
		N74:  rand.Int(),
		N75:  rand.Int(),
		N76:  rand.Int(),
		N77:  rand.Int(),
		N78:  rand.Int(),
		N79:  rand.Int(),
		N80:  rand.Int(),
		N81:  rand.Int(),
		N82:  rand.Int(),
		N83:  rand.Int(),
		N84:  rand.Int(),
		N85:  rand.Int(),
		N86:  rand.Int(),
		N87:  rand.Int(),
		N88:  rand.Int(),
		N89:  rand.Int(),
		N90:  rand.Int(),
		N91:  rand.Int(),
		N92:  rand.Int(),
		N93:  rand.Int(),
		N94:  rand.Int(),
		N95:  rand.Int(),
		N96:  rand.Int(),
		N97:  rand.Int(),
		N98:  rand.Int(),
		N99:  rand.Int(),
		N100: rand.Int(),
		N101: rand.Int(),
		N102: rand.Int(),
		N103: rand.Int(),
		N104: rand.Int(),
		N105: rand.Int(),
		N106: rand.Int(),
		N107: rand.Int(),
		N108: rand.Int(),
		N109: rand.Int(),
		N110: rand.Int(),
		N111: rand.Int(),
		N112: rand.Int(),
		N113: rand.Int(),
		N114: rand.Int(),
		N115: rand.Int(),
		N116: rand.Int(),
		N117: rand.Int(),
		N118: rand.Int(),
		N119: rand.Int(),
		N120: rand.Int(),
		N121: rand.Int(),
		N122: rand.Int(),
		N123: rand.Int(),
		N124: rand.Int(),
		N125: rand.Int(),
		N126: rand.Int(),
		N127: rand.Int(),
		N128: rand.Int(),
	}
}

func randomRecord128sPointer(n int) []*Record128 {
	records := make([]*Record128, n)
	for i := 0; i < n; i++ {
		records[i] = randomRecord128Pointer()
	}
	return records
}
