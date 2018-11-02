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
    string raw ="12+23+34.56-\n76+4.5";
    double result = calculator.calculate(raw);
    std::cout<<result;
}
