#include <memory>
using namespace std;

int main() {

    shared_ptr<int> sp1 = make_shared<int>(1);
    shared_ptr<int> sp2(new int(2));
    shared_ptr<int> sp3(sp2);
    cout<<sp<<endl;

    return 0;
}