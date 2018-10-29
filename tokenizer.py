import re
class Tokenizer:
    def __init__(self):
        self.token_table = {
            "NUMBER": r'[0-9]+',
            "OPERATOR": r"[+\-*/%]",
            "LBRA": r"\(",
            "RBRA": r"\)",
            "SEPARATOR": r"( |\n)"
        }

    """
    def build_tokens(self, raw_str):
        tokens = list()
        if(type(raw_str)!=str):
            raise TypeError("error type in raw_str")
        tmp = ""
        for char in raw_str:
            if re.match(r"([0-9]|\.)+", char):
                tmp += char
            elif char != ' ' and char != '\n':
                tokens.append(self.match_token(char))
            else:
                tokens.append(self.match_token(tmp))
                tmp = ""
        else:
            if tmp !="":
                tokens.append(self.match_token(tmp))

        return tokens
    """
    def build_tokens(self, raw):
        tokens = list()
        while len(raw) > 0:
            for typ in self.token_table:
                if re.match(self.token_table[typ], raw):
                    token_text = re.match(self.token_table[typ], raw).group()
                    if typ!="SEPARATOR":
                        tokens.append([typ, token_text])
                    raw = raw[len(token_text):]
        return tokens


    def match_token(self, chars):
        for typ in self.token_table:
            if chars is not None and chars !=' ' and chars !='\n':
                if re.match(self.token_table[typ], chars):
                    return [typ, chars]
            else:
                return ["SEPARATOR", chars]
        else:
            return ["UNKNOWN", chars]

tokenizer = Tokenizer()

if __name__ == '__main__':
    a = tokenizer.build_tokens("1 + 2 + 3 + 5 + 7")
    b = tokenizer.build_tokens("1+2+3+4")
    c = tokenizer.build_tokens("(1+2) + (3+4)")
    pass