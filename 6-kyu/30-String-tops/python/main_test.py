from unittest import TestCase

from main import tops

class TestMain(TestCase):

    def test_tops(self):
        tt = [
            {
                "name": "solution 1",
                "input": (""),
                "expected": ""
            },
            {
                "name": "solution 2",
                "input": ("12"),
                "expected": "2"
            },
            {
                "name": "solution 3",
                "input": ("abcdefghijklmnopqrstuvwxyz12345"),
                "expected": "3pgb"
            },
            {
                "name": "solution 4",
                "input": ("abcdefghijklmnopqrstuvwxyz1236789ABCDEFGHIJKLMN"),
                "expected": "M3pgb"
            },
        ]

        for tc in tt:
            res = tops(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
