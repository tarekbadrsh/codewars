namespace Lib;

using System.Linq;

public class Main
{
    public static bool IsValidWalk(string[] walk)
    {
        if (walk.Length != 10)
        {
            return false;
        }
        int x = 0;
        int y = 0;
        foreach (var r in walk)
        {
            switch (r)
            {
                case "s":
                    y--;
                    break;
                case "n":
                    y++;
                    break;
                case "w":
                    x--;
                    break;
                case "e":
                    x++;
                    break;
            }
        }
        return x == 0 && y == 0;
    }

    ///================ other practices ==================///
    public static bool IsValidWalkA(string[] walk)
    {
        return walk.Count() == 10 &&
            walk.Count(c => c == "n") == walk.Count(c => c == "s") &&
            walk.Count(c => c == "w") == walk.Count(c => c == "e");
    }
}