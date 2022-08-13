# -------------- Begin of Helper Methods --------------
''' Reads the .txt input file and converts it into a list of type Line '''
def read_input(path: str) -> list[int]:
    with open(path) as f:
        # put the single line into an int array
        return list(map(int, f.read().split(",")))

# -------------- End of Helper Methods --------------

def part_one(input_path: str) -> int:
    numbers = read_input(input_path)
    print("Initial state: ", numbers)
    for day in range(1,81):
        for i in range(len(numbers)):
            lanternfish = numbers[i]
            numbers[i] -= 1
            if lanternfish == 0:
                numbers[i] = 6
                numbers.append(8)
        #print("After day: ", day, numbers)
    return len(numbers)

if __name__ == "__main__":
    input_path = "./2021/day06/input.txt"
    print("The result of Part 1:", part_one(input_path))