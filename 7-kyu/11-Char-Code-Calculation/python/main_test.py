from unittest import TestCase

from main import get_middle

class TestMain(TestCase):

    def test_get_middle(self):
        tt = [
            {
                "name": "get_middle 1",
                "input": ("test"),
                "expected": "es"
            },
            {
                "name": "get_middle 2",
                "input": ("testing"),
                "expected": "t"
            },
            {
                "name": "get_middle 3",
                "input": ("middle"),
                "expected": "dd"
            },
            {
                "name": "get_middle 3",
                "input": ("A"),
                "expected": "A"
            },
        ]

        for tc in tt:
            res = get_middle(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
