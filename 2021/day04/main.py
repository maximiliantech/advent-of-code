from dataclasses import dataclass

# ------------- Begin of BingCard ----------------
@dataclass
class BingoCard:
    def __init__(self, board: str) -> None:
        self.board = self.string_board_to_2D_array(board)
        self.check_board = self.create_empty_2D_array(self.board)
    
    def string_board_to_2D_array(self, board: str) -> list:
        board_2D = []
        array = board.split("\n")
        for row in array:
            board_2D.append(convert_array_from_string_to_int(remove_all_occ_of_el_from_arr(row.split(" "), "")))
        return board_2D

    def create_empty_2D_array(self,array: list) -> list:
        n_rows = len(array)
        n_columns = len(array[0])
        return [[0] * n_columns for _ in range(n_rows)] 

    def check_bingo(self) -> list:
        return []  

# ----------- End of BingoCard ------------------

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

def part_one(path: str) -> int:
    sequence, boards = read_input(path)
    return 0

if __name__ == "__main__": 
    input_path = "./2021/day04/input_test.txt"
    print("The result of Part 1:", part_one(input_path))