# Part 1
# read input file and return the list of measurements
def read_diagnostic_report(path):
    with open(path) as file:
        return [line.rstrip() for line in file]