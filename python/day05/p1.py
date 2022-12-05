import re
from collections import deque

if __name__ == "__main__":
    with open("input.txt", "r") as f:
        stacks_init_state_raw, moves_raw = f.read().split("\n\n")
        stacks_init_state = stacks_init_state_raw.split("\n")

        # get stacks count from last line of init state
        N_STACKS = int(stacks_init_state.pop()[-1])
        stacks = tuple(deque() for _ in range(N_STACKS))

        # prepare initial stacks data
        SLICE_BY = 4  # -> [n]<space>
        for line in stacks_init_state:
            # go over every slice and check if theres a crate
            for i in range(0, len(line), SLICE_BY):
                slice = line[i : i + SLICE_BY]
                if "[" in slice:
                    stack_number = i // SLICE_BY
                    stacks[stack_number].appendleft(slice[1])

        # actual work
        for move in moves_raw.split("\n"):
            how_many, from_, to = map(int, re.findall(r"\d+", move))
            for _ in range(how_many):
                stacks[to - 1].append(stacks[from_ - 1].pop())

        print("".join((stack.pop() for stack in stacks)))
