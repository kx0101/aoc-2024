class Program
{
    static void Main()
    {
        string[] input = File.ReadAllLines("../input.txt");

        List<string> rules = new List<string>();
        List<string> updates = new List<string>();

        int i = 0;
        while (!string.IsNullOrWhiteSpace(input[i]))
        {
            rules.Add(input[i]);
            i++;
        }

        while (i < input.Length)
        {
            updates.Add(input[i]);
            i++;
        }

        List<(int, int)> rulePairs = rules.Select(rule =>
        {
            var parts = rule.Split('|').Select(int.Parse).ToArray();
            return (parts[0], parts[1]);
        }).ToList();

        int sumOfMiddlePages = 0;

        foreach (string update in updates)
        {
            var pages = update.Split(',').Select(int.Parse).ToList();

            if (!IsValidOrder(pages, rulePairs))
            {
                var sortedPages = TopologicalSort(pages, rulePairs);
                int middlePage = sortedPages[sortedPages.Count / 2];
                sumOfMiddlePages += middlePage;
            }
        }

        Console.WriteLine("Sum of middle pages: " + sumOfMiddlePages);
    }

    static bool IsValidOrder(List<int> pages, List<(int, int)> rules)
    {
        foreach (var rule in rules)
        {
            int indexX = pages.IndexOf(rule.Item1);
            int indexY = pages.IndexOf(rule.Item2);

            if (indexX != -1 && indexY != -1 && indexX > indexY)
            {
                return false;
            }
        }

        return true;
    }

    static List<int> TopologicalSort(List<int> pages, List<(int, int)> rules)
    {
        Dictionary<int, List<int>> graph = pages.ToDictionary(page => page, _ => new List<int>());
        Dictionary<int, int> inComingEdges = pages.ToDictionary(page => page, _ => 0);

        foreach (var rule in rules)
        {
            if (pages.Contains(rule.Item1) && pages.Contains(rule.Item2))
            {
                graph[rule.Item1].Add(rule.Item2);
                inComingEdges[rule.Item2]++;
            }
        }

        Queue<int> queue = new Queue<int>(inComingEdges.Where(pair => pair.Value == 0).Select(pair => pair.Key));
        List<int> sorted = new List<int>();

        while (queue.Count > 0)
        {
            int current = queue.Dequeue();
            sorted.Add(current);

            foreach (var neighbor in graph[current])
            {
                inComingEdges[neighbor]--;

                if (inComingEdges[neighbor] == 0)
                {
                    queue.Enqueue(neighbor);
                }
            }
        }

        return sorted;
    }
}
