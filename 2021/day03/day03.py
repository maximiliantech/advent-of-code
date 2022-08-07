# -------------- Part 1 -----------------
''' read input file and return a 2D-array of the input binary numbers '''
def read_diagnostic_report(path):
    with open(path) as file:
        # returns a 2D-array of the input binary numbers
        return [[int(bit) for bit in line.rstrip()] for line in file]

''' Calculate the most and least common bits by counting the number of 1s and 0s in the array.'''
def get_most_and_least_common_bits(array):
    count_zero = 0
    count_one = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    # decide which bit is the most common and least common and return the result
    return (0,1) if count_zero > count_one else (1,0)

def zip_binary_array(array):
    return [list(bit) for bit in zip(*array)]

def calculate_power_consumption(input):
    # zip the original input array to iterate over the columns
    zipped_input = zip_binary_array(input)
    gamma = ''
    epsilon = ''
    for array in zipped_input:
        # find the most and least common bits
        most_common_bit, least_common_bit = get_most_and_least_common_bits(array)
        gamma += str(most_common_bit)
        epsilon += str(least_common_bit)
    # mulitply the two binary numbers to get the integer power consumption
    return (int(gamma, 2) * int(epsilon, 2))

report = read_diagnostic_report('./2021/day03/input_test.txt')
power_consumption = calculate_power_consumption(report)
print("The power consumption is: ", power_consumption)

# --------------- Part 2 ------------------
def get_most_common_value(array):
    count_zero = 0
    count_one = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    # decide which bit is the most common and least common and return the result
    return (1,0) if count_one >= count_zero else (0,1)

def get_least_common_value(array):
    count_zero = 0
    count_one = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    # decide which bit is the most common and least common and return the result
    return (0,1) if count_zero >= count_one else (1,0)

def calculate_oxygen_generator_rating(input):
    zipped_input = zip_binary_array(input)
    help_array = []
    for array in zipped_input-1:
        most_common_value = get_most_common_value()

def calculate_co2_scrubber_rating(input):
    pass

def calculate_life_supporting_rate(input):
    oxygen_generator_rating = calculate_oxygen_generator_rating(input)
    co2_scrubber_rating = calculate_co2_scrubber_rating(input)

    return (int(oxygen_generator_rating, 2) * int(co2_scrubber_rating, 2))

life_supporting_rate = calculate_life_supporting_rate(report)
print("The life supporting rate is: ", life_supporting_rate)