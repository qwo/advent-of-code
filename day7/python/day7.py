def open_file(p):
    f = open(p, 'r')
    contents = f.read()
    f.close()
    return contents
base = (2 ** 16) - 1

if __name__ == "__main__":
    signals = {}
    lines = open_file("../day7.txt")
    for line in lines.split("\n"):
        line = line.split(" ")
        count = len(line)
        print(line)
        if (count == 3):
            # Assignmentif
            signals[line[2]] = line[0]
        elif (count == 4):
            # Not
            print("Not")
            if signals.get(line[1]):
                signals[line[3]] = ~(signals.get(line[1])) % base
        elif (count == 5):
            # LSHIFT, RSHIFT, AND, OR
            print ("Other")
            print(signals.get(line[0]), signals.get(line[0]))
            x = signals.get(line[0])
            y = signals.get(line[2])
 #             if signals.get(line[0]) else int(line[0])
            if (x and y):
                if (line[1] == "LSHIFT"):
                    signals[line[4]] = int(x) << int(y)
                    print()
                elif (line[1] == "RSHIFT"):
                    signals[line[4]] = int(x) >> int(y)
                elif (line[1] == "AND"):
                    print()
                    signals[line[4]] = int(x) & int(y)
                elif (line[1] == "OR"):
                    print()
                    signals[line[4]] = int(x) | int(y)
    print (signals)
