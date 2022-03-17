#include "myfun1.h"
#include <stdio.h>

int setStruct1(S11 *s){
    printf("S1 size:%d\n",sizeof(S11));
    s->a = 11;
    s->b = 'a';
}