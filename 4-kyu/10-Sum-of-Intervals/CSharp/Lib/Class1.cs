using System.Collections.Generic;
using System.Linq;
using System;

namespace Lib;
public class Main
{
    public static int SumIntervals((int, int)[] intervals)
    {
        if (intervals.Length == 0)
        {
            return 0;
        }
        Array.Sort(intervals, (x, y) => x.Item1.CompareTo(y.Item1));
        int sum = 0;
        int low = intervals[0].Item1;
        foreach (var v in intervals)
        {
            if (v.Item2 >= low)
            {
                sum += v.Item2;
                if (v.Item1 > low)
                {
                    sum -= v.Item1;
                }
                else
                {
                    sum -= low;
                }
                low = v.Item2;
            }
        }
        return sum;
    }
}
