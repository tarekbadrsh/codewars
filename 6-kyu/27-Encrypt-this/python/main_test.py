from unittest import main, TestCase

from main import encrypt_this

class TestMain(TestCase):

    def test_encrypt_this(self):
        tt = [ 
            {"input":"Hello", "expected": "72olle"},
            {"input":"good", "expected": "103doo"},
            {"input":"hello world", "expected": "104olle 119drlo"},
            {"input":"", "expected": ""},
            {"input":"A wise old owl lived in an oak", "expected": "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"},
            {"input":"The more he saw the less he spoke", "expected": "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"},
            {"input":"The less he spoke the more he heard", "expected": "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"},
            {"input":"Why can we not all be like that wise old bird", "expected": "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"},
            {"input":"Thank you Piotr for all your help", "expected": "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"}
        ]

        for tc in tt:
            res = encrypt_this(tc["input"])
            self.assertEqual(res, tc["expected"])




# python3 -m unittest discover -v -p '*_test.py'

		
