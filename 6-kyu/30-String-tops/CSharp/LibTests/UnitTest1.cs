namespace LibTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using Lib;

class pram
{
    public string input1;
    public string expected;
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestSolution()
    {
        pram[] tt = new pram[] {
        new pram { input1 = "", expected = "" },
        new pram { input1 = "12", expected = "2"},
        new pram { input1 = "abcdefghijklmnopqrstuvwxyz12345", expected = "3pgb" },
        new pram { input1 = "abcdefghijklmnopqrstuvwxyz1236789ABCDEFGHIJKLMN", expected = "M3pgb"},
        new pram { input1 = "syF2fh1h6e5nb8t69nAetAFbpFqcH4CHjfpu6 k1dfbzxspw8uisu1Dgqx9fc4EqwzG zqeHiaDDgxw3el7emzzGC4onrEDxymE3DbDwp0zzfbttByeh 87gowGtzz'GipmEw5'34zgjswzHG6zxDi2h22nAHC1lgBj 5srG7G mDoqm50wpGmskamec5cwC''3pHdzF FcqHo umb0e2ssBmyjGsAG4ov2ssACuFbmvp 4n k 12iC5CAgp9s67ktFbtsrpvpqhCiHGxoh31nn d9f6bukem7d'm9sDe9wfi5wzc8Grbw8mzdCCAFyiDf boq9rhzCEk4Ez k9d2b5wzadBsmdqqwlqG2dB3nwqdrnqBHCbDxomm8rd2DHBdni2d 107hBr2dapprhkjfmm8ocEgodqfhnBr47bmxi bq5havafqdqkFm1HtG6Fxgwgwbpm3E8Dgs3hDiDulAFo2 xvfBpmmn'zerEi BdxsjkeDEDEDjme F1z5c9 '7sgvrbCdzGw2zwe5dv839wHqd2488p8ulwoaqlxopubGf1y0B xxhqlFrx4gjFgmorbBmwAzHtqwDd4h1jvjABuFj2m6pAviFG9fkdz2wv6sm wGq5wrdelGvkzjAlr qp8ooknn2lCcpmBih7yjC eu3vdi7ddA1qb4imbidfro8eqE8haoEd0fbqm4eDq2 kdbg49CxG2i mbqwvv9uz59zluvdw6zcaqmjx01DD4E'bwm4q191qwEaC6mebzBn391udfl 546Dku9gu2B xx3p7'3GmgcowqGp2veses bp1jqxFp417bdvA'AEb8tsoxz2y1fm yEzFlsbg DmmAw2psc'vEn03ewojb5yzpAmkzlbpjpE0n2bsH1quhzs v'hqCq2wi7zwmE8jrE23lpzl9wtqCgdhBehC9wp6CmlzD0qejp6EygtqzHo 4dlo0GFuExbd   Fy4bb2Ep'3uh8zq1lCmFutzdk", expected = "qc6qeBDfrq1uw2onGsH61y" },};

        foreach (var t in tt)
        {
            string actual = Main.Tops(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestSolutionA()
    {
        pram[] tt = new pram[] {
        new pram { input1 = "", expected = "" },
        new pram { input1 = "12", expected = "2"},
        new pram { input1 = "abcdefghijklmnopqrstuvwxyz12345", expected = "3pgb" },
        new pram { input1 = "abcdefghijklmnopqrstuvwxyz1236789ABCDEFGHIJKLMN", expected = "M3pgb"},
        new pram { input1 = "syF2fh1h6e5nb8t69nAetAFbpFqcH4CHjfpu6 k1dfbzxspw8uisu1Dgqx9fc4EqwzG zqeHiaDDgxw3el7emzzGC4onrEDxymE3DbDwp0zzfbttByeh 87gowGtzz'GipmEw5'34zgjswzHG6zxDi2h22nAHC1lgBj 5srG7G mDoqm50wpGmskamec5cwC''3pHdzF FcqHo umb0e2ssBmyjGsAG4ov2ssACuFbmvp 4n k 12iC5CAgp9s67ktFbtsrpvpqhCiHGxoh31nn d9f6bukem7d'm9sDe9wfi5wzc8Grbw8mzdCCAFyiDf boq9rhzCEk4Ez k9d2b5wzadBsmdqqwlqG2dB3nwqdrnqBHCbDxomm8rd2DHBdni2d 107hBr2dapprhkjfmm8ocEgodqfhnBr47bmxi bq5havafqdqkFm1HtG6Fxgwgwbpm3E8Dgs3hDiDulAFo2 xvfBpmmn'zerEi BdxsjkeDEDEDjme F1z5c9 '7sgvrbCdzGw2zwe5dv839wHqd2488p8ulwoaqlxopubGf1y0B xxhqlFrx4gjFgmorbBmwAzHtqwDd4h1jvjABuFj2m6pAviFG9fkdz2wv6sm wGq5wrdelGvkzjAlr qp8ooknn2lCcpmBih7yjC eu3vdi7ddA1qb4imbidfro8eqE8haoEd0fbqm4eDq2 kdbg49CxG2i mbqwvv9uz59zluvdw6zcaqmjx01DD4E'bwm4q191qwEaC6mebzBn391udfl 546Dku9gu2B xx3p7'3GmgcowqGp2veses bp1jqxFp417bdvA'AEb8tsoxz2y1fm yEzFlsbg DmmAw2psc'vEn03ewojb5yzpAmkzlbpjpE0n2bsH1quhzs v'hqCq2wi7zwmE8jrE23lpzl9wtqCgdhBehC9wp6CmlzD0qejp6EygtqzHo 4dlo0GFuExbd   Fy4bb2Ep'3uh8zq1lCmFutzdk", expected = "qc6qeBDfrq1uw2onGsH61y" },};

        foreach (var t in tt)
        {
            string actual = Main.TopsA(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }
}