class ProgramPart2
{
    static void Main()
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

        int totalRating = 0;

        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                if (map[i, j] == 0)
                {
                    var memo = new Dictionary<(int, int), int>();
                    totalRating += GetTrailheadRating(map, i, j, rows, cols, 0, memo);
                }
            }
        }

        Console.WriteLine("total: " + totalRating);
    }

    static int GetTrailheadRating(int[,] map, int row, int col, int rows, int cols, int currHeight, Dictionary<(int, int), int> memo)
    {
        if (row < 0 || row >= rows || col < 0 || col >= cols || map[row, col] > 9 || map[row, col] != currHeight)
        {
            return 0;
        }

        if (map[row, col] == 9)
        {
            return 1;
        }

        if (memo.ContainsKey((row, col)))
        {
            return memo[(row, col)];
        }

        int totalTrails = 0;

        totalTrails += GetTrailheadRating(map, row - 1, col, rows, cols, currHeight + 1, memo);
        totalTrails += GetTrailheadRating(map, row, col - 1, rows, cols, currHeight + 1, memo);
        totalTrails += GetTrailheadRating(map, row + 1, col, rows, cols, currHeight + 1, memo);
        totalTrails += GetTrailheadRating(map, row, col + 1, rows, cols, currHeight + 1, memo);

        memo[(row, col)] = totalTrails;

        return totalTrails;
    }
}
