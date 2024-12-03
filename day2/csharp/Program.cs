namespace Aoc2
{
    public class Program
    {
        public static void Main()
        {
            string input = File.ReadAllText("../input.txt");

            var reports = input
                .Trim()
                .Split("\n")
                .Select(line =>
                        line
                        .Trim()
                        .Split(" ")
                        .Select(int.Parse)
                        .ToList())
                .ToList();

            int safe = 0;

            foreach (var report in reports)
            {
                if (IsSafe(report) || CanBeSafeWithOneRemove(report))
                {
                    safe++;
                }
            }

            Console.WriteLine("safe: " + safe);
        }

        static bool IsSafe(List<int> levels)
        {
            if (levels.Count < 2)
            {
                return false;
            }

            bool isIncreasing = levels[1] > levels[0];
            for (int i = 1; i < levels.Count; i++)
            {
                int diff = levels[i] - levels[i - 1];

                if (Math.Abs(diff) < 1 || Math.Abs(diff) > 3)
                {
                    return false;
                }

                if ((isIncreasing && diff < 0) || (!isIncreasing && diff > 0))
                {
                    return false;
                }
            }

            return true;
        }

        static bool CanBeSafeWithOneRemove(List<int> levels)
        {
            for (int i = 0; i < levels.Count; i++)
            {
                var modifiedReport = new List<int>(levels);
                modifiedReport.RemoveAt(i);

                if (IsSafe(modifiedReport))
                {
                    return true;
                }
            }

            return false;
        }
    }
}
