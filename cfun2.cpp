#include "cfun2.h"

#include <string>

using namespace std;

int add(int a, int b){
    return a+b;
}

int lenStr(char *s){
 string str = string(s);
 return str.length();

}