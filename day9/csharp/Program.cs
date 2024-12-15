class Program
{
    static void Main()
    {
        string lines = File.ReadAllText("../input.txt").Trim();

        List<char> blocks = parseDisk(lines);

        MoveBlocks(blocks);

        long checksum = calculateChecksum(blocks);

        Console.WriteLine($"Checksum: {checksum}");
    }

    static List<char> parseDisk(string lines)
    {
        var blocks = new List<char>();
        int fileId = 0;

        for (int i = 0; i < lines.Length; i += 2)
        {
            int fileLength = lines[i] - '0';
            int freeSpaceLength = (i + 1 < lines.Length) ? lines[i + 1] - '0' : 0;

            for (int j = 0; j < fileLength; j++)
            {
                blocks.Add((char)('0' + fileId));
            }

            for (int j = 0; j < freeSpaceLength; j++)
            {
                blocks.Add('.');
            }

            fileId++;
        }

        return blocks;
    }

    static void MoveBlocks(List<char> blocks)
    {
        int length = blocks.Count;
        int maxFileId = blocks.Max(b => b != '.' ? b - '0' : -1);

        for (int fileId = maxFileId; fileId > 0; fileId--)
        {
            int fileStartIndex = blocks.FindIndex(b => b == (char)('0' + fileId));
            if (fileStartIndex == -1)
            {
                continue;
            }

            int fileLength = blocks.Skip(fileStartIndex).TakeWhile(b => b == (char)('0' + fileId)).Count();

            int freeSpaceStartIndex = -1;

            for (int i = 0; i <= length - fileLength; i++)
            {
                if (blocks.Skip(i).Take(fileLength).All(b => b == '.'))
                {
                    freeSpaceStartIndex = i;

                    if (fileStartIndex < freeSpaceStartIndex)
                    {
                        freeSpaceStartIndex = -1;
                        break;
                    }

                    break;
                }
            }

            if (freeSpaceStartIndex != -1)
            {
                for (int i = 0; i < fileLength; i++)
                {
                    blocks[freeSpaceStartIndex + i] = (char)('0' + fileId);
                    blocks[fileStartIndex + i] = '.';
                }
            }
        }
    }

    static long calculateChecksum(List<char> blocks)
    {
        long checksum = 0;

        for (int i = 0; i < blocks.Count; i++)
        {
            if (blocks[i] != '.')
            {
                int fileId = blocks[i] - '0';
                checksum += i * fileId;
            }
        }

        return checksum;
    }
}
