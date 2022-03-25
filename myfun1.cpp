#include "myfun1.h"
#include <stdio.h>


void CGO2C_S11(S11 *s11, CGO_S11 *cgo_s11) {
    s11->a = cgo_s11->a;
    s11->b = cgo_s11->b;
    s11->c = cgo_s11->c;
}

void C2CGO_S11(CGO_S11 *cgo_s11, S11 *s11) {
    cgo_s11->a = s11->a;
    cgo_s11->b = s11->b;
    cgo_s11->c = s11->c;
}

int setStruct1(S11 *s) {
//    printf("S1 size:%d\n",sizeof(S11));
    s->a = 11;
    s->b = 'a';
    s->c = 0x1234;
}

int CGO_setStruct1(CGO_S11 *cgo_s11) {
//    S11 s11;
//    setStruct1(&s11);
//    C2CGO_S11(cgo_s11, &s11);
}

void pass_struct(Foo *in){
    printf("%d:%d\n",in->a,in->b);
//    in->a +=1;
//    in->b +=1;
}
void pass_structs(Foo **in,int len){
    for(int i=0;i <len;i++){
        pass_struct(in[i]);
        in[i]->a +=1;
        in[i]->b +=1;
    }
}