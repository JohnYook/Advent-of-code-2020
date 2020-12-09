import argparse

parser = argparse.ArgumentParser()

parser.add_argument("--input", "-i", type=str, required=True)
args = parser.parse_args()

target_value = 2020
all_nums = set()


def solve_day1_1():
    for num in all_nums:
        if target_value - num in all_nums:
            print(
                f"The numbers adding up to {target_value} are: {num}, {target_value - num}"
            )
            print(f"and their product is: {num * (target_value - num)}")
            break


def solve_day1_2():
    found = False
    for first in all_nums:
        for second in all_nums.difference({first}):
            if (target_value - first - second) in all_nums.difference({first, second}):
                print(
                    f"The numbers adding up to {target_value} are: {first}, {second}, and {target_value - first - second}"
                )
                print(
                    f"and their product is: {first * second * (target_value - first - second)}"
                )
                found = True
                break
        if found:
            break


with open(args.input) as data_file:
    data = [int(line) for line in data_file.read().splitlines()]
    for num in data:
        if not num in all_nums:
            all_nums.add(num)
    solve_day1_1()
    solve_day1_2()
