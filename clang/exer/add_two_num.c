#include <stdio.h>


int main(void)
{
    int num1;
    int num2;

    printf("Number 1: ");
    scanf("%d", &num1);
    printf("Number 2: ");
    scanf("%d", &num2);

    int result = num1 + num2;

    printf("%d + %d = %d\n", num1, num2, result);
}

