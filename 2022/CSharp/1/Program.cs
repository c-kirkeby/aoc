
// string? line;
// List<int> elves = new List<int>();
// List<int> calories = new List<int>();
// // int index = 0;
// while ((line = Console.ReadLine()) is not null) {
//   if (line == "") {
//     calories.Add(null)
//   }
//   calories.Add(int.Parse(line));
// };
//
string? line;

// List<List<int>> elves = new List<List<int>>();
List<int> elves = new List<int>();

int elf = 0;

while ((line = Console.ReadLine()) is not null) {
  if (line != "") {
    elves[elf] = elves[elf] + int.Parse(line);
  } else {
    elves.Add(0);
    elf++;
  }
}

// foreach (int elf in elves) {
//   Console.WriteLine(elf.ToString());
// }
