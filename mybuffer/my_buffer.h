#include <string>
#include <vector>

class MyBuffer {
public:
	std::string* s_;

	std::vector<int> vectorInt;

	MyBuffer(int size);
	~MyBuffer();

	int Size() const;
	char* Data();

    void Push(int value);

    int Pop();
};
