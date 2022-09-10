namespace LibTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using Lib;

class pram
{
    public long input1;
    public long expected;
}

[TestClass]
public class UnitTest1
{
    [TestMethod]
    public void TestSolution()
    {
        pram[] tt = new pram[] {
        new pram {input1 = 12, expected = 21},
        new pram {input1= 2017, expected= 2071},
        new pram {input1= 414, expected= 441},
        new pram {input1= 144, expected= 414},
        new pram {input1= 8, expected= -1},
        new pram {input1= 123456789, expected= 123456798},
        new pram {input1= 1234567890, expected= 1234567908},
        new pram {input1= 9876543210, expected= -1},
        new pram {input1= 9999999999, expected= -1},
        new pram {input1= 59884848459853, expected= 59884848483559},
        new pram {input1= 44444, expected= -1},};

        foreach (var t in tt)
        {
            long actual = Main.NextBiggerNumber(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }

    [TestMethod]
    public void TestSolutionA()
    {
        pram[] tt = new pram[] {
        new pram {input1 = 12, expected = 21},
        new pram {input1= 2017, expected= 2071},
        new pram {input1= 414, expected= 441},
        new pram {input1= 144, expected= 414},
        new pram {input1= 8, expected= -1},
        new pram {input1= 123456789, expected= 123456798},
        new pram {input1= 1234567890, expected= 1234567908},
        new pram {input1= 9876543210, expected= -1},
        new pram {input1= 9999999999, expected= -1},
        new pram {input1= 59884848459853, expected= 59884848483559},
        new pram {input1= 44444, expected= -1},};
        foreach (var t in tt)
        {
            long actual = Main.NextBiggerNumberA(t.input1);
            Assert.AreEqual(t.expected, actual);
        }
    }
}