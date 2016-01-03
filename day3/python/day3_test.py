import day3

import unittest


class TestVehicle(unittest.TestCase):

  # Part 1
  def test_sequence1(self):
      sequence1 = ">"
      sleigh = day3.Vehicle((0, 0))
      for char in sequence1:
          sleigh.process_char(char)
      self.assertEqual(len(sleigh.visited_houses), 2, sleigh)

  def test_sequence2(self):
      sequence2 = "^>v<"
      sleigh = day3.Vehicle((0, 0))
      for char in sequence2:
          sleigh.process_char(char)
      self.assertEqual(len(sleigh.visited_houses), 4, sleigh)

  def test_sequence3(self):
      test_sequence3 = "^v^v^v^v^v"
      sleigh = day3.Vehicle((0, 0))
      for char in test_sequence3:
          sleigh.process_char(char)
      self.assertEqual(len(sleigh.visited_houses), 2, sleigh)

  # Part 2
  def test_sequence_1_pt2(self):
      sequence = "^v"
      santas_turn = True
      santa = day3.Vehicle((0, 0))
      robo_santa = day3.Vehicle((0, 0))
      for char in sequence:
        if santas_turn:
            santa.process_char(char)
            santas_turn = False
        else:
            robo_santa.process_char(char)
            santas_turn = True

      self.assertEqual(len(santa.compare_visited_houses(robo_santa)), 3)

  def test_sequence_2_pt2(self):
      sequence = "^>v<"
      santas_turn = True
      santa = day3.Vehicle((0, 0))
      robo_santa = day3.Vehicle((0, 0))
      for char in sequence:
        if santas_turn:
            santa.process_char(char)
            santas_turn = False
        else:
            robo_santa.process_char(char)
            santas_turn = True

      self.assertEqual(len(santa.compare_visited_houses(robo_santa)), 3)

  def test_sequence_3_pt2(self):
      sequence = "^v^v^v^v^v"
      santas_turn = True
      santa = day3.Vehicle((0, 0))
      robo_santa = day3.Vehicle((0, 0))
      for char in sequence:
        if santas_turn:
            santa.process_char(char)
            santas_turn = False
        else:
            robo_santa.process_char(char)
            santas_turn = True

      self.assertEqual(len(santa.compare_visited_houses(robo_santa)), 11)

if __name__ == '__main__':
    unittest.main()