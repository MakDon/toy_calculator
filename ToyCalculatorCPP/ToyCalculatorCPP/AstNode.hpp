//
//  AstNode.hpp
//  ToyCalculatorCPP
//
//  Created by makdon on 2/11/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#ifndef AstNode_hpp
#define AstNode_hpp
#include <string>
#include <vector>
#include "structs.h"
#include <stdio.h>
class AstNode
{
public:
    AstNode(std::string typ){type = typ;};
    std::string type;
    std::string text;
    std::vector<AstNode> childs;
    int build_ast(const std::vector<Token>&, int);
    bool match_token(Token);
};
#endif /* AstNode_hpp */
