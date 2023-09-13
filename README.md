DS: Way of organizing data so that it can be accessed/queried/updated quickly efficently.

Abstract Data Type: ADT is abstraction of DS which only provides the interface to which data structure must adhere to.

ADT(Abstraction)			| 			DS(Implementation)

List 					|			Dynamic Array, Linked List 
Queue					|			Linked List,Array,Stack based on Queue
Map 					| 			Tree Map, Hash Map


-- Computational Complexity Analysis.
How much time? How much space?

- Big(O) -> Upper bound of worst case, measure performance as increase in input size.

 -Constant Time: 	O(1)
 -Logarathmic Time:	O(log(n))
 -Linear Time:		O(n)
 -Linearithmic Time:	O(nlog(n))
 -Quadric Time:		O(n^2)
 -Cubic Time:		O(n^3)
 -Exponential Time:	O(b^n)   b>1 
 -Factorial Time:	O(n!)



- Static & Dynamic Array
Static Array: Fixed Length containter contanting n elements indexable from range [0, n-1]


		Static Array |  Dynamic Array
Access		O(1)		O(1)
Search		O(n)		O(n)
Insertion	N/A 		O(n)
Appending	N/A 		O(1)
Deletion 	N/A 		O(n)



- Singly & Doubly Linked List 

Linked List: Sequential list of nodes htat hold data which point to other nodes containing data.
Used while creating List Queue & Stack Implementatin.
Great for creating circular list.

Head, Tail, Node, Pointer

 

 
