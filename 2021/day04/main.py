from dataclasses import dataclass

# ----------------- Begin of BingCard -------------------
@dataclass
class BingoCard:
    def __init__(self, board: str) -> None:
        self.board = self.string_board_to_2D_array(board)
        self.n_rows = len(self.board)
        self.n_columns = len(self.board[0])
        self.check_board = create_empty_2D_array(self.n_rows, self.n_columns)
    
    def string_board_to_2D_array(self, board: str) -> list:
        board_2D = []
        array = board.split("\n")
        for row in array:
            board_2D.append(convert_array_from_string_to_int(remove_all_occ_of_el_from_arr(row.split(" "), "")))
        return board_2D

    def get_sum_of_unmarked_numbers() -> list:
        pass   

    def check_bingo(self) -> tuple[bool, list[int]]:
        return []

    def add_number(number: int) -> None:
        pass
# ----------------- End of BingoCard ---------------------

# ----------- Begin of General/Helper Methods --------------
def transform_array(array):
    return [list(number) for number in zip(*array)]

def create_empty_2D_array(n_rows, n_columns) -> list:
    return [[0] * n_columns for _ in range(n_rows)]     

def remove_all_occ_of_el_from_arr(array, element) -> list:
    return [x for x in array if x != element]

def convert_array_from_string_to_int(array) -> list:
    return [int(numeric_string) for numeric_string in array]     

''' Reads input .txt from path.
    Transfrom the first line of numbers to the sequence for Bingo Game
    The rest of the input are transformed to BingoCards '''
def read_input(path: str) -> tuple[list[int], list[BingoCard]]:
    with open(path) as f:
        file = f.read().split("\n\n")

    sequence = list(map(int, file[0].split(",")))

    # get the rest of the boards
    boards = []
    for board in file[1:]:
        boards.append(BingoCard(board))
    
    return sequence, boards

# ----------- End of General/Helper Methods --------------

def part_one(path: str) -> int:
    sequence, boards = read_input(path)
    for number in sequence:
        for board in boards:
            board.add_number(number)
            check, marked_numbers_array = board.check_bingo()
            if check:
                sum = board.get_sum_of_unmarked_numbers()
                return sum * number

    return 0

if __name__ == "__main__": 
    input_path = "./2021/day04/input_test.txt"
    print("The result of Part 1:", part_one(input_path))