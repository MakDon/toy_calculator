//
//  main.cpp
//  ToyCalculatorCPP
//
//  Created by makdon on 31/10/2018.
//  Copyright © 2018 makdon. All rights reserved.
//

#include <iostream>
#include "Calculator.h"

int main(int argc, const char * argv[]) {
    Calculator calculator = Calculator();
    
    /*-----------test----------*/
    assert(calculator.calculate("1+2+3")==6);
    assert(calculator.calculate("1 + 2 + 3")==6);
    assert(calculator.calculate("1+ 2+ 3")==6);
    assert(calculator.calculate("1.2+2.4+3.4")==7);
    assert(calculator.calculate("1+2-3")==0);
    assert(calculator.calculate("1-2+3")==2);
    assert(calculator.calculate("1+2*3")==7);
    assert(calculator.calculate("1*2-3")==-1);
    assert(calculator.calculate("1+3/2")==2.5);
    assert(calculator.calculate("(1+2)*3")==9);
    assert(calculator.calculate("1-(2+3)")==-4);
    std::cout<<"success";
}
