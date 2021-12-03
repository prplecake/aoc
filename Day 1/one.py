#!python3

def loop_input(input_list):
    i = 1
    inc = 0
    while i < len(input_list):
        if input_list[i-1] < input_list[i]:
            inc += 1
        i += 1
    return inc


def main():
    input = [int(i) for i in open("input", "r").readlines()]
    print(f'Part 1 answer: {loop_input(input)}')
    return

if __name__ == '__main__':
    main()
