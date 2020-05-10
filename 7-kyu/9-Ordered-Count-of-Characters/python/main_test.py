from unittest import main, TestCase

from main import ordered_count

# {"abracadabra", []Tuple{{'a', 5}, {'b', 2}, {'r', 2}, {'c', 1}, {'d', 1}}},
# {"Code Wars", []Tuple{{'C', 1}, {'o', 1}, {'d', 1}, {'e', 1}, {' ', 1}, {'W', 1}, {'a', 1}, {'r', 1}, {'s', 1}}},
# {"", []Tuple{}},



class TestMain(TestCase):

    def test_ordered_count(self):
        tt = [
            {
                "name": "ordered_count 1",
                "input": "abracadabra",
                "expected": [('a', 5), ('b', 2), ('r', 2), ('c', 1), ('d', 1)]
            },
         
        ]

        for tc in tt:
            res = ordered_count(tc["input"])
            self.assertEqual(res, tc["expected"], tc["name"])




# python3 -m unittest discover -v -p '*_test.py'
