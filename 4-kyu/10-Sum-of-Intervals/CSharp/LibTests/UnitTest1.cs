namespace LibTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using Lib;

using Interval = System.ValueTuple<int, int>;

class pram
{
    public Interval[] input1;
    public int expected;
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestSolution()
    {
        pram[] tt = new pram[] {
        new pram {input1 = new Interval[] { (4, 4), (6, 6), (8, 8) }, expected = 0},
        new pram {input1 = new Interval[] {  (1, 5), (10, 20), (1, 6), (16, 19), (5, 11) }, expected = 19},
        };

        foreach (var t in tt)
        {
            int actual = Main.SumIntervals(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }
}