def get_value(board: list[list[str]]) -> int:
    value = 0
    for column in board:
        value += sum(i * (c == "O") for i, c in enumerate(column[::-1], 1))

    return value


def tilt_north(board: list[list[str]]):
    tilted_board = []
    for column in board:
        new_col = []
        for part in column.split("#"):
            if part:
                new_col.append("".join(["O"] * part.count("O") +
                                       ["."] * part.count(".")))
            else:
                new_col.append("")
        tilted_board.append("#".join(new_col))

    return tilted_board


def rotate(board: list[list[str]]):
    return ["".join(line) for line in zip(*map(reversed, board))]


def do_cycle(board: list[list[str]]):
    # Run the move + rotate four times
    for _ in range(4):
        board = rotate(tilt_north(board))

    return board


def parse(s: str):
    cols = ["".join(char) for char in zip(*s.strip().splitlines())]

    p1 = get_value(tilt_north(cols))

    limit = 1000000000
    t = 0
    cycled_states = [cols]

    while True:
        end = do_cycle(cycled_states[t])
        if end in cycled_states:
            i = cycled_states.index(end)
            cycle_length = t + 1 - i
            break
        cycled_states.append(end)
        t += 1

    p2 = get_value(cycled_states[(limit - i) % cycle_length + i])

    return p1, p2


def solve(s: str) -> tuple[int, int]:
    s1, s2 = parse(s)

    return s1, s2


input_data = open("data/14.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
