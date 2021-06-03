from unittest import main, TestCase

from main import RomanNumerals


class TestMain(TestCase):

    def test_to_roman(self):
        tt = [
            {"input": 1000, "expected": 'M'},
            {"input": 4, "expected": 'IV'},
            {"input": 1, "expected": 'I'},
            {"input": 1990, "expected": 'MCMXC'},
            {"input": 2008, "expected": 'MMVIII'},
        ]

        for tc in tt:
            res = RomanNumerals.to_roman(tc["input"])
            self.assertEqual(res, tc["expected"])

    def test_from_roman(self):
        tt = [
            {"input": 'XXI', "expected": 21},
            {"input": 'I', "expected": 1},
            {"input": 'IV', "expected": 4},
            {"input": 'MMVIII', "expected": 2008},
            {"input": 'MDCLXVI', "expected": 1666},
        ]

        for tc in tt:
            res = RomanNumerals.from_roman(tc["input"])
            self.assertEqual(res, tc["expected"])


# python -m unittest discover -v -p '*_test.py'
