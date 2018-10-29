import re
operator = {
        "+":lambda a,b:a+b,
        "-":lambda a,b:a-b,
        "*":lambda a,b:a*b,
        "/":lambda a,b:a/b
    }
"""
Expr      ->    Term ExprTail
ExprTail  ->    + Term ExprTail
          |     - Term ExprTail
          |     null

Term      ->    Factor TermTail
TermTail  ->    * Factor TermTail
          |     / Factor TermTail
          |     null

Factor    ->    (Expr)
          |     num
reference:https://zhuanlan.zhihu.com/p/24035780
"""
grammars = {
            "E":    ["T ET"],
            "ET":   ["+ T ET",
                     "- T ET",
                     "null"],
            "T":    ["F TT"],
            "TT":   ["* F TT",
                     "/ F TT",
                     "null",],
            "F":    ["BRA",
                     "NUMBER"],
            "BRA":  ["LBRA E RBRA",],
            "END_STATE": r"(null)|(NUMBER)|[+\-*/]|(LBRA)|(RBRA)"
}

class AstBuilder:

    def build_ast(self, tokens):
        root = Node("E")
        if root.build_ast(tokens):
            return root
        else:
            raise ValueError("Error Grammar")


class Node:

    def match_token(self, token):
        token_type = token[0]
        if self.type == "null":
            return True
        if self.type == token_type:
            return True
        if token_type == 'OPERATOR':
            if self.type == token[1]:
                return True
            else:
                return False
        return False

    def __init__(self, type):
        self.type = type
        self.text = None
        self.child = list()

    def __str__(self):
        childs = list()
        for child in self.child:
            childs.append(child.__str__())
        out = "({type}, {text})".format(type=self.type, text=self.text)
        for child in childs:
            if child:
                for line in child.split("\n"):
                        out = out + "\n\t" + line
        return out

    def __repr__(self):
        return self.__str__()


    def build_ast(self, tokens: list):
        # if re.match(grammars["END_STATE"], self.type):
        #     if self.match_token(tokens[0]):
        #         self.text = tokens[0][1]
        #         tokens.pop(0)
        #         return tokens
        #     else:
        #         raise ValueError("Error Grammar")
        if len(tokens) == 0:
            return None
        for grammar in grammars[self.type]:
            # print(grammar)
            # print(tokens)
            # print("\n")
            pop_counter = 0
            grammar_tokens = grammar.split()
            try:
                tmp_nodes = list()
                for grammar_token in grammar_tokens:
                    node = Node(grammar_token)
                    tmp_nodes.append(node)
                    if re.match(grammars["END_STATE"], grammar_token):
                        if node.match_token(tokens[0+pop_counter]):

                            if grammar_token!="null":
                                node.text = tokens[0 + pop_counter][1]#
                                pop_counter += 1
                        else:
                            raise ValueError("Error Grammar")
                    else:
                        offset = node.build_ast(tokens[pop_counter:])
                        if offset is not None:
                            pop_counter += offset

                else:
                    self.child = tmp_nodes
                    return pop_counter
            except ValueError as err:
                pass
        raise ValueError("Error Grammar")



ast_builder = AstBuilder()