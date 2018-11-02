import unittest
from calculator import Calculator


calculator = Calculator()


class TestCalculator(unittest.TestCase):
    def test_calculate0(self):
        question = "1+2+3+5+7"
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
        question = "1.1+2.3*(3.0+5)+4"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate8(self):
        question = "1+2*(3.235+5)+4"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate9(self):
        question = "1+2*(3+5.78)+4"
        self.assertEqual(calculator.calculate(question), eval(question))

    def test_calculate10(self):
        question = "1+2*(3+5+4"
        try:
            calculator.calculate(question)
        except ValueError as e:
            self.assertEqual(e.args[0], "Error Grammar")

    def test_calculate11(self):
        question = "1+2(3+5)+4"
        try:
            calculator.calculate(question)
        except ValueError as e:
            self.assertEqual(e.args[0], "Error Grammar")

    def test_calculate12(self):
        question = "1+2*(3+5)+"
        try:
            calculator.calculate(question)
        except ValueError as e:
            self.assertEqual(e.args[0], "Error Grammar")

    def test_calculate13(self):
        question = "1.+2*(3+5)+"
        try:
            calculator.calculate(question)
        except ValueError as e:
            self.assertEqual(e.args[0], "Unknown Token")

if __name__=='__main__':
     unittest.main()