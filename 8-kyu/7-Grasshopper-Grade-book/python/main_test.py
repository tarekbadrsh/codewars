from unittest import main, TestCase

from main import get_grade


class TestMain(TestCase):

    def test_get_grade(self):
        tt = [
            {
                "name": "get_grade 1",
                "input": [95, 90, 93],
                "expected": 'A'
            },
            {
                "name": "get_grade 1",
                "input": [100, 85, 96],
                "expected": 'A'
            }, 
            {
                "name": "get_grade 1",
                "input": [92, 93, 94],
                "expected": 'A'
            },
            {
                "name": "get_grade 1",
                "input": [70, 70, 100],
                "expected": 'B'
            }, 
            {
                "name": "get_grade 1",
                "input": [82, 85, 87],
                "expected": 'B'
            }, 
            {
                "name": "get_grade 1",
                "input": [84, 79, 85],
                "expected": 'B'
            }, 
            {
                "name": "get_grade 1",
                "input": [89, 89, 90],
                "expected": 'B'
            },
            {
                "name": "get_grade 1",
                "input": [70, 70, 70],
                "expected": 'C'
            }, 
            {
                "name": "get_grade 1",
                "input": [75, 70, 79],
                "expected": 'C'
            }, 
            {
                "name": "get_grade 1",
                "input": [60, 82, 76],
                "expected": 'C'
            },
            {
                "name": "get_grade 1",
                "input": [65, 70, 59],
                "expected": 'D'
            }, 
            {
                "name": "get_grade 1",
                "input": [66, 62, 68],
                "expected": 'D'
            }, 
            {
                "name": "get_grade 1",
                "input": [58, 62, 70],
                "expected": 'D'
            },
            {
                "name": "get_grade 1",
                "input": [44, 55, 52],
                "expected": 'F'
            }, 
            {
                "name": "get_grade 1",
                "input": [48, 55, 52],
                "expected": 'F'
            }, 
            {
                "name": "get_grade 1",
                "input": [58, 59, 60],
                "expected": 'F'
            },
        ]

        for tc in tt:
            res = get_grade(*tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])


# python3 -m unittest discover -v -p '*_test.py'
