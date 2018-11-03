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
#include "structs.h"
#include "AstNode.hpp"
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

typedef double(*pDouble)(double);
map<string, std::function<double(double&,double&)>> op_actions={
    {"+",[](double& a, double& b){return a+b;}},
    {"-",[](double& a, double& b){return a-b;}},
    {"*",[](double& a, double& b){return a*b;}},
    {"/",[](double& a, double& b){return a/b;}}
};


double Calculator::calculate(const string& raw)
{
    vector<Token> tokens = tokenize(raw);
    AstNode root = parse(tokens);
    vector<Instruction> instructions = generate_instructions(root);
    double result = run_instructions(instructions);
    return result;
}


vector<Token> Calculator::tokenize(const string& raw)
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
                if(pattern_pair.first!="SEPARATOR")
                {
                    Token token = Token(pattern_pair.first, match_result.str());
                    tokens.push_back(token);
                }
                offset += match_result.str().size();
                match_success_flag = true;
            }
        }
    }
    return tokens;
}

AstNode Calculator::parse(const vector<Token>& tokens)
{
    AstNode root = AstNode("E");
    root.build_ast(tokens, 0);
    return root;
    
}
vector<Instruction> Calculator::generate_instructions(const AstNode& node)
{
    //vector<Instruction> instructions = vector<Instruction>();
    if(node.type == "E" or node.type == "T" or node.type == "F" or node.type == "BRA")
    {
        vector<Instruction> instructions = vector<Instruction>();
        for(auto child:node.childs)
        {
            vector<Instruction> tmp =generate_instructions(child);
            instructions.insert(instructions.cend(), tmp.cbegin(), tmp.cend());
        }
        return instructions;
    }
    
    else if(node.type == "ET" or node.type == "TT")
    {
        vector<Instruction> instructions = vector<Instruction>();
        if(node.childs.size()>1)
        {
            vector<Instruction> tmp =generate_instructions(node.childs.at(1));
            instructions.insert(instructions.cend(), tmp.cbegin(), tmp.cend());
            tmp =generate_instructions(node.childs.at(0));
            instructions.insert(instructions.cend(), tmp.cbegin(), tmp.cend());
            tmp =generate_instructions(node.childs.at(2));
            instructions.insert(instructions.cend(), tmp.cbegin(), tmp.cend());
        }
        return instructions;
    }
    
    else if(node.type == "LBRA" or node.type == "RBRA")
    {
        return vector<Instruction>();
    }
    
    else if(node.type == "NUMBER")
    {
        vector<Instruction> instructions={
            Instruction("PUSH", node.text)
        };
        return instructions;
    }
    
    else
    {
        vector<Instruction> instructions={
            Instruction(node.text, node.text)
        };
        return instructions;
    }
    //return vector<Instruction>();
}

double Calculator::run_instructions(const vector<Instruction>& instructions)
{
    vector<double> stack = vector<double>();
    for(auto instruction:instructions)
    {
        if(instruction.opcode=="PUSH")
        {
            stack.push_back(stod(instruction.operand));
        }
        else
        {
            
            double b = stack.at(stack.size()-1);
            double a = stack.at(stack.size()-2);
            double tmp = op_actions.at(instruction.opcode)(a, b);
            stack.pop_back();
            stack.pop_back();
            stack.push_back(tmp);
        }
    }
    if(stack.size()==1)
    {
        return stack[0];
    }
    else
    {
        throw "More than 1 frame in stack";
    }
    return 0;
}
