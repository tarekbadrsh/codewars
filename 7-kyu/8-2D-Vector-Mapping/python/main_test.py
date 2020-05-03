from unittest import main, TestCase

from main import map_vector, map_vector2


class TestMain(TestCase):

    def test_map_vector(self):
        tt = [
            {
                "name": "map_vector 1",
                "input": [(46, 58), (0, 0, 100), (0, 0, 100)],
                "expected": (46, 58)
            },
            {
                "name": "map_vector 1",
                "input": [(50, 88), (-25, 128, 100), (50, 50, 100)],
                "expected": (125, 10)
            },
            {
                "name": "map_vector 1",
                "input": [(120, 58), (100, 76, 36), (120, -62, 50)],
                "expected": (147.78, -87.0)
            },
            {
                "name": "map_vector 1",
                "input": [(5, 5), (10, 20, 42), (-100, -42, 60)],
                "expected": (-107.14, -63.43)
            },
        ]

        for tc in tt:
            res = map_vector(*tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])


        for tc in tt:
            res = map_vector2(*tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])



# python3 -m unittest discover -v -p '*_test.py'
