import re


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
            "F":    ["NUMBER",
                     "BRA"],
            "BRA":  ["LBRA E RBRA",],
            "END_STATE": r"(null)|(NUMBER)|[+\-*/]|(LBRA)|(RBRA)"
}


def build_ast(tokens):
    root = Node("E")
    offset = root.build_ast(tokens, token_index=0)
    if offset == len(tokens):
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
        return False

    def __init__(self, type):
        self.type = type
        self.text = None
        self.child = list()

    def build_ast(self, tokens: list, token_index=0):

        if re.match(grammars["END_STATE"], self.type):
            if self.type != "null":
                if token_index >= len(tokens):
                    raise ValueError("Error Grammar")
                if self.match_token(tokens[token_index]):
                    self.text = tokens[token_index][1]
                else:
                    raise ValueError("Error Grammar")
                return 1
            return 0

        for grammar in grammars[self.type]:
            offset = 0
            grammar_tokens = grammar.split()
            tmp_nodes = list()
            try:
                for grammar_token in grammar_tokens:
                    node = Node(grammar_token)
                    tmp_nodes.append(node)
                    offset += node.build_ast(tokens, offset+token_index)
                else:
                    self.child = tmp_nodes
                    return offset
            except ValueError:
                pass
        raise ValueError("Error Grammar")

    """
    # previous version:
    def build_ast(self, tokens: list, token_index=0):
        for grammar in grammars[self.type]:
            offset = 0
            grammar_tokens = grammar.split()
            try:
                tmp_nodes = list()
                for grammar_token in grammar_tokens:
                    node = Node(grammar_token)
                    tmp_nodes.append(node)
                    if re.match(grammars["END_STATE"], grammar_token):
                        if grammar_token != "null":
                            if offset + token_index >= len(tokens):
                                raise ValueError("Error Grammar")
                            if node.match_token(tokens[token_index + offset]):
                                node.text = tokens[token_index + offset][1]
                                offset += 1
                            else:
                                raise ValueError("Error Grammar")
                    else:
                        offset_ = node.build_ast(tokens, offset+token_index)
                        if offset_ is not None:
                            offset += offset_
                else:
                    self.child = tmp_nodes
                    return offset
            except ValueError:
                pass
        raise ValueError("Error Grammar")"""

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



