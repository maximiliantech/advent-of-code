import sys

# -------- Part 1 ----------

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

measurements = read_measurements('./2021/day01/input.txt')
print("Number of times the measurement increases: " + str(find_number_of_times_measurement_increases(measurements)))

# -------- Part 2 ----------
def find_number_of_times_measurement_increases_in_three_window(measurements):
    number_of_increases = 0
    for i in range(len(measurements) - 3):
        if measurements[i] + measurements[i+1] + measurements[i+2] < measurements[i+1] + measurements[i+2] + measurements[i+3]:
            number_of_increases += 1
    return number_of_increases

print("Number of times the measurement increases in a three window: " + str(find_number_of_times_measurement_increases_in_three_window(measurements)))