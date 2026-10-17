#include <stdio.h>



typedef struct {
    char *name;
    char *phone;
} Person;


int main(void)
{
    Person contacts[3];
    // names
    contacts[0].name = "John";
    contacts[1].name = "Bob";
    contacts[2].name = "James";
    // phones
    contacts[0].phone = "123";
    contacts[1].phone = "321";
    contacts[2].phone = "456";
    printf("something\n");


}
