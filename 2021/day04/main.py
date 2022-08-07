from dataclasses import dataclass

@dataclass
class BingoCard:
    pass

def read_input(path: str) -> tuple[list[int], list[BingoCard]]:
    with open(path) as f:
        file = f.read().split("\n\n")

    sequence = list(map(int, file[0].split(",")))

    boards = []
    return sequence, boards   

def part_one(path: str) -> int:
    sequence, borads = read_input(path)
    return 0

if __name__ == "__main__": 
    input_path = "./2021/day04/input_test.txt"
    print("The result of Part 1:", part_one(input_path))