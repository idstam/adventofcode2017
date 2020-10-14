using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Security.Cryptography;
using System.Xml.Linq;

namespace _18_Many_Worlds_Interpretation
{
    class Program
    {
        static void Main(string[] args)
        {
            
            var lines = File.ReadAllLines("data.txt");
            var maze = new Maze();
            maze.Init(lines);


            //maze.DumpPoints(3);

            List<MazeItem> keys;
            int totalSteps = 0;

            do
            {
                var current = maze.FindAll(i => i == "@").First();
                maze.Distance(current);

                maze.DumpPoints(3);

                keys = maze.FindAll(i => Char.IsLower(i[0])).Where(k => k.Steps > 0).ToList();
                
                if (keys.Any())
                {
                    var key = keys[0];
                    var doors = maze.FindAll(i =>Char.IsUpper(i[0]) && i == key.Name.ToUpperInvariant());
                    if (doors.Any())
                    {
                        maze.SetMapItem(doors[0].X, doors[0].Y, ".");
                    }
                    else
                    {
                        Console.Write("No doors left");
                        break;
                    }
                    

                    totalSteps += key.Steps;
                    maze.SetMapItem(current.X, current.Y, ".");
                    maze.SetMapItem(key.X, key.Y, "@");

                }
            } while (keys.Any());

            Console.WriteLine("totalSteps");
            Console.WriteLine(totalSteps);

        }



        
    }
}
