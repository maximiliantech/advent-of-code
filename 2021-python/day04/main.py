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

    def get_sum_of_unmarked_numbers(self) -> int:
        result = 0
        for row_number in range(self.n_rows):
            for column_number in range(self.n_columns):
                if self.check_board[row_number][column_number] == 0:
                    result += self.board[row_number][column_number]
        return result
    
    def check_bingo_in_rows(self) -> tuple[bool, list[int]]:
        for row_number in range(self.n_rows):
            if self.check_board[row_number].count(1) == self.n_columns:
                return (True, self.check_board[row_number])
        return (False, [])

    def check_bingo_in_columns(self) -> tuple[bool, list[int]]:
        t_board = transform_array(self.board)
        t_check_board = transform_array(self.check_board)
        for column_number in range(self.n_columns):
            if t_check_board[column_number].count(1) == self.n_rows:
                return (True, t_board[column_number])
        return (False, [])

    def check_bingo(self) -> tuple[bool, list[int]]:
        rows_check, row_array = self.check_bingo_in_rows()
        column_check, column_array = self.check_bingo_in_columns()
        if rows_check:
            return (rows_check, row_array)
        if column_check:
            return (column_check, column_array)    
        return (False, [])

    def add_number(self, number: int) -> None:
        for row_number in range(self.n_rows):
            for column_number in range(self.n_columns):
                if self.board[row_number][column_number] == number:
                    self.check_board[row_number][column_number] = 1
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

def print_array(array: list) -> None:
    print('\n'.join([''.join(['{:4}'.format(item) for item in row]) for row in array]))
    print("\n")

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

''' Last board has won if winning_boards has only 1s in the array'''
def last_board_wins(winning_boards: list) -> bool:
    if winning_boards.count(1) == len(winning_boards):
        return True
    return False   

# ----------- End of General/Helper Methods --------------

def part_one(path: str) -> int:
    sequence, boards = read_input(path)
    for number in sequence:
        for board in boards:
            board.add_number(number)
            check, marked_numbers_array = board.check_bingo()
            if check:
                sum = board.get_sum_of_unmarked_numbers()
                return (sum * number)

    return 0

def part_two(path: str) -> int:
    sequence, boards = read_input(path)
    # create winning_boards array to mark which BingoCards have won
    winning_boards = [0 for _ in range(len(boards))]
    
    for number in sequence:
        for board_number in range(len(boards)):
            board = boards[board_number]
            board.add_number(number)
            check, marked_numbers_array = board.check_bingo()
            # if this board has a Bingo then mark this with one
            if check:
                winning_boards[board_number] = 1
            # if last board has won then calculate the result of the last winning board
            if last_board_wins(winning_boards):
                sum = board.get_sum_of_unmarked_numbers()
                return (sum * number)

    return 0

if __name__ == "__main__": 
    input_path = "./2021-python/day04/input.txt"
    print("The result of Part 1:", part_one(input_path))
    print("The result of Part 2:", part_two(input_path))