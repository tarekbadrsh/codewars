using System.Collections.Generic;
using System.Linq;

namespace Lib;
public class Main
{
    public static int[] CountBy(int x, int n)
    {
        List<int> termsList = new List<int>();

        for (int i = x; i <= (x * n); i += x)
        {
            termsList.Add(i);
        }

        return termsList.ToArray();
    }

    public static int[] CountByA(int x, int n)
    {
        return Enumerable.Range(1, n).Select(i => i * x).ToArray();
    }
}
