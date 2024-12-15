class Program
{
    static void Main()
    {
        string[] lines = File.ReadAllLines("../input.txt");
        int rows = lines.Length;
        int cols = lines[0].Length;

        char[,] garden = new char[rows, cols];
        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                garden[i, j] = lines[i][j];
            }
        }

        bool[,] visited = new bool[rows, cols];
        int totalPricePart2 = 0;

        int[] dRow = { -1, 1, 0, 0 };
        int[] dCol = { 0, 0, -1, 1 };


        (int area, int sides) BFS(int startRow, int startCol, char plantType)
        {
            int area = 0;
            int sides = 0;

            Queue<(int, int)> queue = new Queue<(int, int)>();
            queue.Enqueue((startRow, startCol));
            visited[startRow, startCol] = true;

            while (queue.Count > 0)
            {
                var (row, col) = queue.Dequeue();
                area++;

                for (int i = 0; i < 4; i++)
                {
                    int newRow = row + dRow[i];
                    int newCol = col + dCol[i];

                    if (newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || garden[newRow, newCol] != plantType)
                    {
                        sides++;
                    }
                    else if (!visited[newRow, newCol])
                    {
                        visited[newRow, newCol] = true;
                        queue.Enqueue((newRow, newCol));
                    }
                }
            }

            return (area, sides);
        }

        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                if (!visited[i, j])
                {
                    char plantType = garden[i, j];
                    var (area, sides) = BFS(i, j, plantType);

                    int pricePart2 = area * sides;
                    totalPricePart2 += pricePart2;
                }
            }
        }

        Console.WriteLine($"Total: {totalPricePart2}");
    }
}

