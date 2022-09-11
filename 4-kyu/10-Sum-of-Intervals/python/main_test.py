from unittest import TestCase

from main import sum_of_intervals, sum_of_intervalsA
 
class TestMain(TestCase):

    def test_sum_of_intervals(self):
        tt = [
            {
                "name": "solution 1",
                "input": ([(1, 5)]),
                "expected": 4
            },
            {
                "name": "solution 2",
                "input": ([(1, 5), (6, 10)]),
                "expected": 8
            },
            {
                "name": "solution 3",
                "input": ([(1, 5), (1, 5)]),
                "expected": 4
            },
            {
                "name": "solution 4",
                "input": ([(1, 4), (7, 10), (3, 5)]),
                "expected": 7
            },
            {
                "name": "solution 5",
                "input": ([(-1_000_000_000, 1_000_000_000)]),
                "expected": 2_000_000_000
            },
            {
                "name": "solution 6",
                "input": ([(0, 20), (-100_000_000, 10), (30, 40)]),
                "expected": 100_000_030
            },
        ]

        for tc in tt:
            res = sum_of_intervals(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])
        for tc in tt:
            res = sum_of_intervalsA(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
