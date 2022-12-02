from util.utils import read_input_file


def decrypt_choice(choice):
    match choice:
        case 'A' | 'X':
            return 'rock'
        case 'B' | 'Y':
            return 'paper'
        case 'C' | 'Z':
            return 'scissors'


# https://github.com/prplecake/Python-ConsoleGames/blob/3618249db3aca909c97d10e240034e29db859f71/RockPaperScissors.py#L26-L44
def ScoreGame(opponentChoice, userChoice):
    _matrix = {
        "rock": {
            "rock": "tie",
            "paper": "lose",
            "scissors": "win",
        },
        "paper": {
            "rock": "win",
            "paper": "tie",
            "scissors": "lose",
        },
        "scissors": {
            "rock": "lose",
            "paper": "win",
            "scissors": "tie",
        },
    }
    match userChoice:
        case 'rock':
            points = 1
        case 'paper':
            points = 2
        case 'scissors':
            points = 3
    result = _matrix[userChoice][opponentChoice]
    return (points, result)


def parse_choice(line):
    choices = line.strip().split(' ')
    op_choice = decrypt_choice(choices[0])
    my_choice = decrypt_choice(choices[1])
    return (my_choice, op_choice)


def main():
    _input = read_input_file('./data/day2')
    total_score = 0
    for line in _input:
        (my_choice, op_choice) = parse_choice(line)
        my_choice_points, result = ScoreGame(op_choice, my_choice)
        match result:
            case 'win':
                this_score = my_choice_points + 6
            case 'tie':
                this_score = my_choice_points + 3
            case 'lose':
                this_score = my_choice_points
        total_score += this_score
        print(f'{line}::{result}//{my_choice_points}::total score this round: {this_score}')
    print(f'total score: {total_score}')
    return


if __name__ == '__main__':
    main()
