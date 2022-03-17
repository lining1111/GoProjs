##20220317
    最近接触到的一个代码开发就是做一个服务，功能涉及文件上传下载、文件解压缩、shell执行、
    以及基于网络协议和c库SDK的通信。这里有之前做的一个go web的项目做基石，同时涉及一些
    c语言才好做的硬件通信功能，综合考虑下来，go和c/c++的混编最为合适。
    go调用c/c++时，可以分两种情况
    1,先将c/c++编译成so库，基于函数和变量符号唯一性考虑，这里需要注意的是c++文件内不能有类存在，
        所以h文件内需要添加条件编译限制 
            头部
            #ifdef __cplusplus
            extern "C" {
            #endif
            尾部
            #ifdef __cplusplus
            }
            #endif
        在go文件中 
        头部 同时 import "C"后面空一行
        /*
        #include "../c/myfun1.h"
        #cgo LDFLAGS: -Wl,-rpath="../c" -L../c -lmyfun1
        */
        import "C"
    2. 只编写c c++文件，没有h文件，同时在c c++文件头部添加 #ifdef xxx 尾部添加 #endif ，go文件中
         /*
        #cgo CFLAGS: -Dxxx
        extern 函数名;
        */
        import "C"
        上面方法中有一个可以兼容c/c++ 文件编写习惯的方式，就是c/c++ h 文件按照原本的习惯编写
        在go文件中头部使用函数调用的方式，将想要调用的函数包裹进来
        /*
        #include "xxx.h"
        int add2(int a, int b){
            return add(a,b);
        }
        */
        import "C"

    以上内容主要参考了go sqlite库的方式