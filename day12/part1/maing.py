# this solves part 1 and 2

import sys


def split_plan(plan, mid_len):
    res = []
    for i in range(0, len(plan)):
        left = plan[:i]
        right = plan[i:]
        if len(right) < mid_len:
            break
        mid = right[:mid_len]
        right = right[mid_len:]
        res += [(left, mid, right)]
    print('split_plan', plan, mid_len, res)
    return res


def search(plan, nums):
    print('search', plan, nums)
    global cache
    if (plan, tuple(nums)) in cache:
        return cache[(plan, tuple(nums))]
    if len(nums) == 0:
        if '#' in plan:
            return 0
        return 1
    elif plan == '':
        return 0
    mid_num_pos = len(nums)//2
    left_nums = nums[:mid_num_pos]
    mid_num = nums[mid_num_pos]
    right_nums = nums[mid_num_pos+1:]

    print('nums', left_nums, mid_num, right_nums)

    res = 0
    for left_str, mid_str, right_str in split_plan(plan, mid_num):
        if '.' in mid_str:
            # skipping pos {pos} because of mid_str
            continue
        if len(left_str) > 0:
            if left_str[-1] == '#':
                # skipping pos because of left char
                continue
            left_str = left_str[:-1]
        if len(right_str) > 0:
            if right_str[0] == '#':
                # skipping pos because of right char
                continue
            right_str = right_str[1:]
        left_score = search(left_str, left_nums)
        print('left_score', left_str, left_nums, left_score)
        if left_score == 0:
            # no results on left, so don't check right
            continue
        right_score = search(right_str, right_nums)
        print('right_score', right_str, right_nums, right_score)
        res += left_score * right_score
    cache[(plan, tuple(nums))] = res
    return res


if __name__ == '__main__':
    lines = sys.stdin.read().strip().split('\n')
    cache = {}
    for part in range(1, 2):
        answer = 0
        for line_num, line in enumerate(lines):
            if part == 1:
                mult = 1
            else:
                mult = 5
            words = line.split()
            plans = '?'.join([words[0]] * mult)
            counts = list(map(int, words[1].split(','))) * mult
            res = search(plans, counts)
            print(f'{line_num}: {plans} = {res}')
            answer += res

        print(f'answer{part} = {answer}')