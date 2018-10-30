from tokenizer import tokenizer
from ast_builder import build_ast
from code_generator import generate
from toy_vm import vm


class Calculator:

    def __init__(self):
        self.tokenizer = tokenizer
        self.build_ast = build_ast
        self.generate = generate
        self.vm = vm

    def calculate(self, input_str):
        tokens = self.tokenizer.build_tokens(input_str)
        ast_root = self.build_ast(tokens)
        ins = self.generate(ast_root)
        result = self.vm.run(ins)
        return result


if __name__ == '__main__':
    calculator = Calculator()
    print(calculator.calculate("1+2*(3+5+4"))
