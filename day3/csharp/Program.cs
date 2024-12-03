using System.Text.RegularExpressions;

namespace Aoc3
{
    public class Program
    {
        public static void Main(string[] args)
        {
            string input = File.ReadAllText("../input.txt");

            string mulPattern = @"mul\((\d+),(\d+)\)";
            string togglePattern = @"(do\(\)|don't\(\))";

            bool isMulEnabled = true;
            int sum = 0;

            int currPosition = 0;

            while (currPosition < input.Length)
            {
                Match mulMatch = Regex.Match(input.Substring(currPosition), mulPattern);
                Match toggleMatch = Regex.Match(input.Substring(currPosition), togglePattern);

                if (mulMatch.Success && (mulMatch.Index < toggleMatch.Index || !toggleMatch.Success))
                {
                    int x = int.Parse(mulMatch.Groups[1].Value);
                    int y = int.Parse(mulMatch.Groups[2].Value);

                    if (isMulEnabled)
                    {
                        sum += x * y;
                    }

                    currPosition += mulMatch.Index + mulMatch.Length;
                }
                else if (toggleMatch.Success)
                {
                    if (toggleMatch.Value == "do()")
                    {
                        isMulEnabled = true;
                    }

                    else if (toggleMatch.Value == "don't()")
                    {
                        isMulEnabled = false;
                    }

                    currPosition += toggleMatch.Index + toggleMatch.Length;
                }
                else
                {
                    currPosition++;
                }
            }

            Console.WriteLine(sum);
        }
    }
}
