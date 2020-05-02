from unittest import main, TestCase

from main import abbrevName, abbrevNameB

class TestData():
    def __init__(self, input, expected): 
        self.input = input
        self.expected = expected

class TestMain(TestCase):

    def test_sumin(self):
        tt = [
            TestData("Sam Harris", "S.H"),
            TestData("Patrick Feenan", "P.F"),
            TestData("Evan Cole", "E.C"),
            TestData("P Favuzzi", "P.F"),
            TestData("David Mendieta", "D.M"),
            TestData("te be", "T.B"),
            TestData("aaa sss", "A.S"),
        ]
        
        for tc in tt:
            res = abbrevName(tc.input)
            self.assertEqual(res,tc.expected)

        for tc in tt:
            res = abbrevNameB(tc.input)
            self.assertEqual(res,tc.expected)


# python3 -m unittest discover -s . -p '*_test.py'
# python3 -m unittest discover -v -s . -p '*_test.py'
# python3 -m unittest discover -v -p '*_test.py'
