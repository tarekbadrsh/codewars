from unittest import TestCase

from main import count_by

class TestMain(TestCase):

    def test_count_by(self):
        tt = [
            {
                "name": "count_by 1",
                "input": (1, 5),
                "expected": [1, 2, 3, 4, 5]
            },
            {
                "name": "count_by 2",
                "input": (2, 5),
                "expected": [2, 4, 6, 8, 10]
            },
            {
                "name": "count_by 3",
                "input": (3, 5),
                "expected": [3, 6, 9, 12, 15]
            },
            {
                "name": "count_by 4",
                "input": (50, 5),
                "expected": [50, 100, 150, 200, 250]
            },
            {
                "name": "count_by 5",
                "input": (100, 5),
                "expected": [100, 200, 300, 400, 500]
            },
         
        ]

        for tc in tt:
            res = count_by(tc["input"][0],tc["input"][1])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
