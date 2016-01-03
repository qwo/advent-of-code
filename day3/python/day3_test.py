import day3

import unittest


class TestVehicle(unittest.TestCase):

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

if __name__ == '__main__':
    unittest.main()