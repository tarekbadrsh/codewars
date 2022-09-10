from unittest import TestCase

from main import next_bigger

class TestMain(TestCase):

    def test_next_bigger(self):
        tt = [
            {
                "name": "solution 1",
                "input": (12),
                "expected": 21
            },
            {
                "name": "solution 2",
                "input": (513),
                "expected": 531
            },
            {
                "name": "solution 3",
                "input": (2017),
                "expected": 2071
            },
            {
                "name": "solution 4",
                "input": (414),
                "expected": 441
            }
        ]

        for tc in tt:
            res = next_bigger(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
