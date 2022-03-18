#include <string>

class MyBuffer {
public:
	std::string* s_;

	MyBuffer(int size);
	~MyBuffer();

	int Size() const;
	char* Data();
};
