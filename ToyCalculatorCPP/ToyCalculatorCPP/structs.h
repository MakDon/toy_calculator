//
//  structs.h
//  ToyCalculatorCPP
//
//  Created by makdon on 2/11/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#ifndef structs_h
#define structs_h
#include <string>
struct Instruction
{
    std::string opcode;
    std::string operand;
    Instruction(std::string opc, std::string ope=NULL)
    {
        opcode = opc;
        operand = ope;
    }
};

struct Token
{
    std::string type;
    std::string text;
    Token(std::string typ, std::string txt)
    {
        type = typ;
        text = txt;
    }
};

#endif /* structs_h */
