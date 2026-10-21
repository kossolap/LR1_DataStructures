#pragma once
#include <iostream>
#include <string>
#include "list.h"
using namespace std;

struct Queue {
    Node* head;
    Node* tail;
};

Queue* createQueue();
void enqueue(Queue* q, string value);
void dequeue(Queue* q);
void printQueue(Queue* q);
