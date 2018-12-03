"""
Advent of Code Day 1
"""

def get_input():
    with open('input.txt', 'r') as f:
        return f.read().split(', ')

direction = {
    0: 'North',
    1: 'East',
    2: 'South',
    3: 'West'
}

weight = {
    'N': (1, 1),
    'S': (1, -1),
    'W': (-1, 1),
    'E': (1, 1)
}


#def main():
if __name__ == '__main__':
    arr = get_input()
    print([(i[0], i[1:] ) for i in arr])
    compass = 0 # Start Pointing at North 
    starting = (0, 0)
    ending = (0, 0)
    blocks = 0
    last_move = ''
    for k, i in enumerate(arr):
        next_move = i[0]
        count = int(i[1:])
        if k == 0:
            last_move = next_move
            blocks = count
            print('Start', next_move, last_move, count)

        if next_move != last_move:
            print('Alternating', next_move, last_move, count)
            last_move = next_move
            blocks += count
        else:
            last_move = next_move
            blocks -= count

    pass

"""
if __name__ == '__main__':
    import sys
    sys.exit(int(main() or 0))
"""