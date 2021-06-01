from unittest import main, TestCase

from main import anagrams


class TestMain(TestCase):

    def test_dirReduc(self):
        tt = [
            {
                "input1": "abba", "input2": ['aabb', 'abcd', 'bbaa', 'dada'], "expected": ['aabb', 'bbaa']
            },
            {"input1": "racer", "input2":  [
                'crazer', 'carer', 'racar', 'caers', 'racer'], "expected": ['carer', 'racer']},
            {"input1": "a", "input2":  [
                'a', 'b', 'c', 'd'], "expected": ['a']},
            {"input1": "ab", "input2":  [
                'cc', 'ac', 'bc', 'cd', 'ab', 'ba', 'racar', 'caers', 'racer'], "expected": ['ab', 'ba']},
            {"input1": "abba", "input2":  ['a', 'b', 'c', 'd', 'aabb', 'bbaa', 'abab', 'baba', 'baab', 'abcd',
                                           'abbba', 'baaab', 'abbab', 'abbaa', 'babaa'], "expected": ['aabb', 'bbaa', 'abab', 'baba', 'baab']},
            {"input1": "big", "input2":  [
                'gig', 'dib', 'bid', 'biig'], "expected": []},
        ]

        for tc in tt:
            res = anagrams(tc["input1"], tc["input2"])
            self.assertEqual(res, tc["expected"])


# python -m unittest discover -v -p '*_test.py'
