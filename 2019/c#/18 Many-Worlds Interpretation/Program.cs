using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Security.Cryptography;
using System.Threading.Tasks;
using System.Xml.Linq;

namespace _18_Many_Worlds_Interpretation
{
    class Program
    {
        public static int shortestFullPath = int.MaxValue; //I got this from an invalid answer. 2698 i too many steps.
        public static Dictionary<string, int> visited = new Dictionary<string, int>();
        static void Main(string[] args)
        {
            Console.WriteLine(DateTime.Now.ToShortTimeString());
            var lines = File.ReadAllLines("data.txt");
            var maze = new Maze();
            maze.Init(lines);

            int totalSteps = 0;

            List<MazeItem> path;

            (_, path) = ShortestPath(maze, new List<MazeItem>());

            totalSteps = path.Sum(i => i.Steps);
            
            Console.WriteLine("totalSteps");
            Console.WriteLine(totalSteps);

            //Console.WriteLine(String.Join(",", path.Select(p => p.Name)));
        }

        public static (Maze, List<MazeItem>) ShortestPath(Maze maze, List<MazeItem> path){
            //Console.WriteLine(String.Join(",", path.Select(p => p.Name)));

            var current = maze.FindAll(i => i == "@").First();
            maze.Distance(current);

            var minMaze = maze.Clone();
            var minPath = path;
            var keys = maze.FindAll(i => Char.IsLower(i[0])).Where(k => k.Steps > 0).ToList();
            var pathLength = path.Sum(i => i.Steps);

            if(pathLength > shortestFullPath) {
                return (maze, path);
            }

            var collectedKeys = string.Join("", path.OrderBy(mi => mi.Name).Select(mi => mi.Name));
            var curHashKey = current.X.ToString() + ":" + current.Y.ToString() + ":" + collectedKeys;
            lock (visited)
            {
                if (visited.ContainsKey(curHashKey))
                {
                    var oldSteps = visited[curHashKey];
                    if (oldSteps > pathLength)
                    {
                        visited[curHashKey] = pathLength;
                    }
                    else
                    {
                        path.Add(new MazeItem() { Steps = 100000, Name = "ö" });
                        return (maze, path);
                    }
                }
                else
                {
                    visited.Add(curHashKey, pathLength);
                }
            }

            if (!keys.Any()){
                if(pathLength < shortestFullPath){
                    shortestFullPath = pathLength;//
                    Console.WriteLine(shortestFullPath);
                    Console.WriteLine(String.Join(",", path.Select(p => p.Name)));
                }
                return (maze, path);
            }

            //keys.Shuffle();
            keys = keys.OrderBy(k => k.Name).ToList();
            Parallel.ForEach(keys, key =>
            //foreach(var key in keys)
            {
                int tmpSteps = 0;
                var tmpMaze = maze.Clone();
                var tmpPath = path.Select(item => item.Clone()).ToList();
                tmpPath.Add(key);


                (tmpSteps, tmpMaze) = GoToKey(key, tmpMaze);

                (tmpMaze, tmpPath) = ShortestPath(tmpMaze, tmpPath);

                if (tmpPath.Sum(i => i.Steps) < shortestFullPath)
                {
                    minPath = tmpPath.Select(item => item.Clone()).ToList(); ;
                    minMaze = tmpMaze.Clone();
                }

            }
            );
            return (minMaze, minPath);
            
            
        }
        public static (int, Maze) GoToKey(MazeItem key, Maze maze){
            var current = maze.FindAll(i => i == "@").First();
            maze.Distance(current);
            //totalSteps += key.Steps;
            //Console.WriteLine("Key: " + key.Name);

            var doors = maze.FindAll(i =>Char.IsUpper(i[0]));
            maze.SetMapItem(current.X, current.Y, ".");
            maze.SetMapItem(key.X, key.Y, "@");

            if (doors.Any())
            {
                var door = doors.Where(d => d.Name.ToLowerInvariant() == key.Name).FirstOrDefault();
                if(door != null){
                    maze.SetMapItem(door.X, door.Y, ".");
                }
            }
            else
            {
                //maze.DumpMap(1);
                //Console.Write("No doors left ");
                //break;
                //return (-1, null);
            }
            

            return (key.Steps, maze);

        }

     
    }

    static class MyExtensions{
        private static Random rng = new Random();  

        public static void Shuffle<T>(this IList<T> list)  
        {  
            int n = list.Count;  
            while (n > 1) {  
                n--;  
                int k = rng.Next(n + 1);  
                T value = list[k];  
                list[k] = list[n];  
                list[n] = value;  
            }  
        }   
    }
}
