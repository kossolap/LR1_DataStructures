#include "queue.h"

Queue* createQueue() {
    Queue* q = new Queue;
    q->head = nullptr;
    q->tail = nullptr;
    return q;
}

void enqueue(Queue* q, string value) {
    Node* new_node = new Node{value, nullptr};
    if (q->tail == nullptr) { q->head = q->tail = new_node; return; }
    q->tail->next = new_node;
    q->tail = new_node;
}

void dequeue(Queue* q) {
    if (q->head == nullptr) return;
    Node* temp = q->head;
    q->head = q->head->next;
    if (q->head == nullptr) q->tail = nullptr;
    delete temp;
}

void printQueue(Queue* q) {
    Node* current = q->head;
    cout << "[ ";
    while (current != nullptr) {
        cout << current->data << " ";
        current = current->next;
    }
    cout << "]" << endl;
}
