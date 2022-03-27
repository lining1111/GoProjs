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
##20220318
###mybuffer
    mybuffer.h是一个只含public成员和函数的类，编写方式符合c++风格
    通过my_buffer_capi的两个文件，将 struct形式重新定义了mybuffer类，供给go文件使用
    这里查到的c++上 class和struct的区别，就是class默认权限是private，而struct是public
##20220325
    通常的go c/c++混编，是基于go的package模式的，就是说 go和调用的c/c++文件在同一级目录上
##20220327
    myfun1：go通过 swig调用c++的例子
    myfun1 是参考swig example go class的例子做的，需要强调的一点是，函数有返回值的必须返回。
    