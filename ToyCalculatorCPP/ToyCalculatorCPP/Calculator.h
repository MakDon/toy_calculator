//
//  Calculator.h
//  ToyCalculatorCPP
//
//  Created by makdon on 31/10/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#ifndef Calculator_h
#define Calculator_h
#include <string>
#include <vector>
using namespace std;


struct Instruction
{
    int opcode;
    double operand;
};

struct Token
{
    string type;
    string text;
    Token(string typ, string txt)
    {
        type = typ;
        text = txt;
    }
};



class Calculator
{
public:
    double calculate(string raw);
    vector<Token> tokenize(string raw);
    AstNode parse(vector<Token>&);
    vector<Instruction> generate_instructions(AstNode);
    //double run_instructions(vector<Instruction>);
};
#endif /* Calculator_h */
