public class Helpers{
    public static int[] stringToIntArray(string input){
                    var ret = new List<int>();
            foreach(var c in input){
                ret.Add(int.Parse(c));
            }

            return ret.ToArray();
    }
}