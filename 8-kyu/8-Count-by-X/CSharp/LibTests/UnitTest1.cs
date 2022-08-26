namespace LibTests;
using Lib;


class pram
{
    public int inputx;
    public int inputn;
    public int[] expected { get; set; }
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestCountBy()
    {
        pram[] tt = new pram[] {
        new pram { inputx = 1, inputn = 5, expected = new int[] { 1, 2, 3, 4, 5 } },
        new pram { inputx = 2, inputn = 5, expected = new int[] { 2, 4, 6, 8, 10 } },
        new pram { inputx = 3, inputn = 5, expected = new int[] { 3, 6, 9, 12, 15 } },
        new pram { inputx = 50, inputn = 5, expected = new int[] { 50, 100, 150, 200, 250 } },
        new pram { inputx = 100, inputn = 5, expected = new int[] { 100, 200, 300, 400, 500 } }};

        foreach (var t in tt)
        {
            int[] actual = Main.CountBy(t.inputx, t.inputn);
            CollectionAssert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestCountByA()
    {
        pram[] tt = new pram[] {
        new pram { inputx = 1, inputn = 5, expected = new int[] { 1, 2, 3, 4, 5 } },
        new pram { inputx = 2, inputn = 5, expected = new int[] { 2, 4, 6, 8, 10 } },
        new pram { inputx = 3, inputn = 5, expected = new int[] { 3, 6, 9, 12, 15 } },
        new pram { inputx = 50, inputn = 5, expected = new int[] { 50, 100, 150, 200, 250 } },
        new pram { inputx = 100, inputn = 5, expected = new int[] { 100, 200, 300, 400, 500 } }};

        foreach (var t in tt)
        {
            int[] actual = Main.CountByA(t.inputx, t.inputn);
            CollectionAssert.AreEqual(t.expected, actual);
        }
    }
}