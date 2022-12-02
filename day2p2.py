from util.utils import read_input_file


def decrypt_choice(choice):
    match choice:
        case 'A':
            return 'rock'
        case 'B':
            return 'paper'
        case 'C':
            return 'scissors'


def decrypt_guide(outcome):
    match outcome:
        case 'X':
            return 'lose'
        case 'Y':
            return 'tie'
        case 'Z':
            return 'win'


def ScoreGame(opponentChoice, outcome):
    _matrix = {
        "rock": {
            "win": "paper",
            "lose": "scissors",
        },
        "paper": {
            "win": "scissors",
            "lose": "rock",
        },
        "scissors": {
            "win": "rock",
            "lose": "paper",
        },
    }
    if outcome == 'tie':
        userChoice = opponentChoice
    else:
        userChoice = _matrix[opponentChoice][outcome]
    match userChoice:
        case 'rock':
            points = 1
        case 'paper':
            points = 2
        case 'scissors':
            points = 3
    return points


def parse_choice(line):
    choices = line.strip().split(' ')
    op_choice = decrypt_choice(choices[0])
    outcome = decrypt_guide(choices[1])
    return (op_choice, outcome)


def main():
    _input = read_input_file('./data/day2')
    total_score = 0
    for line in _input:
        (op_choice, outcome) = parse_choice(line)
        my_choice_points = ScoreGame(op_choice, outcome)
        match outcome:
            case 'win':
                this_score = my_choice_points + 6
            case 'tie':
                this_score = my_choice_points + 3
            case 'lose':
                this_score = my_choice_points
        total_score += this_score
        print(f'{line}::{outcome}//{my_choice_points}::total score this round: {this_score}')
    print(f'total score: {total_score}')
    return


if __name__ == '__main__':
    main()
