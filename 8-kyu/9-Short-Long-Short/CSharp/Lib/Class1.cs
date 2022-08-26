using System.Collections.Generic;
using System.Linq;

namespace Lib;
public class Main
{
    public static string Solution(string a, string b)
    {
        if (a.Length < b.Length)
        {
            return String.Format("{0}{1}{2}", a, b, a);
        }
        return String.Format("{0}{1}{2}", b, a, b);
    }

    ///================ other practices ==================///
    public static string SolutionA(string a, string b)
    {
        return (a.Length > b.Length) ? (b + a + b) : (a + b + a);
    }
}
