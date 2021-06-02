from unittest import TestCase

from main import strip_comments, strip_comments2

# python -m unittest discover -v -p '*_test.py'


class TestMain(TestCase):

    def test_strip_comments(self):
        tt = [
            {"input1": "apples, pears # and bananas\ngrapes\nbananas !apples",
                "input2": ["#", "!"], "expected": "apples, pears\ngrapes\nbananas"},
            {"input1": "a #b\nc\nd $e f g", "input2": [
                "#", "$"], "expected": "a\nc\nd"},
            {"input1": "apples, pears # and bananas\ngrapes\nbananas !#apples",
                "input2": ["#", "!"], "expected": "apples, pears\ngrapes\nbananas"},
            {"input1": "apples, pears # and bananas\ngrapes\nbananas #!apples",
                "input2": ["#", "!"], "expected": "apples, pears\ngrapes\nbananas"},
            {"input1": "apples, pears # and bananas\ngrapes\navocado @apples",
                "input2": ["@", "!"], "expected": "apples, pears # and bananas\ngrapes\navocado"},
            {"input1": "apples, pears § and bananas\ngrapes\navocado *apples",
                "input2": ["*", "§"], "expected": "apples, pears\ngrapes\navocado"},
            {"input1": "", "input2": ["#", "!"], "expected": ""},
            {"input1": "#", "input2": ["#", "!"], "expected": ""},
            {"input1": "\n§", "input2": ["#", "§"], "expected": "\n"},
        ]

        for tc in tt:
            res = strip_comments(tc["input1"], tc["input2"])
            self.assertEqual(res, tc["expected"])
        for tc in tt:
            res = strip_comments2(tc["input1"], tc["input2"])
            self.assertEqual(res, tc["expected"])
