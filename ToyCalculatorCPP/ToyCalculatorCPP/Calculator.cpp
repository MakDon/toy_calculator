//
//  Calculator.cpp
//  ToyCalculatorCPP
//
//  Created by makdon on 31/10/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <regex>
#include <map>
#include "Calculator.h"
using namespace std;

map<string, string> token_patterns={
    {"NUMBER","^([0-9]+\\.)?[0-9]+"},
    {"+", "^\\+"},
    {"-", "^-"},
    {"*", "^\\*"},
    {"/", "^/"},
    {"LBRA", "^\\("},
    {"RBRA", "^\\)"},
    {"SEPARATOR", "^( |\\n)"},
};




double Calculator::calculate(string raw)
{
    vector<Token> tokens = tokenize(raw);
    AstNode root = parse(tokens);
    vector<Instruction> instructions = generate_instructions(root);
    //double result = run_instructions(instructions);
    return 0;
}


vector<Token> Calculator::tokenize(string raw)
{
    vector<Token> tokens = vector<Token>();
    int offset = 0;
    bool match_success_flag = true;
    while(match_success_flag && offset<raw.size())
    {
        match_success_flag = false;
        for(auto pattern_pair:token_patterns)
        {
            smatch match_result;
            basic_regex<char> pattern = regex(pattern_pair.second);
            //raw = string(raw.begin()+offset, raw.end());
            //if(regex_search(raw, match_result, regex(pattern_pair.second)))
            if(offset<raw.size() && regex_search(raw.cbegin()+offset, raw.cend(), match_result, pattern, regex_constants::match_not_null))
            {
                Token token = Token(pattern_pair.first, match_result.str());
                tokens.push_back(token);
                offset += match_result.str().size();
                match_success_flag = true;
            }
        }
    }
    return tokens;
}


AstNode Calculator::parse(const vector<Token>& tokens)
{
    AstNode root = AstNode();
    return root;
    
}
vector<Instruction> Calculator::generate_instructions(AstNode)
{
    vector<Instruction> instructions = vector<Instruction>();
    return instructions;
}

//double run_instructions(vector<Instruction>)
//{
//    return 0;
//}
