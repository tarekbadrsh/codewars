using System.Collections.Generic;
using System.Linq;

namespace Lib;
public class Main
{
    public static object[] Alternate(int n, object firstValue, object secondValue)
    {
        object[] res = new object[n];
        for (int i = 0; i < n; i++)
        {
            if (i % 2 == 0)
            {
                res[i] = firstValue;
            }
            else if (i % 2 == 1)
            {
                res[i] = secondValue;
            }
        }
        return res;
    }

    ///================ other practices ==================///
    public static object[] AlternateA(int n, object firstValue, object secondValue)
    {
        return Enumerable.Range(0, n).Select(x => x % 2 == 1 ? secondValue : firstValue).ToArray();
    }
}
