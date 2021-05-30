from unittest import main, TestCase

from main import dirReduc


class TestMain(TestCase):

    def test_dirReduc(self):
        tt = [
            {
                "input": ["WEST", "WEST", "EAST", "EAST", "WEST", "SOUTH", "NORTH", "SOUTH"],
                "expected": ["WEST", "SOUTH"]
            },
            {"input": ["NORTH", "WEST", "EAST", "EAST", "SOUTH"],
                "expected":["NORTH", "EAST", "SOUTH"]},
            {"input": ["SOUTH", "SOUTH", "SOUTH", "NORTH", "SOUTH",
                       "NORTH", "WEST", "EAST"], "expected":["SOUTH", "SOUTH"]},
            {"input": ["WEST", "EAST", "EAST", "SOUTH",
                       "NORTH"], "expected":["EAST"]},
            {"input": ["WEST", "EAST", "SOUTH", "WEST", "EAST"],
                "expected":["SOUTH"]},
            {"input": ["NORTH", "EAST", "SOUTH", "NORTH",
                       "SOUTH", "NORTH"], "expected":["NORTH", "EAST"]},
            {"input": ["NORTH", "SOUTH", "WEST", "NORTH", "NORTH"],
                "expected":["WEST", "NORTH", "NORTH"]},
            {"input": ["NORTH", "NORTH", "NORTH", "NORTH", "SOUTH",
                       "SOUTH", "WEST"], "expected":["NORTH", "NORTH", "WEST"]},
            {"input": ["NORTH", "WEST", "EAST", "WEST",
                       "EAST", "WEST"], "expected":["NORTH", "WEST"]},
            {"input": ["EAST", "WEST", "NORTH", "EAST", "NORTH"],
                "expected":["NORTH", "EAST", "NORTH"]},
            {"input": ["WEST", "NORTH", "EAST", "NORTH", "WEST"],
                "expected":["WEST", "NORTH", "EAST", "NORTH", "WEST"]},
            {"input": ["EAST", "SOUTH", "SOUTH", "SOUTH", "EAST", "NORTH", "NORTH"], "expected":[
                "EAST", "SOUTH", "SOUTH", "SOUTH", "EAST", "NORTH", "NORTH"]},
            {"input": ["NORTH", "EAST", "SOUTH", "EAST",
                       "WEST", "NORTH"], "expected":["NORTH", "EAST"]},
            {"input": ["WEST", "EAST", "SOUTH", "EAST",
                       "WEST", "WEST"], "expected":["SOUTH", "WEST"]},
            {"input": ["SOUTH", "NORTH", "SOUTH", "EAST", "SOUTH", "EAST", "EAST"], "expected":[
                "SOUTH", "EAST", "SOUTH", "EAST", "EAST"]},

        ]

        for tc in tt:
            res = dirReduc(tc["input"])
            self.assertEqual(res, tc["expected"])


# python3 -m unittest discover -v -p '*_test.py'
