#include "myfun.h"

int setStruct(S1 *s1){
    printf("S1 size:%d\n",sizeof(S1));
    s1->a = 10;
    s1->b = 'c';
}