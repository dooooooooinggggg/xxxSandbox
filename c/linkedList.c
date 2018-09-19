#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Linked List
struct listStruct
{
    char value[20];
    struct listStruct *next;
};

void printLinkedList(struct listStruct *thisList)
{
    printf("this value is [%s]\n", thisList->value);
    if (thisList->next == NULL)
        return;
    return printLinkedList(thisList->next);
}

void addValueToList(struct listStruct *linkedList, char value[20])
{
    if (linkedList->next != NULL)
        return addValueToList(linkedList->next, value);

    // それ以降
    struct listStruct *thisList = (struct listStruct *)malloc(sizeof(struct listStruct));
    strcpy(thisList->value, value);
    thisList->next = NULL;
    linkedList->next = thisList;
    return;
}

void initList(struct listStruct *linkedList, char value[20])
{
    strcpy(linkedList->value, value);
    linkedList->next = NULL;
    return;
}

int main()
{
    struct listStruct *linkedList;
    initList(linkedList, "aa");
    addValueToList(linkedList, "bb");
    addValueToList(linkedList, "cc");
    printLinkedList(linkedList);
    return 1;
}
