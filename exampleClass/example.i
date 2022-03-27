/* File : example.i */
%module exampleClass

%{
#include "example.h"
%}

/* Let's just grab the original header file here */
%include "example.h"
