from unittest import TestCase

from main import solution

class TestMain(TestCase):

    def test_solution(self):
        tt = [
            {
                "name": "solution 1",
                "input": ("45", "1"),
                "expected": "1451"
            },
            {
                "name": "solution 2",
                "input": ("13", "200"),
                "expected": "1320013"
            },
            {
                "name": "solution 3",
                "input": ("Soon", "Me"),
                "expected": "MeSoonMe"
            },
            {
                "name": "solution 4",
                "input": ("U", "False"),
                "expected": "UFalseU"
            },
        ]

        for tc in tt:
            res = solution(tc["input"][0],tc["input"][1])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
