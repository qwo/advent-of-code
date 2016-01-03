#!/usr/local/bin/python3

# current_position = (0, 0)
# visited_houses = set(current_position)


def open_file(p):
    f = open(p, 'r')
    contents = f.read()
    f.close()
    return contents


class Vehicle(object):
    """docstring for Vehicle"""
    def __init__(self, current_position):
        super(Vehicle, self).__init__()
        self.current_position = current_position
        self.visited_houses = set()
        self.visited_houses.add(current_position)

    def process_char(self, c):
        cp = self.current_position
        if c == "^":
            self.current_position = (cp[0], cp[1] + 1)
            self.visited_houses.add(self.current_position)
        elif c == "v":
            self.current_position = (cp[0], cp[1] - 1)
            self.visited_houses.add(self.current_position)
        elif c == ">":
            self.current_position = (cp[0] + 1, cp[1])
            self.visited_houses.add(self.current_position)
        elif c == "<":
            self.current_position = (cp[0] - 1, cp[1])
            self.visited_houses.add(self.current_position)

    def __str__(self):
        return str({"current_position": self.current_position, "visited_houses": self.visited_houses})

if __name__ == "__main__":
    f = open_file('test_input.txt')

    sleigh = Vehicle((0, 0))
    for char in f:
        sleigh.process_char(char)
    print (sleigh)
    # Parse Character
    # Put Direction
    # calculate house visited
    # visited array
