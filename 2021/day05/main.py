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
    def __init__(self,x_min: int, y_min: int, x_max: int, y_max: int) -> None:
        self.x_min = x_min
        self.y_min = y_min
        self.x_max = x_max
        self.y_max = y_max
        self.grid = create_empty_2D_array(y_max, x_max)

    def draw_horizontal_and_vertical_lines(self, lines: list[Line]) -> None:
        for line in lines:
            # only draw horizontal or vertical lines
            if line.is_horizontal or line.is_vertical:
                self.draw_line(line)

    def draw_line(self, line: Line) -> None:
        pass

    def get_number_of_coordinates_with_overlap(self, threshold: int) -> int:
        number_overlap = 0
        for y_pos in range(len(self.grid)):
            for x_pos in range(len(self.grid[0])):
                if self.grid[y_pos][x_pos] >= threshold:
                    number_overlap += 1
        return number_overlap

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
def create_empty_2D_array(n_rows: int, n_columns: int) -> list:
    return [[0] * n_columns for _ in range(n_rows)]   

''' This Method takes an input line as a string ("x1,y1 -> x2,y2"), extract the Coordinates and give them back as a tuple. '''
def get_coordinates_of_line(input_line: str) -> tuple[Coordinate, Coordinate]:
    coordinates = input_line.replace(" ", "").split("->")
    begin_x_y = coordinates[0].split(",")
    begin_Coordinate = Coordinate(int(begin_x_y[0]),int(begin_x_y[1]))
    end_x_y = coordinates[1].split(",")
    end_Coordinate = Coordinate(int(end_x_y[0]),int(end_x_y[1]))

    return (begin_Coordinate, end_Coordinate)

''' Calculate the Coordinate system boarders from the Line list '''
def get_coordinate_system_boarders(lines: list[Line]) -> tuple[int, int, int, int]:
    x_min = 0
    y_min = 0
    x_max = 0
    y_max = 0
    # iterate over the Lines and their Coordinates to calculate the boarders of the CoordinateSystem
    for line in lines:
        if line.begin.x < x_min:
            x_min = line.begin.x
        if line.end.x < x_min:
            x_min = line.end.x
        #------------------------
        if line.begin.y < y_min:
            y_min = line.begin.y
        if line.end.y < y_min:
            y_min = line.end.y 
        #------------------------
        if line.begin.x > x_max:
            x_max = line.begin.x
        if line.end.x > x_max:
            x_max = line.end.x
        #------------------------
        if line.begin.y > y_max:
            y_max = line.begin.y
        if line.end.y > y_max:
            y_max = line.end.y 

    return (x_min, y_min, x_max, y_max)

# --------------- End of Helper Methods ---------------    

def part_one(input_path: str) -> int:
    lines = read_input(input_path)
    x_min, y_min, x_max, y_max = get_coordinate_system_boarders(lines)
    cs = CoordinateSystem(x_min, y_min, x_max, y_max)
    cs.draw_horizontal_and_vertical_lines(lines)
    result = cs.get_number_of_coordinates_with_overlap(2)
    return result

if __name__ == "__main__":
    input_path = "./2021/day05/input_test.txt"
    print("The result of Part 1:", part_one(input_path))