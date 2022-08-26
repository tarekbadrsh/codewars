from unittest import TestCase

from main import alternate

class TestMain(TestCase):

    def test_alternate(self):
        tt = [
            {
                "name": "alternate 1",
                "input": (5, "true", "false"),
                "expected": ["true", "false", "true", "false", "true"]
            },
            {
                "name": "alternate 2",
                "input": (20, "blue", "red"),
                "expected": ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
            },
            {
                "name": "alternate 3",
                "input": (0, "", ""),
                "expected": []
            },
        ]

        for tc in tt:
            res = alternate(tc["input"][0],tc["input"][1], tc["input"][2])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
