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

report = read_diagnostic_report('./2021/day03/input.txt')
power_consumption = calculate_power_consumption(report)
print("The power consumption is: ", power_consumption)

# --------------- Part 2 ------------------
def get_most_common_indices(array):
    count_zero = 0
    count_one = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    # decide which bit is the most common
    common_value = 1 if count_one >= count_zero else 0

    return [i for i, x in enumerate(array) if x == common_value]

def get_least_common_indices(array):
    count_zero = 0
    count_one = 0

    for bit in array:
        if bit == 1:
            count_one += 1
        else:
            count_zero += 1
    
    # decide which bit is the least common
    least_value = 0 if count_zero <= count_one else 1 

    return [i for i, x in enumerate(array) if x == least_value]

def shortn_array(array, indices):
    arr = []
    for index in indices:
        arr.append(array[index])
    return arr

def bit_array_to_string(bit_array):
    bit_string = ""
    for bit in bit_array:
        bit_string += str(bit)
    return bit_string    

def calculate_oxygen_generator_rating(input):
    find_oxygen_rate_array = input

    for bit_position in range(len(input[0])):
        # -------- For the Oxygen Part ----------
        zip_oxygen_array = zip_binary_array(find_oxygen_rate_array)
        common_indices = get_most_common_indices(zip_oxygen_array[bit_position])
        find_oxygen_rate_array = shortn_array(find_oxygen_rate_array, common_indices)

        # break if there is only one bit array left
        if len(find_oxygen_rate_array) == 1:
            break
    
    # return the bits as string
    return bit_array_to_string(find_oxygen_rate_array[0])    

def calculate_co2_scrubber_rating(input):
    find_scrubber_rate_array = input

    for bit_position in range(len(input[0])):
        # -------- For the Scrubber Part -------------
        zip_srcubber_array = zip_binary_array(find_scrubber_rate_array)
        least_indices = get_least_common_indices(zip_srcubber_array[bit_position])
        find_scrubber_rate_array = shortn_array(find_scrubber_rate_array, least_indices)

        # break if there is only one bit array left
        if len(find_scrubber_rate_array) == 1:
            break
    
    # return the bits as string
    return bit_array_to_string(find_scrubber_rate_array[0]) 

def calculate_life_supporting_rate(input):
    oxygen_generator_rating = calculate_oxygen_generator_rating(input)
    co2_scrubber_rating = calculate_co2_scrubber_rating(input)
    return (int(oxygen_generator_rating, 2) * int(co2_scrubber_rating, 2))

life_supporting_rate = calculate_life_supporting_rate(report)
print("The life supporting rate is: ", life_supporting_rate)