import re


operator = {
        "+":lambda a,b:a+b,
        "-":lambda a,b:a-b,
        "*":lambda a,b:a*b,
        "/":lambda a,b:a/b
    }


class VM:
    def __init__(self):
        self.stack = list()

    def push(self, num):
        num = float(num)
        self.stack.append(num)

    def pop(self):
        return self.stack.pop()

    def calculate(self, opcode):
        a = self.pop()
        b = self.pop()
        result = operator[opcode](b,a)
        self.push(result)

    def clear(self):
        self.stack.clear()


    def run(self, instructions):
        self.clear()
        for instruction in instructions:
            if instruction.opcode == "PUSH":
                self.push(instruction.operand)
            else:
                if re.match(r"[+\-*/]", instruction.opcode):
                    self.calculate(instruction.opcode)
                else:
                    raise ValueError("unsupport opcode: " + instruction.opcode)
        if len(self.stack) == 1:
            return self.stack[0]
        elif len(self.stack) == 0:
            raise ValueError("empty stack")
        elif len(self.stack) >1 :
            raise ValueError("more than 1 element in the stack")

vm = VM()