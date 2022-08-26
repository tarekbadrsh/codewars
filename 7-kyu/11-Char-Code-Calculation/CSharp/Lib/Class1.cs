namespace Lib;
public class Main
{
    public static string GetMiddle(string s)
    {
        int n = s.Length;
        if (n % 2 == 1)
        {
            return s.Substring(n / 2, 1);
        }
        return s.Substring(n / 2 - 1, 2);

    }

    ///================ other practices ==================///
    public static string GetMiddleA(string s)
    {
        return string.IsNullOrEmpty(s)
            ? s
            : s.Substring((s.Length - 1) / 2, 2 - s.Length % 2);
    }
}
