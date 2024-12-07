class Program
{
    static void Main()
    {
        string[] lines = File.ReadAllLines("../input.txt");
        long totalResult = 0;

        foreach (string line in lines)
        {
            string[] parts = line.Split(":");

            long testValue = long.Parse(parts[0].Trim());
            long[] numbers = Array.ConvertAll(parts[1].Trim().Split(' '), long.Parse);

            if (MatchesTestValue(testValue, numbers))
            {
                totalResult += testValue;
            }
        }

        Console.WriteLine("total: " + totalResult);
    }

    private static bool MatchesTestValue(long testValue, long[] numbers)
    {
        return TryOperators(numbers, 1, numbers[0], testValue);
    }

    private static bool TryOperators(long[] numbers, int index, long currResult, long target)
    {
        if (index == numbers.Length)
        {
            return currResult == target;
        }

        long nextNumber = numbers[index];

        if (TryOperators(numbers, index + 1, currResult + nextNumber, target))
        {
            return true;
        }

        if (TryOperators(numbers, index + 1, currResult * nextNumber, target))
        {
            return true;
        }

        long combinedNumber = CombineString(currResult, nextNumber);
        if (TryOperators(numbers, index + 1, combinedNumber, target))
        {
            return true;
        }

        return false;
    }

    static long CombineString(long left, long right)
    {
        string combinedString = left.ToString() + right.ToString();
        return long.Parse(combinedString);
    }
}
