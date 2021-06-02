from unittest import TestCase

from main import same_structure_as, same_structure_as2


class TestMain(TestCase):

    def test_dirReduc(self):
        tt = [
            {"input1": [1, 1, 1], "input2": [2, 2, 2], "expected": True},
            {"input1": [1, [1, 1]], "input2": [2, [2, 2]], "expected": True},
            {"input1": [1, [1, 1]], "input2": [[2, 2], 2], "expected": False},
            {"input1": [1, [1, 1]], "input2": [2, [2]], "expected": False}
        ]

        for tc in tt:
            res = same_structure_as(tc["input1"], tc["input2"])
            self.assertEqual(res, tc["expected"])
        for tc in tt:
            res = same_structure_as2(tc["input1"], tc["input2"])
            self.assertEqual(res, tc["expected"])


# python -m unittest discover -v -p '*_test.py'
