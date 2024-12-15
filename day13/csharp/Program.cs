class ClawContraption
{
    static void Main(string[] args)
    {
        string[] input = File.ReadAllLines("../input.txt");
        List<Machine> machines = ParseInput(input);

        long minTokens = 0;

        foreach (var machine in machines)
        {
            machine.XPrize += 10000000000000;
            machine.YPrize += 10000000000000;
        }

        foreach (var machine in machines)
        {
            var minimumTokens = SolveMachine(machine);

            if (minimumTokens.HasValue)
            {
                minTokens += minimumTokens.Value;
            }
        }

        Console.WriteLine($"tokens: {minTokens}");
    }

    static List<Machine> ParseInput(string[] input)
    {
        var machines = new List<Machine>();

        for (int i = 0; i < input.Length; i += 4)
        {
            var xStepA = int.Parse(input[i].Split(' ')[2].TrimEnd(',').Split('+')[1]);
            var yStepA = int.Parse(input[i].Split(' ')[3].Split('+')[1]);

            var xStepB = int.Parse(input[i + 1].Split(' ')[2].TrimEnd(',').Split('+')[1]);
            var yStepB = int.Parse(input[i + 1].Split(' ')[3].Split('+')[1]);

            var prize = input[i + 2].Split(' ');
            var xPrize = int.Parse(prize[1].TrimEnd(',').Split('=')[1]);
            var yPrize = int.Parse(prize[2].Split('=')[1]);

            machines.Add(new Machine
            {
                XStepA = xStepA,
                YStepA = yStepA,
                XStepB = xStepB,
                YStepB = yStepB,
                XPrize = xPrize,
                YPrize = yPrize
            });
        }

        return machines;

    }

    static long? SolveMachine(Machine machine)
    {
        long determinant = machine.XStepA * machine.YStepB - machine.YStepA * machine.XStepB;

        if (determinant == 0)
        {
            return null;
        }

        long determinantA = machine.XPrize * machine.YStepB - machine.YPrize * machine.XStepB;
        long determinantB = machine.XStepA * machine.YPrize - machine.YStepA * machine.XPrize;

        if (determinantA % determinant != 0 || determinantB % determinant != 0)
        {
            return null;
        }

        long buttonAPresses = determinantA / determinant;
        long buttonBPresses = determinantB / determinant;

        if (buttonAPresses < 0 || buttonBPresses < 0)
        {
            return null;
        }

        return buttonAPresses * 3 + buttonBPresses * 1;
    }

    public class Machine
    {
        public int XStepA { get; set; }
        public int YStepA { get; set; }
        public int XStepB { get; set; }
        public int YStepB { get; set; }
        public long XPrize { get; set; }
        public long YPrize { get; set; }
    }

}
