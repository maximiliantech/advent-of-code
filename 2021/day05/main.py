from dataclasses import dataclass

@dataclass
class Coordinate:
    pass

@dataclass
class CoordinateSystem:
    pass

def part_one(input_path: str) -> int:
    return 0

if __name__ == "__main__":
    input_path = "./2021/day05/input_test.txt"
    print("The result of Part 1:", part_one(input_path))