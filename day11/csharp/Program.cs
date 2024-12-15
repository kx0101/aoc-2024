class Program
{
    static void Main()
    {
        string input = File.ReadAllText("../input.txt").Trim();
        string[] initialStones = input.Split(' ');

        int blinks = 75;
        long totalStones = CountStonesAfterBlinks(initialStones, blinks);

        Console.WriteLine(totalStones);
    }

    static long CountStonesAfterBlinks(string[] initialStones, int totalBlinks)
    {
        Dictionary<string, long> currCounts = new Dictionary<string, long>();

        foreach (string stone in initialStones)
        {
            if (!currCounts.ContainsKey(stone))
            {
                currCounts[stone] = 0;
            }

            currCounts[stone]++;
        }

        for (int blink = 0; blink < totalBlinks; blink++)
        {
            Dictionary<string, long> nextCounts = new Dictionary<string, long>();

            foreach (var pair in currCounts)
            {
                string stone = pair.Key;
                long count = pair.Value;

                if (stone == "0")
                {
                    AddToDictionary(nextCounts, "1", count);
                }
                else if (stone.Length % 2 == 0)
                {
                    int mid = stone.Length / 2;
                    string left = stone.Substring(0, mid).TrimStart('0');
                    string right = stone.Substring(mid).TrimStart('0');

                    AddToDictionary(nextCounts, string.IsNullOrEmpty(left) ? "0" : left, count);
                    AddToDictionary(nextCounts, string.IsNullOrEmpty(right) ? "0" : right, count);
                }
                else
                {
                    long number = long.Parse(stone);
                    long transformed = number * 2024;

                    AddToDictionary(nextCounts, transformed.ToString(), count);
                }
            }

            currCounts = nextCounts;
        }

        long totalStones = 0;
        foreach (var count in currCounts.Values)
        {
            totalStones += count;
        }

        return totalStones;
    }

    static void AddToDictionary(Dictionary<string, long> dict, string key, long count)
    {
        if (!dict.ContainsKey(key))
        {
            dict[key] = 0;
        }

        dict[key] += count;
    }
}
