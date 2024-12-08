class Program
{
    static void Main()
    {
        string[] input = File.ReadAllLines("../input.txt");
        int rows = input.Length;
        int cols = input[0].Length;

        Dictionary<char, List<(int x, int y)>> antennas = new Dictionary<char, List<(int x, int y)>>();

        for (int y = 0; y < rows; y++)
        {
            for (int x = 0; x < cols; x++)
            {
                char cell = input[y][x];
                if (char.IsLetterOrDigit(cell))
                {
                    if (!antennas.ContainsKey(cell))
                    {
                        antennas[cell] = new List<(int x, int y)>();
                    }

                    antennas[cell].Add((x, y));
                }
            }
        }

        HashSet<(int x, int y)> antinodes = new HashSet<(int x, int y)>();

        foreach (var pair in antennas)
        {
            List<(int x, int y)> positions = pair.Value;

            foreach (var pos in positions)
            {
                antinodes.Add(pos);
            }

            for (int i = 0; i < positions.Count; i++)
            {
                for (int j = i + 1; j < positions.Count; j++)
                {
                    var (x1, y1) = positions[i];
                    var (x2, y2) = positions[j];

                    AddAntinodes(x1, y1, x2, y2, rows, cols, antinodes);
                }
            }
        }

        Console.WriteLine(antinodes.Count);
    }

    static void AddAntinodes(int x1, int y1, int x2, int y2, int rows, int cols, HashSet<(int x, int y)> antinodes)
    {
        if (x1 == x2)
        {
            int minY = Math.Min(y1, y2);
            int maxY = Math.Max(y1, y2);

            for (int y = minY; y <= maxY; y++)
            {
                if (IsWithinBounds(x1, y, rows, cols))
                {
                    antinodes.Add((x1, y));
                }
            }
        }
        else if (y1 == y2)
        {
            int minX = Math.Min(x1, x2);
            int maxX = Math.Max(x1, x2);

            for (int x = minX; x <= maxX; x++)
            {
                if (IsWithinBounds(x, y1, rows, cols))
                {
                    antinodes.Add((x, y1));
                }
            }
        }
        else
        {
            int dx = x2 - x1;
            int dy = y2 - y1;

            // traverse from (x1, y1) to (x2, y2)
            TraverseDiagonal(x1, y1, dx, dy, rows, cols, antinodes);

            // traverse from (x, y2) to (x1, y1)
            TraverseDiagonal(x2, y2, -dx, -dy, rows, cols, antinodes);
        }
    }

    static void TraverseDiagonal(int x, int y, int dx, int dy, int rows, int cols, HashSet<(int x, int y)> antinodes)
    {
        while (true)
        {
            if (IsWithinBounds(x, y, rows, cols))
            {
                antinodes.Add((x, y));
            }

            if (!IsWithinBounds(x, y, rows, cols))
            {
                break;
            }

            x += dx;
            y += dy;
        }
    }

    static bool IsWithinBounds(int x, int y, int rows, int cols)
    {
        return x >= 0 && y >= 0 && x < cols && y < rows;
    }
}
