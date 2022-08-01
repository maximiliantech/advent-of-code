import sys

# Part 1
# read input file and return the list of measurements
def read_measurements(path):
    with open(path) as file:
        return [line.rstrip() for line in file]

def calculate_horizontal_and_depth(measurements):
    horizontal_count = 0
    depth_count = 0
    for measurement in measurements:
        instruction = measurement.split(' ')
        match instruction[0]:
            case 'forward':
                horizontal_count += int(instruction[1])
            case 'up':
                depth_count -= int(instruction[1])
            case 'down':
                depth_count += int(instruction[1])        

    return horizontal_count * depth_count

print("Ergebnis von Tag 2.1:", calculate_horizontal_and_depth(read_measurements('./2021/day02/input.txt')))