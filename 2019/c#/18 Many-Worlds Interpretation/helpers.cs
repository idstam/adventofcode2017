using System.Collections.Generic;

public class Helpers{
    public static int[] stringToIntArray(string input){
                    var ret = new List<int>();
            foreach(var c in input){
                ret.Add(int.Parse(c.ToString()));
            }

            return ret.ToArray();
    }

    public static string[] intArrayToStringArray(int[] input)
    {
        var ret = new List<string>();
        foreach(var i in input)
        {
            ret.Add(i.ToString());
        }

        return ret.ToArray();
    }

    
    
}