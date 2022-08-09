from dataclasses import dataclass
from tkinter import Y
# ------------- Begin of Classes ----------------
@dataclass
class Coordinate:
    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y

@dataclass
class Line:
    def __init__(self, begin: Coordinate, end: Coordinate) -> None:
        self.begin = begin
        self.end = end

    def is_horizontal(self) -> bool:
        return True if self.begin.x == self.end.x else False

    def is_vertical(self) -> bool:
        return True if self.begin.y == self.end.y else False    

@dataclass
class CoordinateSystem:
    def __init__(self, x_max: int, y_max: int) -> None:
        self.grid = create_empty_2D_array(y_max, x_max)

    def draw_horizontal_and_vertical_lines(self, lines: list[Line]):
        pass    

# --------------- End of Classes ---------------

# -------------- Begin of Helper Methods --------------
''' Reads the .txt input file and converts it into a list of type Line '''
def read_input(path: str) -> list[Line]:
    with open(path) as f:
        # put the line into a string array
        input_line_list = f.read().split("\n")

    export_line_list = list()
    for input_line in input_line_list:
        # get the coordinates from an input line in the list
        begin, end = get_coordinates_of_line(input_line)
        
        # create a Line and append the line at the export_line_list
        export_line_list.append(Line(begin, end))

    return export_line_list

''' Prints a 2D array fancy '''
def print_2D_array(array: list) -> None:
    print('\n'.join([''.join(['{:4}'.format(item) for item in row]) for row in array]))
    print("\n")

''' Creates a zero-2D-array based on number of columns and rows. '''
def create_empty_2D_array(n_rows, n_columns) -> list:
    return [[0] * n_columns for _ in range(n_rows)]   

''' This Method takes an input line as a string ("x1,y1 -> x2,y2"), extract the Coordinates and give them back as a tuple. '''
def get_coordinates_of_line(input_line: str) -> tuple[Coordinate, Coordinate]:
    coordinates = input_line.replace(" ", "").split("->")
    begin_x_y = coordinates[0].split(",")
    begin_Coordinate = Coordinate(int(begin_x_y[0]),int(begin_x_y[1]))
    end_x_y = coordinates[1].split(",")
    end_Coordinate = Coordinate(int(end_x_y[0]),int(end_x_y[1]))

    return (begin_Coordinate, end_Coordinate)

# --------------- End of Helper Methods ---------------    

def part_one(input_path: str) -> int:
    lines = read_input(input_path)
    return 0

if __name__ == "__main__":
    input_path = "./2021/day05/input_test.txt"
    print("The result of Part 1:", part_one(input_path))