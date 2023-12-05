example = """
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""


def solve(s: str) -> int:
    lines = s.splitlines()
    cards = [1] * len(lines)
    total_wins = 0
    for line in lines:
        game_id, numbers = line.split(":", 1)
        game_id = int(game_id.replace("Card ", ""))

        winning_numbers, our_numbers = numbers.split(" | ", 1)
        winning_numbers = [int(x) for x in winning_numbers.split()]
        our_numbers = [int(x) for x in our_numbers.split()]

        wins = len(set(winning_numbers).intersection(set(our_numbers)))

        amount = cards[game_id - 1]

        for i in range(game_id, game_id + wins):
            cards[i] += 1 * amount

        if wins > 0:
            total_wins += 2 ** (wins - 1)

    return total_wins, sum(cards)


ex1, ex2 = solve(example.strip())
s1, s2 = solve(open("data/4.txt").read())

print(ex1)
print(ex2)
print(s1)
print(s2)
