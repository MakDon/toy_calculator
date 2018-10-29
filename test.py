import unittest
from calculator import Calculator


calculator = Calculator()


class demoTest(unittest.TestCase):
    def test_calculate0(self):
        question = "1+2+3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate1(self):
        question = "1 + 2 + 3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate2(self):
        question = "1+ 2+ 3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate3(self):
        question = "1*2+3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate4(self):
        question = "1+2*3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate5(self):
        question = "1-2+3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate6(self):
        question = "1+2/3"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate7(self):
        question = "1+2*(3+5)+4"
        self.assertEqual(calculator.calculate(question), eval(question))

if __name__=='__main__':
     unittest.main()