namespace LibTests;
using Lib;


class pram
{
    public string input1;
    public string input2;
    public string expected;
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestSolution()
    {
        pram[] tt = new pram[] {
        new pram { input1 = "45", input2 = "1", expected = "1451" },
        new pram { input1 = "13", input2 = "200", expected = "1320013"},
        new pram { input1 = "Soon", input2 = "Me", expected = "MeSoonMe" },
        new pram { input1 = "U", input2 = "False", expected = "UFalseU"},};

        foreach (var t in tt)
        {
            string actual = Main.Solution(t.input1, t.input2);
            Assert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestSolutionA()
    {
        pram[] tt = new pram[] {
        new pram { input1 = "45", input2 = "1", expected = "1451" },
        new pram { input1 = "13", input2 = "200", expected = "1320013"},
        new pram { input1 = "Soon", input2 = "Me", expected = "MeSoonMe" },
        new pram { input1 = "U", input2 = "False", expected = "UFalseU"},};

        foreach (var t in tt)
        {
            string actual = Main.SolutionA(t.input1, t.input2);
            Assert.AreEqual(t.expected, actual);
        }
    }
}