#ifndef __MYFUN1_H
#define __MYFUN1_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

#pragma pack(1)
typedef struct {
    int a;
    char b;
    short c;
} S11;
#pragma pack()


typedef struct {
    int a;
    char b;
    short c;
} CGO_S11;

void CGO2C_S11(S11 *s11, CGO_S11 *cgo_s11);

void C2CGO_S11(CGO_S11 *cgo_s11, S11 *s11);

int setStruct1(S11 *s);

int CGO_setStruct1(CGO_S11 *cgo_s11);


typedef struct{
 int a;
 int b;
}Foo;

void pass_struct(Foo *in);
void pass_structs(Foo **in,int len);



#ifdef __cplusplus
}
#endif

#endif