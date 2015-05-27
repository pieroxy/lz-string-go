package encoding

import (
	"testing"
)

var testingValues map[string]string = map[string]string{
	"IZA":   "a",
	"IbI":   "aa",
	"Aco":   "ààà",
	"Ac1A":  "àààààà",
	"C4e1Q": "toto",
	"KORCJCIEBFhgRAU4WASCwiQ1ESCBvwjIBA":                                                                                              "Å∑⁄¢∆Ê‚√∫ﬂı",
	"MYUwBALgTghglgOzCAzhMADsATVFEz4D2CQA":                                                                                            "ce train est à destination",
	"MIUwBALgTghglgOzCAzhMADsATVFEz4D2SuYACgIIBKAkgMpj2W0ByAKmADKUBaNAUQBcYAGpEArimRIAVxFSYwAFgCsAUgDcYAPIS4YAMYByPJCgAL6QBsSAcwB0QA": "Ce train est à destination de PARIS SAINT LAZARE: Vous en êtes à 45%; Oui c'est très long.",
	"N4IglgJiBcIIwCYDMAWArKgbAtOUgBoQBDAFxgSNIFtiAPGRABiLGoAcYBtUSGeQiABGxAHaiApgCcYoAO6M0ADiREAFoyYBONEXYB7AM4wA7ERGlSM6FzhIAugF9HBAAS8osBIJHjpskAVoEyUUdU0WEANjaDDhMituOydHe3NiCAA3bhAAY30OMQBPOAA6fOpBCvZihHKCkDSQQzBSCQC+L1Q0TCRBUWJqdtgAZVaJVwBBACEAYQARQQgC4jBRfjlN0uIAV1y1MQgpMEyJUoAzGT1iAHMJXLIcpBMmJkw4JkaiQ3vSMH1RA9yDYQM9Xth8E12MdMsRckUDAAbMDwxjXO78NSWdjQAD0uJabWIQlyEHq1FxNTupSxlSIUgk50x2LxuIZ52kx1EN0JZwqbMZnLWNypZ1pgnYOyEyMMan80A8-ElQkQ3X6g2GIAAClKZXKpFMQC48gC2qJgaAANYSIpyfRSCAxLgga22+0QVzEQSuu0O1xCb0230e3KNZxECASTIo4a8TiwTAoUqIJM4JDJwQ7L2wACy+gAXmBEYjiLi0KUmK4ABQ5uFrUhGNQAblcAFUWwBJc0SRGuWu5VwAeRGrgAGq4PqVMC2JKIALStkYtqSZaBlLSlOpwTAASlcAHF7pb9LiEEw4B8kHAtK4AGJgdn6Oi49OYZOYQRGfjDseCc4l2VTmseAmAzIgACsYjgY0dh+axFVgdBiEMCBzi0JQTEwHoQi0DJzgyFBMBCBk5CwoilC0HwdiKaQdk6EA0FQ8jkIgCiMLIljzhw9i0HwCMyGzHgQAGIZ+EmPYDlEeYYXaNxmgkG4hnNbhQBEzUREMa1yDk2FER2TUuHsI0nBSIA": "{\"id\":\"1234534625254\",\"at\":2,\"tmax\":120,\"imp\":[{\"id\":\"1\",\"banner\":{\"w\":1583,\"h\":1095,\"pos\":7,\"battr\":[13]}}, {\"id\":\"2\",\"banner\":{\"w\":784,\"h\":100,\"pos\":4,\"battr\":[13]}}],\"badv\":[\"company1.com\",\"company2.com\"],\"site\":{\"id\":\"234563\",\"name\":\"Site ABCD\",\"domain\":\"www.auchandrive.fr\",\"pagecat\":[\"3700610\"],\"sectioncat\":[\"3700624\"],\"privacypolicy\":1,\"page\":\"http://siteabcd.com/page.htm\",\"ref\":\"http://referringsite.com/referringpage.htm\",\"publisher\":{\"id\":\"pub12345\",\"name\":\"Publisher A\"},\"content\":{\"keywords\":[\"keyword a\",\"keyword b\",\"keyword c\"]}},\"device\":{\"ip\":\"64.124.253.1\",\"ua\":\"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16\",\"os\":\"OS X\",\"flashver\":\"10.1\",\"js\":1},\"user\":{\"id\":\"45asdf987656789adfad4678rew656789\",\"buyeruid\":\"5df678asd8987656asdf78987654\",\"data\":[{\"name\":\"AuchanDrive\", \"segment\":[{\"name\":\"basket\", \"value\":\"[]\"}]}]}}",
}

func TestDecompress(t *testing.T) {
	for k, v := range testingValues {
		result, err := DecompressFromEncodedUriComponent(k)
		if err != nil {
			t.Errorf("Unexpected error", err)
		}
		if result != v {
			t.Errorf("Result should be :\n", v, "\n instead of :\n", result)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testingValues {
			DecompressFromEncodedUriComponent(k)
		}
	}
}
