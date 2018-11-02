from ast_builder import Node


class Instruction:
    def __init__(self, opcode, operand=None):
        self.opcode = opcode
        self.operand = operand

    def __str__(self):
        return "({0},{1})".format(self.opcode, self.operand)

    def __repr__(self):
        return self.__str__()


def generate(ast_node: Node):
    if ast_node.type == "E" or ast_node.type == "T":
        tmp = list()
        for child in ast_node.child:
            tmp.extend(generate(child))
        return tmp

    elif ast_node.type == "ET" or ast_node.type == "TT":
        tmp = list()
        if len(ast_node.child) > 1:
            tmp.extend(generate(ast_node.child[1]))
            tmp.extend(generate(ast_node.child[0]))
            tmp.extend(generate(ast_node.child[2]))
        return tmp

    elif ast_node.type == 'F':
        return generate(ast_node.child[0])

    elif ast_node.type == "BRA":
        tmp = list()
        for child in ast_node.child:
            child_code = generate(child)
            if child_code is not None:
                tmp.extend(generate(child))

    elif ast_node.type == "LBRA" or ast_node.type == 'RBRA':
        return None

    elif ast_node.type == "NUMBER":
        tmp = list()
        instruction = Instruction("PUSH", ast_node.text)
        tmp.append(instruction)

    else:
        tmp = list()
        instruction = Instruction(ast_node.text)
        tmp.append(instruction)

    return tmp
