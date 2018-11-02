//
//  AstNode.cpp
//  ToyCalculatorCPP
//
//  Created by makdon on 2/11/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#include "AstNode.hpp"


#include <string>
#include <vector>
#include <map>
#include <regex>
#include "Calculator.h"

/*
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
 */

vector<string> split(string s,char separator=' ')
{
    vector<string> result = vector<string>();
    int begin = 0;
    if(s.size()==0)
    {
        return result;
    }
    else if (s.size()==1)
    {
        if(s[0]==separator)
        {
            return result;
        }
        else
        {
            result.push_back(s);
            return result;
        }
    }
    for(int index=0;index<s.size();index++)
    {
        if(s[index]==separator)
        {
            result.push_back(string(s.begin()+begin, s.end()+index));
            begin = index+1;
        }
    }
    if(begin<s.size())
    {
        result.push_back(string(s.begin()+begin, s.end()));
    }
    return result;
}



map<string, vector<string>> grammars = {
    {"E",    {"T ET"}},
    {"ET",   {"+ T ET",
        "- T ET",
        "null"},},
    {"T",    {"F TT"},},
    {"TT",   {"* F TT",
        "/ F TT",
        "null"}},
    {"F",    {"BRA",
        "NUMBER"}},
    {"BRA",  {"LBRA E RBRA"}}
};
string end_state = "(null)|(NUMBER)|[+\\-*/]|(LBRA)|(RBRA)";

int AstNode::build_ast(vector<Token>& tokens,int token_index=0)
{
    if(regex_match(type, regex(end_state)))
    {
        if(type!="null")
        {
            if(token_index>=tokens.size()){throw "Error Grammar";}
            else if(match_token(tokens.at(token_index))){text = tokens.at(token_index).text;}
            else{throw "Error Grammar";}
            return 1;
        }
        return 0;
    }
    
    /*
     def build_ast(self, tokens: list, token_index=0):
     for grammar in grammars[self.type]:
     offset = 0
     grammar_tokens = grammar.split()
     try:
     tmp_nodes = list()
     for grammar_token in grammar_tokens:
     node = Node(grammar_token)
     tmp_nodes.append(node)
     offset_ = node.build_ast(tokens, offset+token_index)
     offset += offset_
     else:
     self.child = tmp_nodes
     return offset
     except ValueError:
     pass
     raise ValueError("Error Grammar")
     */
    
    
    
    
    
    int offset = 0;
    vector<string> rules = grammars.at(type);
    for(auto rule:rules)
    {
        vector<string> elements = split(rule);
        for(auto element:elements)
        {
            AstNode node = AstNode(element);
            
        }
    }
}
