using System;
using System.Collections.Generic;
using System.Text;

namespace _18_Many_Worlds_Interpretation
{
    public class Maze
    {
        private int[,] _pointMap;
        private string[,] _map;
        private int _maxX;
        private int _maxY;

        public string Wall;
        public string Room;
        public List<MazeItem> Items;
        
        public void Init(string[] lines)
        {
            Wall = "#";
            Room = ".";
            Items = new List<MazeItem>();

            _pointMap = new int[lines[0].Length, lines.Length];
            _map = new string[lines[0].Length, lines.Length];
            _maxX = lines[0].Length - 1;
            _maxY = lines.Length - 1;

            InitMapsAndItems(lines);
        }

        private void InitMapsAndItems(string[] lines)
        {
            for (int x = 0; x < lines[0].Length; x++)
            {
                for (int y = 0; y < lines.Length; y++)
                {
                    _pointMap[x, y] = -1;
                    var line = lines[y];
                    var itemName = line[x].ToString();

                    _map[x, y] = itemName;

                    if (itemName != Wall && itemName != Room)
                    {
                        var i = new MazeItem
                        {
                            X = x,
                            Y = y,
                            Name = itemName
                        };
                        Items.Add(i);
                    }

                }
            }
        }
        private void InitPointMap()
        {
            for (int x = 0; x <= _maxX; x++)
            {
                for (int y = 0; y <= _maxY; y++)
                {
                    _pointMap[x, y] = -1;
                }
            }
        }

        public int Distance(MazeItem origin, MazeItem dest)
        {
            
            return Distance(origin.X, origin.Y, dest.X, dest.Y);
            

        }
        
        public List<MazeItem> Walk(MazeItem origin, MazeItem dest)
        {
            InitPointMap();
            return walk(origin, dest, new List<MazeItem>());
        }
        public List<MazeItem> walk(MazeItem origin, MazeItem dest, List<MazeItem> steps)
        {
            if(origin.X == dest.X && origin.Y == dest.Y)
            {
                return steps;
            }

            var y = origin.Y;
            var x = origin.X;
            List<MazeItem> newSteps ;

            _pointMap[x, y] = 0;

            steps.Add(origin);
            var nextStep = new MazeItem() { ItemType = Room };
            nextStep.X = origin.X;
            nextStep.Y = origin.Y;

            if (y > 0 && _map[x, y - 1] != Wall && _pointMap[x, y - 1] == -1)
            {
                nextStep.Y = origin.Y - 1;
                newSteps = walk(nextStep, dest, steps);
                if (newSteps != null) return newSteps;
            }

            if (y < _maxY && _map[x, y + 1] != Wall && _pointMap[x, y + 1] == -1)
            {
                nextStep.Y = origin.Y + 1;
                newSteps = walk(nextStep, dest, steps);
                if (newSteps != null) return newSteps;
            }

            if (x > 0 && _map[x - 1, y] != Wall && _pointMap[x - 1, y] == -1)
            {
                nextStep.X = origin.X - 1;
                newSteps = walk(nextStep, dest, steps);
                if (newSteps != null) return newSteps;
            }

            if (x < _maxX && _map[x + 1, y] != Wall && _pointMap[x + 1, y] == -1)
            {
                nextStep.X = origin.X + 1;
                newSteps = walk(nextStep, dest, steps);
                if (newSteps != null) return newSteps;
            }


            return null;
        }

        public int Distance(int x1, int y1, int x2, int y2)
        {
            InitPointMap();
            setDistance(x1, y1, 0);
            return _pointMap[x2, y2];

        }

        private void setDistance(int x, int y, int step)
        {
            _pointMap[x, y] = step;
            if (y > 0 && _map[x, y - 1] != Wall && _pointMap[x, y - 1] == -1) setDistance(x, y - 1, step + 1);
            if (y < _maxY  && _map[x, y + 1] != Wall && _pointMap[x, y + 1] == -1) setDistance(x, y + 1, step + 1);
            if (x > 0 && _map[x-1, y] != Wall && _pointMap[x - 1, y] == -1) setDistance(x-1, y, step + 1);
            if (x < _maxX  && _map[x+1, y] != Wall && _pointMap[x + 1, y] == -1) setDistance(x+1, y, step + 1);

        }
        public void DumpPoints(int pad)
        {
            for (int y = 0; y <= _maxY; y++)
            {
                for (int x = 0; x <= _maxX; x++)
                {
                    
                    Console.Write(_pointMap[x, y].ToString().PadLeft(pad));
                }
                Console.WriteLine();
            }

        }

    }

    public class MazeItem
    {
        public int X;
        public int Y;
        public string ItemType;
        public string Name;
    }
}
