#include <iostream>
#include <vector>

using namespace std;

int64_t NOD(int64_t a, int64_t b) {
    if (a < b) {
        swap(a, b);
    }
    while (a % b != 0) {
        a = a % b;
        swap(a, b);
    }
    return b;
}

int main() {
    // считаем сумму чисел от 1 до n и затем вычитаем сумму чисел которые делятся на a, сумму чисел которые делятся на b
    // и прибавляем к ним сумму числе которые делятся и на a и на b

    int64_t n, a, b;
    int64_t result = 0;
    cin >> n >> a >> b;

    result = n * (n + 1) / 2;


    if (a % b == 0 || b % a == 0) {
        a = min(a, b);
        result -= (2 * a + (n / a - 1) * a) * (n / a) / 2;
    } else {
        int64_t nod = NOD(a, b);
        int64_t c = a * b / nod;

        result = result - ((2 * a + (n / a - 1) * a) * (n / a) / 2) - ((2 * b + (n / b - 1) * b) * (n / b) / 2) +
                (2 * c + (n / c - 1) * c) * (n / c) / 2;
    }

    cout << result << '\n';
    return 0;
}
