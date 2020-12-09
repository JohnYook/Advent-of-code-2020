import argparse

parser = argparse.ArgumentParser()

parser.add_argument("--input", "-i", type=str, required=True)
args = parser.parse_args()

target_value = 2020


def solve_day1(data):
    all_nums = set()
    for num in data:
        if target_value - num in all_nums:
            print(
                f"The numbers adding up to {target_value} are: {num}, {target_value - num}"
            )
            print(f"and their product is: {num * (target_value - num)}")
            break
        elif not num in all_nums:
            all_nums.add(num)


with open(args.input) as data_file:
    data = [int(line) for line in data_file.read().splitlines()]
    solve_day1(data)
