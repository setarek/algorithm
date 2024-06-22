import unittest
from bubblesort import bubble_sort

class TestBubbleSort(unittest.TestCase):
    
    def test_bubble_sort(self):
        test_cases = [
            ([5, 2, 9, 1, 5, 6], [1, 2, 5, 5, 6, 9]),
            ([3, 0, 2, 5, -1, 4, 1], [-1, 0, 1, 2, 3, 4, 5]),
            ([], []),
            ([1], [1]),
            ([2, 1], [1, 2])
        ]
        
        for input_arr, expected in test_cases:
            with self.subTest(input_arr=input_arr, expected=expected):
                self.assertEqual(bubble_sort(input_arr), expected)

if __name__ == '__main__':
    unittest.main()