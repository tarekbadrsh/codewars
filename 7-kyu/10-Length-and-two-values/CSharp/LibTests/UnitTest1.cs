namespace LibTests;
using Lib;


class pram
{
    public int n;
    public object input1;
    public object input2;
    public object[] expected;
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestAlternate()
    {
        pram[] tt = new pram[] {
        new pram {n=5 , input1 = "true", input2 = "false", expected = new object[]{"true", "false", "true", "false", "true"} },
        new pram {n=20 , input1 = "blue", input2 = "red", expected = new object[]{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}},
         new pram {n=0 , input1 = "", input2 = "", expected = new object[]{}},};

        foreach (var t in tt)
        {
            object[] actual = Main.Alternate(t.n, t.input1, t.input2);
            Assert.AreEqual(t.expected.Length, actual.Length);
        }
    }

    [TestMethod]
    public void TestAlternateA()
    {
        pram[] tt = new pram[] {
        new pram {n=5 , input1 = "true", input2 = "false", expected = new object[]{"true", "false", "true", "false", "true"} },
        new pram {n=20 , input1 = "blue", input2 = "red", expected = new object[]{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}},
        new pram { n = 0, input1 = "", input2 = "", expected = new object[] { } },};

        foreach (var t in tt)
        {
            object[] actual = Main.AlternateA(t.n, t.input1, t.input2);
            Assert.AreEqual(t.expected.Length, actual.Length);
        }
    }
}