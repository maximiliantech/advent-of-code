# Part 1
# read input file and return the list of measurements
from math import gamma


def read_diagnostic_report(path):
    with open(path) as file:
        # returns a 2D-array of the input binary numbers
        return [[int(bit) for bit in line.rstrip()] for line in file]

def get_most_and_least_common_bits(array):
    count_zero = 0
    count_one = 0
    most_common_bit = 0
    least_common_bit = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    if count_zero > count_one:
        most_common_bit = 0
        least_common_bit = 1
    else:
        most_common_bit = 1
        least_common_bit = 0
    
    return most_common_bit, least_common_bit


def calculate_power_consumption(input):
    # zip the input array for better iteration
    zipped_input = [list(bit) for bit in zip(*input)]
    gamma = ''
    epsilon = ''
    for array in zipped_input:
        most_common_bit, least_common_bit = get_most_and_least_common_bits(array)
        gamma += str(most_common_bit)
        epsilon += str(least_common_bit)

    return (int(gamma, 2) * int(epsilon, 2))

report = read_diagnostic_report('./2021/day03/input.txt')
power_consumption = calculate_power_consumption(report)
print("The power consumption is: ", power_consumption)