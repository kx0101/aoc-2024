namespace Aoc4
{
    class Program
    {
        static int[] dx = { 1, 1, 0, -1, -1, -1, 0, 1 };
        static int[] dy = { 0, 1, 1, 1, 0, -1, -1, -1 };

        static void Main()
        {
            char[,] grid = ReadGridFromFile("../input.txt");

            int count = CountXMASOccurrences(grid);
            int count2 = CountXMASPatterns(grid);

            Console.WriteLine($"part 1: {count}");
            Console.WriteLine($"part 2: {count2}");
        }

        static char[,] ReadGridFromFile(string filePath)
        {
            string[] lines = File.ReadAllLines(filePath);

            int rows = lines.Length;
            int cols = lines[0].Length;
            char[,] grid = new char[rows, cols];

            for (int i = 0; i < rows; i++)
            {
                for (int j = 0; j < cols; j++)
                {
                    grid[i, j] = lines[i][j];
                }
            }

            return grid;
        }

        static int CountXMASOccurrences(char[,] grid)
        {
            int rows = grid.GetLength(0);
            int cols = grid.GetLength(1);
            string needle = "XMAS";
            int count = 0;

            for (int i = 0; i < rows; i++)
            {
                for (int j = 0; j < cols; j++)
                {
                    for (int direction = 0; direction < 8; direction++)
                    {
                        if (CheckWord(grid, i, j, direction, needle))
                        {
                            count++;
                        }
                    }
                }
            }
            return count;
        }

        static bool CheckWord(char[,] grid, int row, int col, int direction, string needle)
        {
            int rows = grid.GetLength(0);
            int cols = grid.GetLength(1);

            for (int k = 0; k < needle.Length; k++)
            {
                int newRow = row + k * dx[direction];
                int newCol = col + k * dy[direction];

                if (newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow, newCol] != needle[k])
                {
                    return false;
                }
            }

            return true;
        }

        static int CountXMASPatterns(char[,] grid)
        {
            int rows = grid.GetLength(0);
            int cols = grid.GetLength(1);
            int count = 0;

            for (int i = 1; i < rows - 1; i++)
            {
                for (int j = 1; j < cols - 1; j++)
                {
                    if (CheckBothDiagonals(grid, i, j))
                    {
                        count++;
                    }
                }
            }

            return count;
        }

        static bool CheckBothDiagonals(char[,] grid, int i, int j)
        {
            // first diagonal \
            // top left, middle, bottom right
            bool diag1MAS = IsMASSequence(grid[i - 1, j - 1], grid[i, j], grid[i + 1, j + 1]);
            bool diag1SAM = IsMASSequence(grid[i + 1, j + 1], grid[i, j], grid[i - 1, j - 1]);

            // second diagonal /
            // top right, middle, bottom left
            bool diag2MAS = IsMASSequence(grid[i - 1, j + 1], grid[i, j], grid[i + 1, j - 1]);
            bool diag2SAM = IsMASSequence(grid[i + 1, j - 1], grid[i, j], grid[i - 1, j + 1]);

            // Both diagonals must spell MAS
            return (diag1MAS || diag1SAM) && (diag2MAS || diag2SAM);
        }

        static bool IsMASSequence(char m, char a, char s)
        {
            return (m == 'M' && a == 'A' && s == 'S');
        }
    }
}
