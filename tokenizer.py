import re


class Tokenizer:
    def __init__(self):
        self.token_table = {
            "NUMBER": r"[0-9]+",
            "+": r"\+",
            "-": r"-",
            "*": r"\*",
            "/": r"/",
            "LBRA": r"\(",
            "RBRA": r"\)",
            "SEPARATOR": r"( |\n)"
        }

    def build_tokens(self, raw):
        tokens = list()
        while len(raw) > 0:
            for typ in self.token_table:
                if re.match(self.token_table[typ], raw):
                    token_text = re.match(self.token_table[typ], raw).group()
                    if typ != "SEPARATOR":
                        tokens.append([typ, token_text])
                    raw = raw[len(token_text):]
        return tokens


tokenizer = Tokenizer()

if __name__ == '__main__':
    a = tokenizer.build_tokens("1 + 2 + 3 + 5 + 7")
    b = tokenizer.build_tokens("1+2+3+4")
    c = tokenizer.build_tokens("(1+2) + (3+4)")
    pass