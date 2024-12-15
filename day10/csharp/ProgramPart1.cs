class ProgramPart1
{
    static void main()
    {
        string[] lines = File.ReadAllLines("../input.txt");

        int rows = lines.Length;
        int cols = lines[0].Length;

        int[,] map = new int[rows, cols];

        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                map[i, j] = lines[i][j] - '0';
            }
        }

        int totalScore = 0;

        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                if (map[i, j] == 0)
                {
                    bool[,] visited = new bool[rows, cols];
                    totalScore += GetTrailheadScore(map, i, j, rows, cols, visited);
                }
            }
        }

        Console.WriteLine("total: " + totalScore);
    }

    static int GetTrailheadScore(int[,] map, int row, int col, int rows, int cols, bool[,] visited)
    {
        HashSet<(int, int)> reachableNines = new HashSet<(int, int)>();
        DFS(map, row, col, rows, cols, visited, 0, reachableNines);

        return reachableNines.Count;
    }

    static void DFS(int[,] map, int row, int col, int rows, int cols, bool[,] visited, int currentHeight, HashSet<(int, int)> reachableNines)
    {
        if (row < 0 || row >= rows || col < 0 || col >= cols || visited[row, col] || map[row, col] != currentHeight)
        {
            return;
        }

        visited[row, col] = true;

        if (map[row, col] == 9)
        {
            reachableNines.Add((row, col));
            return;
        }

        DFS(map, row - 1, col, rows, cols, visited, currentHeight + 1, reachableNines);
        DFS(map, row + 1, col, rows, cols, visited, currentHeight + 1, reachableNines);
        DFS(map, row, col - 1, rows, cols, visited, currentHeight + 1, reachableNines);
        DFS(map, row, col + 1, rows, cols, visited, currentHeight + 1, reachableNines);
    }
}
