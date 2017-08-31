#include<iostream>
using namespace std;

int main()
{
    int a, b, x, c;
    int prev;
    prev = 0;
    while(cin >> a)
    {
        cin >> b;
        x = 1;
        prev = 1;
        for (int i=a; i <= b; i++)
        {
            c = i;

            while (c > 1)
            {
                if (c % 2 == 1)
                {
                    c = c * 3 + 1;
                    x++;
                }
                else
                {
                    c = c / 2;
                    x++;
                }

            }
            if(x > prev)
            {
                prev = x;
            }
            x = 0;
        }

        cout << a << " " << b << " " << prev + 1 << endl;
    }
}
