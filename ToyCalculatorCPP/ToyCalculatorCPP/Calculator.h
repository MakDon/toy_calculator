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
#include "AstNode.hpp"
using namespace std;






class Calculator
{
public:
    double calculate(const string& raw);
    vector<Token> tokenize(const string& raw);
    AstNode parse(const vector<Token>&);
    vector<Instruction> generate_instructions(const AstNode&);
    double run_instructions(const vector<Instruction>&);
};
#endif /* Calculator_h */
