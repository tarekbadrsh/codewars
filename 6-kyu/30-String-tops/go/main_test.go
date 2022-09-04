package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1   string
		expected string
	}{
		{"", ""},
		{"12", "2"},
		{"abcdefghijklmnopqrstuvwxyz12345", "3pgb"},
		{"abcdefghijklmnopqrstuvwxyz1236789ABCDEFGHIJKLMN", "M3pgb"},
		{"syF2fh1h6e5nb8t69nAetAFbpFqcH4CHjfpu6 k1dfbzxspw8uisu1Dgqx9fc4EqwzG zqeHiaDDgxw3el7emzzGC4onrEDxymE3DbDwp0zzfbttByeh 87gowGtzz'GipmEw5'34zgjswzHG6zxDi2h22nAHC1lgBj 5srG7G mDoqm50wpGmskamec5cwC''3pHdzF FcqHo umb0e2ssBmyjGsAG4ov2ssACuFbmvp 4n k 12iC5CAgp9s67ktFbtsrpvpqhCiHGxoh31nn d9f6bukem7d'm9sDe9wfi5wzc8Grbw8mzdCCAFyiDf boq9rhzCEk4Ez k9d2b5wzadBsmdqqwlqG2dB3nwqdrnqBHCbDxomm8rd2DHBdni2d 107hBr2dapprhkjfmm8ocEgodqfhnBr47bmxi bq5havafqdqkFm1HtG6Fxgwgwbpm3E8Dgs3hDiDulAFo2 xvfBpmmn'zerEi BdxsjkeDEDEDjme F1z5c9 '7sgvrbCdzGw2zwe5dv839wHqd2488p8ulwoaqlxopubGf1y0B xxhqlFrx4gjFgmorbBmwAzHtqwDd4h1jvjABuFj2m6pAviFG9fkdz2wv6sm wGq5wrdelGvkzjAlr qp8ooknn2lCcpmBih7yjC eu3vdi7ddA1qb4imbidfro8eqE8haoEd0fbqm4eDq2 kdbg49CxG2i mbqwvv9uz59zluvdw6zcaqmjx01DD4E'bwm4q191qwEaC6mebzBn391udfl 546Dku9gu2B xx3p7'3GmgcowqGp2veses bp1jqxFp417bdvA'AEb8tsoxz2y1fm yEzFlsbg DmmAw2psc'vEn03ewojb5yzpAmkzlbpjpE0n2bsH1quhzs v'hqCq2wi7zwmE8jrE23lpzl9wtqCgdhBehC9wp6CmlzD0qejp6EygtqzHo 4dlo0GFuExbd   Fy4bb2Ep'3uh8zq1lCmFutzdk", "qc6qeBDfrq1uw2onGsH61y"},
	}

	for _, tc := range tt {
		t.Run("Tops", func(t *testing.T) {
			res := Tops(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("Tops", func(t *testing.T) {
			res := TopsA(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkTops(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Tops("syF2fh1h6e5nb8t69nAetAFbpFqcH4CHjfpu6 k1dfbzxspw8uisu1Dgqx9fc4EqwzG zqeHiaDDgxw3el7emzzGC4onrEDxymE3DbDwp0zzfbttByeh 87gowGtzz'GipmEw5'34zgjswzHG6zxDi2h22nAHC1lgBj 5srG7G mDoqm50wpGmskamec5cwC''3pHdzF FcqHo umb0e2ssBmyjGsAG4ov2ssACuFbmvp 4n k 12iC5CAgp9s67ktFbtsrpvpqhCiHGxoh31nn d9f6bukem7d'm9sDe9wfi5wzc8Grbw8mzdCCAFyiDf boq9rhzCEk4Ez k9d2b5wzadBsmdqqwlqG2dB3nwqdrnqBHCbDxomm8rd2DHBdni2d 107hBr2dapprhkjfmm8ocEgodqfhnBr47bmxi bq5havafqdqkFm1HtG6Fxgwgwbpm3E8Dgs3hDiDulAFo2 xvfBpmmn'zerEi BdxsjkeDEDEDjme F1z5c9 '7sgvrbCdzGw2zwe5dv839wHqd2488p8ulwoaqlxopubGf1y0B xxhqlFrx4gjFgmorbBmwAzHtqwDd4h1jvjABuFj2m6pAviFG9fkdz2wv6sm wGq5wrdelGvkzjAlr qp8ooknn2lCcpmBih7yjC eu3vdi7ddA1qb4imbidfro8eqE8haoEd0fbqm4eDq2 kdbg49CxG2i mbqwvv9uz59zluvdw6zcaqmjx01DD4E'bwm4q191qwEaC6mebzBn391udfl 546Dku9gu2B xx3p7'3GmgcowqGp2veses bp1jqxFp417bdvA'AEb8tsoxz2y1fm yEzFlsbg DmmAw2psc'vEn03ewojb5yzpAmkzlbpjpE0n2bsH1quhzs v'hqCq2wi7zwmE8jrE23lpzl9wtqCgdhBehC9wp6CmlzD0qejp6EygtqzHo 4dlo0GFuExbd   Fy4bb2Ep'3uh8zq1lCmFutzdk")
	}
}

func BenchmarkTopsA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		TopsA("syF2fh1h6e5nb8t69nAetAFbpFqcH4CHjfpu6 k1dfbzxspw8uisu1Dgqx9fc4EqwzG zqeHiaDDgxw3el7emzzGC4onrEDxymE3DbDwp0zzfbttByeh 87gowGtzz'GipmEw5'34zgjswzHG6zxDi2h22nAHC1lgBj 5srG7G mDoqm50wpGmskamec5cwC''3pHdzF FcqHo umb0e2ssBmyjGsAG4ov2ssACuFbmvp 4n k 12iC5CAgp9s67ktFbtsrpvpqhCiHGxoh31nn d9f6bukem7d'm9sDe9wfi5wzc8Grbw8mzdCCAFyiDf boq9rhzCEk4Ez k9d2b5wzadBsmdqqwlqG2dB3nwqdrnqBHCbDxomm8rd2DHBdni2d 107hBr2dapprhkjfmm8ocEgodqfhnBr47bmxi bq5havafqdqkFm1HtG6Fxgwgwbpm3E8Dgs3hDiDulAFo2 xvfBpmmn'zerEi BdxsjkeDEDEDjme F1z5c9 '7sgvrbCdzGw2zwe5dv839wHqd2488p8ulwoaqlxopubGf1y0B xxhqlFrx4gjFgmorbBmwAzHtqwDd4h1jvjABuFj2m6pAviFG9fkdz2wv6sm wGq5wrdelGvkzjAlr qp8ooknn2lCcpmBih7yjC eu3vdi7ddA1qb4imbidfro8eqE8haoEd0fbqm4eDq2 kdbg49CxG2i mbqwvv9uz59zluvdw6zcaqmjx01DD4E'bwm4q191qwEaC6mebzBn391udfl 546Dku9gu2B xx3p7'3GmgcowqGp2veses bp1jqxFp417bdvA'AEb8tsoxz2y1fm yEzFlsbg DmmAw2psc'vEn03ewojb5yzpAmkzlbpjpE0n2bsH1quhzs v'hqCq2wi7zwmE8jrE23lpzl9wtqCgdhBehC9wp6CmlzD0qejp6EygtqzHo 4dlo0GFuExbd   Fy4bb2Ep'3uh8zq1lCmFutzdk")
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/6-kyu/30-String-tops/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkTops-8    	  323457	      3664 ns/op	    1088 B/op	      86 allocs/op
BenchmarkTopsA-8   	 1479146	       862.7 ns/op	     304 B/op	      22 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/6-kyu/30-String-tops/go	3.324s
*/
