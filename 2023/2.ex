example = """
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
"""

defmodule Day2Parser do
  @spec matchColors([String.t()]) :: [String.t()]
  def matchColors(content) do
    content
    |> String.split("; ")
    |> Enum.flat_map(fn x -> String.split(x, ", ") end)
    |> Enum.map(fn x -> String.split(x, " ") end)
    |> Enum.group_by(fn x -> Enum.at(x, 1) end, fn x ->
      Enum.at(x, 0) |> Integer.parse() |> elem(0)
    end)
    |> Enum.map(fn {_color, y} ->
      Enum.max(y)
    end)
  end
end

sample =
  example
  |> String.trim()
  |> String.split("\n")
  |> Enum.map(fn x -> String.split(x, ": ") end)
  |> Enum.map(fn x ->
    {Enum.at(x, 0) |> String.split(" ") |> List.last(), Day2Parser.matchColors(Enum.at(x, 1))}
  end)

parsed =
  "data/2.txt"
  |> File.stream!()
  |> Stream.map(&String.trim/1)
  |> Enum.map(fn x -> String.split(x, ": ") end)
  |> Enum.map(fn x ->
    {Enum.at(x, 0) |> String.split(" ") |> List.last(), Day2Parser.matchColors(Enum.at(x, 1))}
  end)

IO.puts("Example, should be 8:")

sample
|> Enum.filter(fn {_id, [b, g, r]} -> b <= 12 && g <= 13 && r <= 14 end)
|> Enum.map(fn {id, _colors} -> elem(Integer.parse(id), 0) end)
|> Enum.sum()
|> IO.inspect()

IO.puts("Part 1:")

parsed
|> Enum.filter(fn {_id, [b, g, r]} -> b <= 14 && g <= 13 && r <= 12 end)
|> Enum.map(fn {id, _colors} -> elem(Integer.parse(id), 0) end)
|> Enum.sum()
|> IO.inspect()

IO.puts("Example for part 2, should be 2286:")

sample
|> Enum.map(fn {_id, colors} -> Enum.product(colors) end)
|> Enum.sum()
|> IO.inspect()

IO.puts("Part 2:")

parsed
|> Enum.map(fn {_id, colors} -> Enum.product(colors) end)
|> Enum.sum()
|> IO.inspect()
