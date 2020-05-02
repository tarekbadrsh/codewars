from unittest import main, TestCase

from main import sumin,sumax,sumsum

class TestData():
    def __init__(self, input, expected): 
        self.input = input
        self.expected = expected

class TestMain(TestCase):

    def test_sumin(self):
        tt = [TestData(5, 55),TestData(100, 338350)]
        
        for tc in tt:
            res = sumin(tc.input)
            self.assertEqual(res,tc.expected)

    def test_sumax(self):
        tt = [
            {
                "name": "sumax 1",
                "input": 5,
                "expected": 95
            },
            {
                "name": "sumax 2",
                "input": 100,
                "expected": 671650
            }
        ]

        for tc in tt:
            res = sumax(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])
    
    def test_sumsum(self):
        tt = [
            {
                "name": "sumsum 1",
                "input": 5,
                "expected": 150
            },
            {
                "name": "sumsum 2",
                "input": 100,
                "expected": 1010000
            }
        ]

        for tc in tt:
            res = sumsum(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])


# python3 -m unittest discover -s . -p '*_test.py'
# python3 -m unittest discover -v -s . -p '*_test.py'
# python3 -m unittest discover -v -p '*_test.py'
