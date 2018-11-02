//
//  structs.h
//  ToyCalculatorCPP
//
//  Created by makdon on 2/11/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#ifndef structs_h
#define structs_h
struct Instruction
{
    int opcode;
    double operand;
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
