from operator import itemgetter

from util.utils import read_input_file

def main():
    _input = read_input_file("data/day1")
    i = 0
    elves = {}
    elves[i] = 0
    for line in _input:
        if (line != '\n'):
            elves[i] += int(line.strip())
        else:
            i += 1
            elves[i] = 0
    max_value = max(elves.values())
    max_key = max(elves, key=elves.get)
    print(f'Elf {max_key} has the most snacks: {max_value}')

    n = 3
    res = dict(sorted(elves.items(), key=itemgetter(1), reverse=True)[:n])
    result = sum(res.values())
    print(f'Top {n} value pairs are {str(res)}, so the top three elves have {result} calories.')
    return


if __name__ == '__main__':
    main()
