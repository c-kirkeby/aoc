int GetCalibrationValue(string line)
{

    List<char> numbers = line.Where(char.IsNumber).ToList();

    return int.Parse(numbers.First().ToString() + numbers.Last());
}

try
{
    StreamReader streamReader = new StreamReader("calibration.txt");
    string? line = streamReader.ReadLine();
    int sum = 0;

    while (line != null)
    {
        sum += GetCalibrationValue(line);
        line = streamReader.ReadLine();
    }

    Console.WriteLine($"Calibration value: {sum}");
}
catch (Exception exception)
{
    Console.WriteLine($"Error reading file: {exception.Message}");
}
