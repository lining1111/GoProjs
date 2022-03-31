%module exampleStruct

%{
#include "example.h"
%}

/* Let's just grab the original header file here */
%include "stdint.i"
%include "example.h"