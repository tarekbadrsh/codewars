from unittest import main, TestCase

from main import decipher_this

class TestMain(TestCase):

    def test_decipher_this(self):
        tt = [ 
            {"input": "65 119esi 111dl 111lw 108dvei 105n 97n 111ka","expected":"A wise old owl lived in an oak"},
            {"input": "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp","expected":"The more he saw the less he spoke"},
            {"input": "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare","expected":"The less he spoke the more he heard"},
            {"input": "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri","expected":"Why can we not all be like that wise old bird"},
            {"input": "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple","expected":"Thank you Piotr for all your help"}
        ]

        for tc in tt:
            res = decipher_this(tc["input"])
            self.assertEqual(res, tc["expected"])




# python3 -m unittest discover -v -p '*_test.py'

		
