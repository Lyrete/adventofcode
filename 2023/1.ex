test = """
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
"""

test2 = """
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
"""

defmodule Day1Parser do
  def convert(string) do
    string
    |> String.graphemes()
    |> Enum.drop_while(fn x ->
      !Enum.member?(["1", "2", "3", "4", "5", "6", "7", "8", "9"], x)
    end)
    |> Enum.reverse()
    |> Enum.drop_while(fn x ->
      !Enum.member?(["1", "2", "3", "4", "5", "6", "7", "8", "9"], x)
    end)
    |> (fn filtered -> Enum.at(filtered, length(filtered) - 1) <> Enum.at(filtered, 0) end).()
    |> Integer.parse()
    |> elem(0)
  end

  def convertPart2(string) do
    written_numbers = %{
      "one" => "1",
      "two" => "2",
      "three" => "3",
      "four" => "4",
      "five" => "5",
      "six" => "6",
      "seven" => "7",
      "eight" => "8",
      "nine" => "9"
    }

    output = [
      string
      |> String.replace(Map.keys(written_numbers), fn number -> written_numbers[number] end,
        global: false
      )
      |> String.graphemes()
      |> Enum.drop_while(fn x ->
        !Enum.member?(["1", "2", "3", "4", "5", "6", "7", "8", "9"], x)
      end)
      |> List.first(),
      string
      |> String.reverse()
      |> String.replace(
        Map.keys(written_numbers) |> Enum.map(fn x -> String.reverse(x) end),
        fn number -> written_numbers[number |> String.reverse()] end,
        global: false
      )
      |> String.graphemes()
      |> Enum.drop_while(fn x ->
        !Enum.member?(["1", "2", "3", "4", "5", "6", "7", "8", "9"], x)
      end)
      |> List.first()
    ]

    output
  end
end

IO.puts("Example, should be 142:")

String.split(test, "\n")
|> Enum.filter(fn x -> x != "" end)
|> Enum.map(fn x -> Day1Parser.convert(x) end)
|> Enum.sum()
|> IO.puts()

IO.puts("Part 1:")

"data/1.txt"
|> File.stream!()
|> Stream.map(&String.trim/1)
|> Enum.map(fn x -> Day1Parser.convert(x) end)
|> Enum.sum()
|> IO.puts()

IO.puts("Example 2, should be 281:")

String.split(test2, "\n")
|> Enum.filter(fn x -> x != "" end)
|> Enum.map(fn x -> Day1Parser.convertPart2(x) end)
|> Enum.map(fn x -> String.to_integer(Enum.join(x)) end)
|> Enum.sum()
|> IO.puts()

IO.puts("Part 2:")

"data/1.txt"
|> File.stream!()
|> Stream.map(&String.trim/1)
|> Enum.map(fn x -> Day1Parser.convertPart2(x) end)
|> Enum.map(fn x -> String.to_integer(Enum.join(x)) end)
|> Enum.sum()
|> IO.puts()
