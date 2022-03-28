#include "myfun1.h"
#include <stdio.h>
#include <cstring>


void setStructS11(S11 *s) {
//    printf("S1 size:%d\n",sizeof(S11));
    s->a = 11;
    s->b = 'a';
    s->c = 0x1234;
}

void setStructCGO_S11(CGO_S11 *s){
    S11 s11;
    setStructS11(&s11);
    s->a = s11.a;
    s->b = s11.b;
    s->c = s11.c;
}

void SerialStructS11(uint8_t *out, uint32_t *len, S11 *s){
    if(out ==nullptr || len == nullptr){
        return;
    }

    memcpy(out, s,sizeof(S11));
    *len = sizeof(S11);

}