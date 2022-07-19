import sys

# read input file and return the list of measurements
def read_measurements(path):
    with open(path) as file:
        return [int(line.rstrip()) for line in file]

# iterate over the list of measurements and find the number of times the measurement increases
def find_number_of_times_measurement_increases(measurements):
    number_of_times_measurement_increases = 0
    for i in range(len(measurements) - 1):
        if measurements[i] < measurements[i + 1]:
            number_of_times_measurement_increases += 1
    return number_of_times_measurement_increases

measurements = read_measurements('input.txt')
