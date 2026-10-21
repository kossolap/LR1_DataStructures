#include <iostream>
#include <fstream>
#include <sstream>
#include <map>
#include <vector>
#include <string>
#include "array.h"
#include "list.h"
#include "dlist.h"
#include "queue.h"
#include "stack.h"
#include "tree.h"

using namespace std;

enum StructType { ARRAY, LIST, DOUBLE_LIST, QUEUE, STACK, TREE };

struct Structure {
    StructType type;
    void* data;
};

map<string, Structure> registr;

vector<string> parseQuery(const string& query) {
    vector<string> tokens;
    istringstream iss(query);
    string token;
    while (iss >> token) tokens.push_back(token);
    return tokens;
}

bool createStruct(const string& name, char typeCode) {
    if (typeCode == 'S') registr[name] = {STACK, createStack()};
    else if (typeCode == 'Q') registr[name] = {QUEUE, createQueue()};
    else if (typeCode == 'F') registr[name] = {LIST, createList()};
    else if (typeCode == 'L') registr[name] = {DOUBLE_LIST, createDList()};
    else if (typeCode == 'M') registr[name] = {ARRAY, createArray(10)};
    else if (typeCode == 'T') registr[name] = {TREE, createTree()};
    else return false;
    return true;
}

bool executeCmd(const string& query) {
    vector<string> tokens = parseQuery(query);
    if (tokens.empty()) return false;

    string cmd = tokens[0];
    if (tokens.size() < 2) return false;
    string name = tokens[1];
    
    char typeCode = cmd[0];
    if (cmd == "ISMEMBER") typeCode = 'T';
    
    bool exists = (registr.find(name) != registr.end());
    if (!exists) {
        if (cmd == "MPUSH" || cmd == "FPUSH" || cmd == "LPUSH" || cmd == "QPUSH" || cmd == "SPUSH" || cmd == "TINSERT" || cmd == "MPUSHIDX") {
            if (!createStruct(name, typeCode)) return false;
            exists = true;
        } else {
            cout << "error: structure not found." << endl;
            return false;
        }
    }

    Structure& str = registr[name];

  
    if (cmd == "MPUSH") {
        MyArray* arr = (MyArray*)str.data;
        if (tokens.size() == 3) pushBack(arr, tokens[2]);
        else if (tokens.size() == 4) pushIndex(arr, stoi(tokens[2]), tokens[3]);
        return true;
    }
    if (cmd == "MPUSHIDX") {
        MyArray* arr = (MyArray*)str.data;
        if (tokens.size() == 4) pushIndex(arr, stoi(tokens[2]), tokens[3]);
        return true;
    }
    if (cmd == "MGET") {
        MyArray* arr = (MyArray*)str.data;
        if (tokens.size() == 2) printArray(arr);
        else if (tokens.size() == 3) cout << getElement(arr, stoi(tokens[2])) << endl;
        return true;
    }
    if (cmd == "MDEL") { deleteElement((MyArray*)str.data, stoi(tokens[2])); return true; }
    if (cmd == "MREPLACE") { replaceElement((MyArray*)str.data, stoi(tokens[2]), tokens[3]); return true; }
    if (cmd == "MLEN") { cout << getLength((MyArray*)str.data) << endl; return true; }

   
    if (cmd == "FPUSH") {
        MyList* list = (MyList*)str.data;
        if (tokens[2] == "front") pushFront(list, tokens[3]);
        else if (tokens[2] == "back") pushBack(list, tokens[3]);
        else if (tokens[2] == "before") pushBefore(list, tokens[3], tokens[4]);
        else if (tokens[2] == "after") pushAfter(list, tokens[3], tokens[4]);
        return true;
    }
    if (cmd == "FDEL") {
        MyList* list = (MyList*)str.data;
        if (tokens.size() == 3) deleteByValue(list, tokens[2]);
        else if (tokens[2] == "before") popBefore(list, tokens[3]);
        else if (tokens[2] == "after") popAfter(list, tokens[3]);
        return true;
    }
    if (cmd == "FPOPFRONT") { popFront((MyList*)str.data); return true; }
    if (cmd == "FPOPBACK") { popBack((MyList*)str.data); return true; }
    if (cmd == "FGET") {
        MyList* list = (MyList*)str.data;
        if (tokens.size() == 3) cout << (findNode(list, tokens[2]) ? "Найден" : "Не найден") << endl;
        return true;
    }

    
    if (cmd == "LPUSH") {
        DoublyList* dl = (DoublyList*)str.data;
        if (tokens[2] == "front") pushFrontD(dl, tokens[3]);
        else if (tokens[2] == "back") pushBackD(dl, tokens[3]);
        else if (tokens[2] == "before") pushBeforeD(dl, tokens[3], tokens[4]);
        else if (tokens[2] == "after") pushAfterD(dl, tokens[3], tokens[4]);
        return true;
    }
    if (cmd == "LDEL") {
        DoublyList* dl = (DoublyList*)str.data;
        if (tokens.size() == 3) deleteByValueD(dl, tokens[2]);
        else if (tokens[2] == "before") popBeforeD(dl, tokens[3]);
        else if (tokens[2] == "after") popAfterD(dl, tokens[3]);
        return true;
    }
    if (cmd == "LPOPFRONT") { popFrontD((DoublyList*)str.data); return true; }
    if (cmd == "LPOPBACK") { popBackD((DoublyList*)str.data); return true; }

    
    if (cmd == "QPUSH") { enqueue((Queue*)str.data, tokens[2]); return true; }
    if (cmd == "QPOP") { dequeue((Queue*)str.data); return true; }

  
    if (cmd == "SPUSH") { push((Stack*)str.data, tokens[2]); return true; }
    if (cmd == "SPOP") { pop((Stack*)str.data); return true; }

   
    if (cmd == "TINSERT") { AVLTree* t = (AVLTree*)str.data; t->root = insert(t->root, tokens[2]); return true; }
    if (cmd == "TDEL") { AVLTree* t = (AVLTree*)str.data; t->root = deleteNode(t->root, tokens[2]); return true; }
    if (cmd == "TGET") { AVLTree* t = (AVLTree*)str.data; cout << (search(t->root, tokens[2]) ? "Найден" : "Не найден") << endl; return true; }
    if (cmd == "ISMEMBER") { AVLTree* t = (AVLTree*)str.data; cout << (search(t->root, tokens[2]) ? "TRUE" : "FALSE") << endl; return true; }

    
    if (cmd == "PRINT") {
        string target = tokens[2];
        if (target == "array") printArray((MyArray*)str.data);
        else if (target == "list") printList((MyList*)str.data);
        else if (target == "dlist") printDList((DoublyList*)str.data);
        else if (target == "queue") printQueue((Queue*)str.data);
        else if (target == "stack") printStack((Stack*)str.data);
        else if (target == "tree") { AVLTree* t = (AVLTree*)str.data; printTree(t->root, tokens[3]); cout << endl; }
        return true;
    }
    return false;
}

bool loadFromFile(const string& filename) {
    ifstream fs(filename);
    if (!fs.is_open()) return false;
    string line;
    while (getline(fs, line)) {
        if (!line.empty()) executeCmd(line);
    }
    fs.close();
    return true;
}

int main(int argc, char* argv[]) {
    string filename = "file.data";
    string query = "";

    for (int i = 1; i < argc; i++) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) filename = argv[++i];
        else if (arg == "--query" && i + 1 < argc) query = argv[++i];
    }

    loadFromFile(filename);

    if (!query.empty()) {
        if (executeCmd(query)) {
            if (query.find("PUSH") != string::npos || 
                query.find("DEL") != string::npos || 
                query.find("POP") != string::npos || 
                query.find("INSERT") != string::npos ||
                query.find("REPLACE") != string::npos) {
                ofstream out(filename, ios::app);
                out << query << endl;
                out.close();
            }
        }
    } else {
        cout << "no query provided." << endl;
    }

    return 0;
}
