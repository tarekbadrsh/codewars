using System.Collections.Generic;
using System.Linq;

namespace Lib;
public class Main
{
    public static long NextBiggerNumber(long n)
    {
        if (n < 12)
        {
            return -1;
        }
        string strBase = n.ToString();
        bool canBe = false;
        for (int i = 0; i < strBase.Length - 1; i++)
        {
            if (strBase[i] < strBase[i + 1])
            {
                canBe = true;
                break;
            }
        }
        if (!canBe)
        {
            return -1;
        }
        for (var nPlus = n + 1; ; nPlus++)
        {
            string strPlus = nPlus.ToString();
            bool valid = true;
            foreach (var item in strBase)
            {
                if (strBase.Count(c => c == item) != strPlus.Count(c => c == item))
                {
                    valid = false;
                    break;
                }
            }
            if (!valid)
            {
                continue;
            }
            return nPlus;
        }
    }

    ///================ other practices ==================///
    public static long NextBiggerNumberA(long n)
    {
        String str = GetNumbers(n);
        for (long i = n + 1; i <= long.Parse(str); i++)
        {
            if (GetNumbers(n) == GetNumbers(i))
            {
                return i;
            }
        }
        return -1;
    }
    public static string GetNumbers(long number)
    {
        return string.Join("", number.ToString().ToCharArray().OrderByDescending(x => x));
    }
}
