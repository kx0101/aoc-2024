namespace Aoc1
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string[] lines = File.ReadAllLines("../input.txt");
            List<int> numbers1 = new List<int>();
            List<int> numbers2 = new List<int>();

            foreach (var line in lines)
            {
                var splitLine = line.Split(' ', StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries);

                numbers1.Add(int.Parse(splitLine[0]));
                numbers2.Add(int.Parse(splitLine[1]));
            }

            numbers1.Sort();
            numbers2.Sort();

            int sumDifference = numbers1.Zip(numbers2, (num1, num2) => Math.Abs(num1 - num2)).Sum();
            Console.WriteLine(sumDifference);

            int weightedSum = numbers1.Sum(num => num * numbers2.Count(x => x == num));
            Console.WriteLine(weightedSum);
        }
    }
}
