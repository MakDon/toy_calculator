from tokenizer import tokenizer
from ast_builder import ast_builder
from code_generator import generate
from toy_vm import vm

class Calculator:

    def __init__(self):
        self.tokenizer = tokenizer
        self.ast_builder = ast_builder
        self.generate = generate
        self.vm = vm

    def calculate(self, input_str):
        tokens = self.tokenizer.build_tokens(input_str)
        ast_head = self.ast_builder.build_ast(tokens)
        ins = self.generate(ast_head)
        result = self.vm.run(ins)
        return result

if __name__ == '__main__':
    calculator = Calculator()
    print(calculator.calculate("1-2+3"))

