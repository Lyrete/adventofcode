example = """
Time:      7  15   30
Distance:  9  40  200
"""

defmodule Day6Parser do
  def roots({b, c}) do
    {floor(b / 2 + (b ** 2 - 4 * c) ** 0.5 / 2), ceil(b / 2 - (b ** 2 - 4 * c) ** 0.5 / 2)}
  end

  def parse(input_string) do
    input_string
    |> String.replace(~r/[A-z:]+/, "")
    |> String.split("\n", trim: true)
    |> Enum.map(fn x ->
      [
        String.split(x, " ", trim: true) |> Enum.map(&String.to_integer/1),
        [String.replace(x, ~r/\s+/, "") |> String.to_integer()]
      ]
    end)
    |> Enum.zip()
    |> Enum.map(fn {x, y} -> Enum.zip(x, y) end)
    |> Enum.map(fn elems ->
      elems
      |> Enum.map(fn {u, w} -> Day6Parser.roots({u, w}) end)
      |> Enum.map(fn {a, b} -> a - b + 1 end)
      |> Enum.product()
    end)
  end
end

Day6Parser.parse(example)
|> IO.inspect()

File.read!("data/6.txt")
|> Day6Parser.parse()
|> IO.inspect()
