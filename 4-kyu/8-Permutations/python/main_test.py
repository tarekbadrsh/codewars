from unittest import main, TestCase

from main import permutations, permutations2


class TestMain(TestCase):

    def test_permutations(self):
        tt = [
            {"input": "aa", "expected": {'aa'}},
            {"input": "ab", "expected": {'ab', 'ba'}},
            {"input": "abc", "expected": {
                'cab', 'abc', 'acb', 'bac', 'bca', 'cba'}}
        ]

        for tc in tt:
            res = permutations(tc["input"])
            self.assertEqual(res, tc["expected"])

        for tc in tt:
            res = permutations2(tc["input"])
            self.assertEqual(res, tc["expected"])


# python -m unittest discover -v -p '*_test.py'
