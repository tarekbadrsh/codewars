namespace LibTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using Lib;

class pram
{
    public string[] input1;
    public bool expected;
}

[TestClass]
public class UnitTest1
{


    [TestMethod]
    public void TestSolution()
    {
        pram[] tt = new pram[] {
        new pram { input1 = new string[] {"n","s","n","s","n","s","n","s","n","s"}, expected = true },
        new pram { input1 = new string[] {"w","e","w","e","w","e","w","e","w","e","w","e"}, expected =  false},
        new pram { input1 =  new string[] { "w" }, expected = false },
        new pram { input1 = new string[] { "n", "n", "n", "s", "n", "s", "n", "s", "n", "s" }, expected = false},};

        foreach (var t in tt)
        {
            bool actual = Main.IsValidWalk(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestSolutionA()
    {
        pram[] tt = new pram[] {
        new pram { input1 = new string[] {"n","s","n","s","n","s","n","s","n","s"}, expected = true },
        new pram { input1 = new string[] {"w","e","w","e","w","e","w","e","w","e","w","e"}, expected =  false},
        new pram { input1 =  new string[] { "w" }, expected = false },
        new pram { input1 = new string[] { "n", "n", "n", "s", "n", "s", "n", "s", "n", "s" }, expected = false},};

        foreach (var t in tt)
        {
            bool actual = Main.IsValidWalkA(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }
}