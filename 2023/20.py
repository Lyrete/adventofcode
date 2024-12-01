from math import lcm


example = """
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
"""

BROADCAST = "broadcaster"


def parse(s: str):
    nodes = {}

    for line in s.strip().splitlines():
        l, r = line.split(" -> ")
        if l == BROADCAST:
            nodes[BROADCAST] = (BROADCAST, r.split(", "), False, {})
        else:
            name = l[1:]
            nodes[name] = (l[0], r.split(", "), False, {})

    for key, val in nodes.items():
        if val[0] != "&":
            continue

        t, out, on, i = val
        i = {k: False for k, (_, o, _, _) in nodes.items() if key in o}
        nodes[key] = (t, out, on, i)

    return nodes


def solve(s: str) -> tuple[int, int]:
    nodes = parse(s)

    p2target = "rx"
    in_nodes = [e for e, (_, o, _, i) in nodes.items() if p2target in o]
    target_flipflops = {}
    while len(in_nodes) > 0:
        name = in_nodes.pop()
        t, _, _, i = nodes[name]
        if t == "%":
            target_flipflops[name] = 0
            continue
        for e in i.keys():
            in_nodes.append(e)

    print(target_flipflops)

    button_presses = 0
    amount = [0, 0]
    n = 1000
    for i in range(n):
        button_presses += 1
        targets = [(BROADCAST, e, False) for e in nodes[BROADCAST][1]]
        while len(targets) > 0:
            new_targets = []
            for i, t, level in targets:
                if level:
                    amount[1] += 1
                else:
                    amount[0] += 1

                # if sending a pulse to nonexistent node
                if t not in nodes:
                    continue

                mod, output, on, inputs = nodes[t]
                if mod == "%" and level:
                    continue
                elif mod == "%":
                    on = not on
                    nodes[t] = mod, output, on, inputs

                    if on and t in target_flipflops:
                        target_flipflops[t] = button_presses

                    out = [(t, e, on) for e in output]
                elif mod == "&":
                    inputs[i] = level
                    pulse = not all(inputs.values())
                    out = [(t, e, pulse) for e in output]
                    nodes[t] = mod, output, on, inputs
                new_targets.extend(out)
            # print(new_targets)
            targets = new_targets

    amount[0] += button_presses
    s1 = amount[0] * amount[1]

    while not all(target_flipflops.values()):
        button_presses += 1
        targets = [(BROADCAST, e, False) for e in nodes[BROADCAST][1]]
        while len(targets) > 0:
            new_targets = []
            for i, t, level in targets:
                if level:
                    amount[1] += 1
                else:
                    amount[0] += 1

                # if sending a pulse to nonexistent node
                if t not in nodes:
                    continue

                mod, output, on, inputs = nodes[t]
                if mod == "%" and level:
                    continue
                elif mod == "%":
                    on = not on
                    nodes[t] = mod, output, on, inputs
                    out = [(t, e, on) for e in output]
                elif mod == "&":
                    inputs[i] = level
                    pulse = not all(inputs.values())
                    if not pulse:
                        for flip in inputs.keys():
                            target_flipflops[flip] = button_presses
                    out = [(t, e, pulse) for e in output]
                    nodes[t] = mod, output, on, inputs
                new_targets.extend(out)
            # print(new_targets)
            targets = new_targets

    return s1, lcm(*target_flipflops.values())


input_data = open("data/20.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
