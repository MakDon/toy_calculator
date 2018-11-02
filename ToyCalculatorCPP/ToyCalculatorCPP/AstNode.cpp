//
//  AstNode.cpp
//  ToyCalculatorCPP
//
//  Created by makdon on 2/11/2018.
//  Copyright © 2018 makdon. All rights reserved.
//




#include <string>
#include <vector>
#include <map>
#include <regex>

#include "AstNode.hpp"
#include "structs.h"
using namespace std;

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
            result.push_back(string(s.begin()+begin, s.begin()+index));
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
    {"E",    {  "T ET"}},
    {"ET",   {  "+ T ET",
                "- T ET",
                "null"},},
    {"T",    {  "F TT"},},
    {"TT",   {  "* F TT",
                "/ F TT",
                "null"}},
    {"F",    {  "NUMBER",
                "BRA",}},
    {"BRA",  {  "LBRA E RBRA"}}
};
string end_state = "(null)|(NUMBER)|[+\\-*/]|(LBRA)|(RBRA)";



int AstNode::build_ast(const vector<Token>& tokens,int token_index=0)
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

    vector<string> rules = grammars.at(type);
    for(auto rule:rules)
    {
        int offset = 0;
        vector<string> elements = split(rule);
        vector<AstNode> tmp_child_nodes = vector<AstNode>();
        try
        {
            for(auto element:elements)
            {
                AstNode node = AstNode(element);
                offset += node.build_ast(tokens, offset+token_index);
                tmp_child_nodes.push_back(node);
            }
            childs.insert(childs.end(), tmp_child_nodes.cbegin(), tmp_child_nodes.cend());
            return offset;
        } catch (const char* msg) {
            
        }
        
    }
    throw "Error Grammar";
}

bool AstNode::match_token(Token token)
{
    string token_type = token.type;
    if(type == "null")
        return true;
    if(type == token_type)
        return true;
    return false;
}
