namespace LibTests;
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
    public void TestGetMiddle()
    {
        pram[] tt = new pram[] {
        new pram {input1 = "test",  expected = "es"},
        new pram {input1 = "testing", expected = "t"},
        new pram {input1 = "middle", expected = "dd"},
        new pram {input1 = "A", expected = "A"}};

        foreach (var t in tt)
        {
            string actual = Main.GetMiddle(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestGetMiddleA()
    {
        pram[] tt = new pram[] {
        new pram {input1 = "test",  expected = "es"},
        new pram {input1 = "testing", expected = "t"},
        new pram {input1 = "middle", expected = "dd"},
        new pram {input1 = "A", expected = "A"}};

        foreach (var t in tt)
        {
            string actual = Main.GetMiddleA(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }
}