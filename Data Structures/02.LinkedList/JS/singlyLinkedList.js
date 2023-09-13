class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}

class LinkedList {
  constructor(element) {
    this.headNode = null;
    this.tailNode = null;
    this.length = 0;
  }
  initiateNodeAndIndex() {
    return { currentNode: this.headNode, currentIndex: 0 };
  }

  size() {
    return this.size;
  }
  head() {
    return this.head;
  }
  tail() {
    return this.tail;
  }
  isEmpty() {
    return this.length === 0;
  }
  addFirst(element) {
    const node = new Node(element);
    if (this.headNode === null) {
      this.tailNode = node;
    }
    node.next = this.headNode;
    this.headNode = node;
    this.length++;
    return this.size();
  }
  addLast(element) {
    if (this.headNode === null) {
      return this.addFirst(element);
    }
    const node = new Node(element);
    this.tailNode.next = node;
    this.tailNode = node;
    this.length++;
    return this.size();
  }
  addAt(index, element) {
    if (index > this.length || index < 0) {
      return console.log("Out of Range index");
    }
    if (index === 0) return this.addFirst(element);
    if (index === this.length) return this.addLast(element);
    let { currentIndex, currentNode } = this.initiateNodeAndIndex();
    const node = new Node(element);

    while (currentIndex !== index - 1) {
      currentIndex++;
      currentNode = currentNode.next;
    }

    // Adding the node at specified index
    const tempNode = currentNode.next;
    currentNode.next = node;
    node.next = tempNode;
    // Incrementing the length
    this.length++;
    return this.size();
  }

  removeFirst() {
    if (this.headNode === null) {
      return null;
    }
    let removedNode = this.headNode;
    this.headNode = this.headNode.next;
    this.length--;
    if (this.isEmpty()) {
      this.tailNode = null;
    }
    return removedNode?.data;
  }

  removeLast() {
    if (this.isEmpty()) return null;
    if (this.length === 1) return this.removeFirst();
    let removedNode = this.tailNode;
    let { currentNode } = this.initiateNodeAndIndex();
    while (currentNode.next.next) {
      currentNode = currentNode.next;
    }
    currentNode.next = null;
    this.tailNode = currentNode;
    this.length--;
    return removedNode?.data;
  }

  remove(element) {
    if (this.isEmpty()) return null;
    let { currentNode } = this.initiateNodeAndIndex();
    let removedNode = null;
    if (currentNode.data === element) {
      return this.removeFirst();
    }
    if (this.tailNode.data === element) {
      return this.removeLast();
    }
    while (currentNode.next) {
      if (currentNode.next.data === element) {
        removedNode = currentNode.next;
        currentNode.next = currentNode.next.next;
        this.length--;
        return removedNode.data;
      }
      currentNode = currentNode.next;
    }
    return removedNode?.data || null;
  }

  removeAt(index) {
    if (index < 0 || index >= this.length) {
      return console.log("Index out of bounds");
    }
    if (index === 0) return this.removeFirst();
    if (index === this.length - 1) return this.removeLast();
    let { currentIndex, currentNode } = this.initiateNodeAndIndex();
    while (currentIndex !== index - 1) {
      currentIndex++;
      currentNode = currentNode.next;
    }
    let removedNode = currentNode.next;
    currentNode.next = currentNode.next.next;
    this.length--;
    return removedNode.data;
  }

  log() {
    console.log(JSON.stringify(this.headNode, null, 2));
  }
}

let ll = new LinkedList();
ll.addFirst(11);
ll.addAt(1, 2);
ll.addFirst(10);
ll.addFirst(10);
ll.addAt(1, [99, 88, 77]);
ll.log();
