using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace _16._1_Flawed_Frequency_Transmission
{
    class Program
    {
        static void Main(string[] args)
        {
            
            var stringInput = "59773775883217736423759431065821647306353420453506194937477909478357279250527959717515453593953697526882172199401149893789300695782513381578519607082690498042922082853622468730359031443128253024761190886248093463388723595869794965934425375464188783095560858698959059899665146868388800302242666479679987279787144346712262803568907779621609528347260035619258134850854741360515089631116920328622677237196915348865412336221562817250057035898092525020837239100456855355496177944747496249192354750965666437121797601987523473707071258599440525572142300549600825381432815592726865051526418740875442413571535945830954724825314675166862626566783107780527347044";
            //stringInput = "03036732577212944063491565474664";
            //stringInput = "12345678";

            var s = new StringBuilder();
            for (int i = 0; i < 10000; i++)
            {
                s.Append(stringInput);
            }
            var input = Helpers.stringToIntArray(s.ToString());
            var output = new int[input.Length];
            
            var foo = "";
            for (int i = 0; i < 7; i++)
            {
                foo += input[i].ToString();

            }
            var offset = int.Parse(foo);
            Console.WriteLine(offset);


            s.Clear();

            for (int phase = 0; phase < 100; phase++)
            {
                Console.WriteLine("Phase: " + phase.ToString() + " " +  DateTime.Now.ToLongTimeString());

                Parallel.For(offset, output.Length, new ParallelOptions() { MaxDegreeOfParallelism = 6 }, (y) =>
                //for (int y = 0; y < output.Length; y++)
                {
                    int foo = 0;
                    //int pn = 0;
                    for (int x = y; x < output.Length; x++)
                    {
                        //pn = 1; //getPatNum(x, y);
                        
                        //if (pn != 0)
                        //{
                            var b = Math.Min((x + y), output.Length - 1);
                            for (int i = x; i <= b; i++)
                            {
                                foo += input[i];
                            }
                        //}
                        x += y;
                    }
                    output[y] = Math.Abs(foo % 10);// * pn;
                }
                );

                input = output;
                output = new int[input.Length];
            }
            
            Console.WriteLine(offset);
            
            for (int i = 0; i < 8; i++)
            {
                Console.Write(input[i + offset]);
            }
            Console.WriteLine();
            Console.WriteLine("84462026 facit");

            Console.WriteLine();

        }

        private static int getPatNum(int numX, int numY){

            int[] basePattern = {0,1,0,-1};
            var ret = basePattern[((numX+1)/(numY+1)) % 4];

            return ret;

        }

    }
}
