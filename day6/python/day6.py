class LightBoard(object):
   """docstring for LightBoard"""

   def __init__(self):
      super(LightBoard, self).__init__()
      self.state = [[False for x in range(1000)] for x in range(1000)]

   def toggle_range(self, i):
       for x in range(int(i['start'][0]), int(i['end'][0]) + 1):
          for y in range(int(i['start'][1]), int(i['end'][1]) + 1):
              if (self.state[x][y] is False):
                 self.state[x][y] = True
              else:
                 self.state[x][y] = False

   def toggle_on_range(self, i):
      for x in range(int(i['start'][0]), int(i['end'][0]) + 1):
         for y in range(int(i['start'][1]), int(i['end'][1]) + 1):
                 self.state[x][y] = True

   def toggle_off_range(self, i):
      for x in range(int(i['start'][0]), int(i['end'][0]) + 1):
         for y in range(int(i['start'][1]), int(i['end'][1]) + 1):
                 self.state[x][y] = False

   def count_lights(self):
       count = 0
       for x in range(0, 1000):
          for y in range(0, 1000):
              if (self.state[x][y] is True):
                 count = count + 1
       return count

   def read_instruction(self, instruct):
       """ Format of instructions are constructed into  {'start': ['475', '711'], 'end': ['921', '882']} """
       instruct = instruct.split(" ")
       # print(self.count_lights())
       if (instruct[0] == "toggle"):
          x = instruct[1].split(",")
          y = instruct[3].split(",")
          self.toggle_range({"start": x, "end": y})
       else:
           if (instruct[1] == "on"):
              x = instruct[2].split(",")
              y = instruct[4].split(",")
              self.toggle_on_range({"start": x, "end": y})
           elif(instruct[1] == "off"):
               x = instruct[2].split(",")
               y = instruct[4].split(",")
               self.toggle_off_range({"start": x, "end": y})
       # print(instruct)


def open_file(p):
    f = open(p, 'r')
    contents = f.read()
    f.close()
    return contents


if __name__ == "__main__":
    lines = open_file("../day7.txt")
    board = LightBoard()
    for line in lines.split("\n"):
        board.read_instruction(line)
    print(board.count_lights())
