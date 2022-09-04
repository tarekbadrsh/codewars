from unittest import TestCase

from main import is_valid_walk

class TestMain(TestCase):

    def test_is_valid_walk(self):
        tt = [
            {
                "name": "solution 1",
                "input": (['n','s','n','s','n','s','n','s','n','s']),
                "expected": True
            },
            {
                "name": "solution 2",
                "input": (['w','e','w','e','w','e','w','e','w','e','w','e']),
                "expected": False
            },
            {
                "name": "solution 3",
                "input": (['w']),
                "expected": False
            },
            {
                "name": "solution 4",
                "input": (['n','n','n','s','n','s','n','s','n','s']),
                "expected": False
            },
        ]

        for tc in tt:
            res = is_valid_walk(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
