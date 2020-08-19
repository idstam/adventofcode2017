using System;
using System.IO;

namespace _18_Many_Worlds_Interpretation
{
    class Program
    {
        static void Main(string[] args)
        {
            
            var lines = File.ReadAllLines("data.txt");
            var maze = new Maze();
            maze.Init(lines);

            foreach(var i in maze.Items)
            {
                if(i.Name == "@")
                {
                    i.ItemType = "entrance";
                }else if(i.Name == i.Name.ToLowerInvariant())
                {
                    i.ItemType = "key";
                }
                else if (i.Name == i.Name.ToUpperInvariant())
                {
                    i.ItemType = "door";
                }
            }

            maze.DumpPoints(3);

            var entrance = maze.Items.Find(i => i.Name == "@");
            var a = maze.Items.Find(i => i.Name == "a");

            var steps = maze.Walk(entrance, a);


        }



        
    }
}
