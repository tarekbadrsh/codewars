namespace Lib;
public class Main
{
    public static string Tops(string msg)
    {
        var c = 1;
        var tops = "";
        for (int i = 1; i < msg.Length; i = i + c)
        {
            tops = msg[i] + tops;
            c = c + 4;
        }
        return tops;
    }

    ///================ other practices ==================///
    public static string TopsA(string msg)
    {
        string s = "";
        for (int x = 1; x <= msg.Length; x += s.Length * 4)
            s = msg[x++] + s;
        return s;
    }
}
