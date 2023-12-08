example = """
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""

defmodule Day6Parser do
  def count_cards(hand) do
    hand
    |> String.split("", trim: true)
    |> Enum.group_by(& &1)
    |> Enum.map(fn {k, v} -> {k, Enum.count(v)} end)
    |> Enum.sort(fn {_, x}, {_, y} -> y < x end)
  end

  def rank_hand(hand) do
    case(count_cards(hand)) do
      [{_, 5}] -> 1
      [{_, 4}, {_, 1}] -> 2
      [{_, 3}, {_, 2}] -> 3
      [{_, 3}, {_, 1}, {_, 1}] -> 4
      [{_, 2}, {_, 2}, {_, 1}] -> 5
      [{_, 2}, {_, 1}, {_, 1}, {_, 1}] -> 6
      _ -> 7
    end
  end

  def replace_jokers(hand) do
    filtered =
      hand
      |> count_cards()
      |> Enum.filter(fn {k, _} -> k != "J" end)

    # if there are no non-jokers, jokers is the only thing in it

    if length(filtered) == 0 do
      hand
    else
      most_common = Enum.at(filtered, 0) |> elem(0)
      String.replace(hand, "J", most_common)
    end
  end

  def rank_card(card) do
    case(card) do
      "T" -> 10
      "J" -> 1
      "R" -> 11
      "Q" -> 12
      "K" -> 13
      "A" -> 14
      _ -> String.to_integer(card)
    end
  end

  def convert_to_int(hand) do
    hand
    |> String.split("", trim: true)
    |> Enum.map(fn x -> rank_card(x) end)
  end

  def compare_hands(hand1, hand2) do
    score1 = rank_hand(hand1 |> replace_jokers())
    score2 = rank_hand(hand2 |> replace_jokers())

    if(score1 == score2) do
      convert_to_int(hand1) < convert_to_int(hand2)
    else
      score1 > score2
    end
  end

  def parse(input_string) do
    input_string
    |> String.split("\n", trim: true)
    |> Enum.map(fn x -> x |> String.replace("J", "R") |> String.split(" ", trim: true) end)
    |> Enum.map(fn [x, y] -> {x, y |> String.to_integer()} end)
    |> Enum.sort(fn {cards, _}, {cards2, _} -> compare_hands(cards, cards2) end)
    |> Enum.map(fn {_, bid} -> bid end)
    |> Enum.with_index()
    |> Enum.reduce(0, fn {bid, index}, acc -> acc + bid * (index + 1) end)
  end

  def parse2(input_string) do
    input_string
    |> String.split("\n", trim: true)
    |> Enum.map(fn x ->
      String.split(x, " ", trim: true)
    end)
    |> Enum.map(fn [x, y] -> {x, y |> String.to_integer()} end)
    |> Enum.sort(fn {cards, _}, {cards2, _} -> compare_hands(cards, cards2) end)
    |> Enum.map(fn {_, bid} -> bid end)
    |> Enum.with_index()
    |> Enum.reduce(0, fn {bid, index}, acc -> acc + bid * (index + 1) end)
  end
end

Day6Parser.parse(example)
|> IO.inspect()

File.read!("data/7.txt")
|> Day6Parser.parse()
|> IO.inspect()

Day6Parser.parse2(example)
|> IO.inspect()

File.read!("data/7.txt")
|> Day6Parser.parse2()
|> IO.inspect()
