#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>



int random_number(int max) 
{
    int x;
    x = rand() % 100 + 1;
    return x;
}



int main() 
{
    int a = random_number(10);
    printf("The number is %d\n", a);
    return 0;
}

