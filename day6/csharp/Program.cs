class Program
{
    static void Main()
    {
        string[] lines = File.ReadAllLines("../input.txt");

        char[,] map = new char[lines.Length, lines[0].Length];
        (int x, int y) guardStartPosition = (-1, -1);
        char guardStartDirection = ' ';

        for (int i = 0; i < lines.Length; i++)
        {
            for (int j = 0; j < lines[i].Length; j++)
            {
                map[i, j] = lines[i][j];
                if ("^v<>".Contains(map[i, j]))
                {
                    guardStartPosition = (i, j);
                    guardStartDirection = map[i, j];
                    map[i, j] = '.';
                }
            }
        }

        int part1Result = Part1(map, guardStartPosition, guardStartDirection);
        int part2Result = Part2(map, guardStartPosition, guardStartDirection);

        Console.WriteLine($"Part 1: {part1Result}");

        // gpt solution i couldnt figure it out smh
        Console.WriteLine($"Part 2: {part2Result}");
    }

    static int Part1(char[,] map, (int x, int y) start, char startDirection)
    {
        var directions = new Dictionary<char, (int dx, int dy)>
        {
            { '^', (-1, 0) },
            { 'v', (1, 0) },
            { '<', (0, -1) },
            { '>', (0, 1) }
        };

        var turnRight = new Dictionary<char, char>
        {
            { '^', '>' },
            { '>', 'v' },
            { 'v', '<' },
            { '<', '^' }
        };

        var visited = new HashSet<(int x, int y)>();
        visited.Add(start);

        (int x, int y) guardPosition = start;
        char guardDirection = startDirection;

        int rows = map.GetLength(0);
        int cols = map.GetLength(1);

        while (true)
        {
            var (dx, dy) = directions[guardDirection];
            (int x, int y) nextPosition = (guardPosition.x + dx, guardPosition.y + dy);

            if (nextPosition.x < 0 || nextPosition.x >= rows || nextPosition.y < 0 || nextPosition.y >= cols)
            {
                break;
            }

            if (map[nextPosition.x, nextPosition.y] == '#')
            {
                guardDirection = turnRight[guardDirection];
            }
            else
            {
                guardPosition = nextPosition;
                visited.Add(guardPosition);
            }
        }

        return visited.Count;
    }

    // gpt solution i couldnt figure it out smh
    static int Part2(char[,] map, (int x, int y) start, char startDirection)
    {
        var directions = new Dictionary<char, (int dx, int dy)>
        {
            { '^', (-1, 0) },
            { 'v', (1, 0) },
            { '<', (0, -1) },
            { '>', (0, 1) }
        };

        var turnRight = new Dictionary<char, char>
        {
            { '^', '>' },
            { '>', 'v' },
            { 'v', '<' },
            { '<', '^' }
        };

        int loopObstructionCount = 0;

        for (int i = 0; i < map.GetLength(0); i++)
        {
            for (int j = 0; j < map.GetLength(1); j++)
            {
                if (map[i, j] != '.' || (i, j) == start)
                    continue;

                if (SimulateWithObstruction(map, start, startDirection, (i, j)))
                {
                    loopObstructionCount++;
                }
            }
        }

        bool SimulateWithObstruction(char[,] map, (int x, int y) start, char startDirection, (int x, int y) obstruction)
        {
            char[,] mapClone = (char[,])map.Clone();
            mapClone[obstruction.x, obstruction.y] = '#';

            (int x, int y) guardPosition = start;
            char guardDirection = startDirection;

            var visitedStates = new HashSet<((int x, int y) position, char direction)>();
            visitedStates.Add((guardPosition, guardDirection));

            int rows = mapClone.GetLength(0);
            int cols = mapClone.GetLength(1);

            while (true)
            {
                var (dx, dy) = directions[guardDirection];
                (int x, int y) nextPosition = (guardPosition.x + dx, guardPosition.y + dy);

                if (nextPosition.x < 0 || nextPosition.x >= rows || nextPosition.y < 0 || nextPosition.y >= cols)
                {
                    return false;
                }

                if (mapClone[nextPosition.x, nextPosition.y] == '#')
                {
                    guardDirection = turnRight[guardDirection];
                }
                else
                {
                    guardPosition = nextPosition;

                    var state = (guardPosition, guardDirection);
                    if (visitedStates.Contains(state))
                    {
                        return true;
                    }

                    visitedStates.Add(state);
                }
            }
        }

        return loopObstructionCount;
    }
}
