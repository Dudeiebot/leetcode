
Just tome.md for my all everything into


So here basically we have the algorithm for leetcode and some explainantion following them



- Duplicate Number (we use map to get things done here or some len shii by ranging twice but memeory shii can come in)
- Valid anagram (we use map also and performs increment on it, then we range through the key and value of the increment)
- Valid palindrome (a very kind of strange lens calculation that goes through the string)
- Single number apperance (Using the bitwise operator will give us the constant time and space needed and very fast also, the bitwise operator here have a kind of calculation in it and it ir represented with the (^) symbol)
- Two Sums in a array and returning the indices of the nos, here we can use brute force which result to O(n2) and also we can use our hashmap to solve it which will make it faster and help save time O(n)
- Finding the nos palindrome using the equation half = half * 10 + x % 10, and then remember it is a loop, so it get to go through all till x > half then we check for the equality in the palindrome or if it is an odd nos we check for the mod with the equality
- converting romans to integer and here we have two approach to it, we loop to check the length of our input string and then we store the result to a map, the other way is replacing the bizzare numerals like (4, 9, 40, 90, 400 and 900) and then we keep on replacing them in instances where they come up
- Here we get our longest prefix, firstly we range through the first word, then range for other words in the remaining slice, if i is equal to length of word or words[i] is not equal to firstlet[i], it then return the firstlet from 0 to i where it stop. We are returning firstlet at the end because it may be the longest prefix also from the beginning.
- checking parenthesis, we are using the LIFO approach of a stack and we range through our input, then we have a big if statement that check for opening bracket first and append it to the stack then we pop out the closing bracket and compare it with the pairs map(or the opening brackets). return false if they are different or return true
- merging two linked sorted list, we create a temp node which is a ref to listNode, we rep current node as the tempNode. we compare between the values and we bring it up to the tempNode through the help of currentNode
- Removing duplicates in a sorted array, here we have to go through the array and use a pointer to make our changes and perform it in constant space of time. the pointer is being incremented when they are not the same and also the right value is moved to the left value.
- Removing the duplicate element of val x in an array, it is the same like the later just that here we are given a val and checking for the duplicates in the array, firstly we create pur pointer and go through the nums and we check if each num is not the same with the value provided, then we move the right to the left if they are not and they increment our pointer, which we later return 
- Removing elements and removing duplicates elements are the same because it uses the same pointer approach with increment on it, for constant space of time 
- Checking the index in which a substring is in a mainstring, we can use the fast route with the gi library strings.Index and then we can use a pointer that point to the end(but not directly the end cause that means we are wasting time), then we use the array indexing library to do the equality(which is this i: i+len(substring))
- Searching for insert in a array and we are returning the index if it is in the array or where it is supposed to be in the array, We are using the O(log n) here which is the binary search (it goes through the middle and does the search from there) 
- we are returning the legth of the last word in a sentences, so firstly take away all the trailing spaces then we count back the last word
- We are incrementing the last word in an array and when the last array is 9 we do the normal addition, we dont forget that we need to start from the last unit. We start our for loop from the end and perform our addition.
- We are adding two binary nos together, so we will just use the normal addition of two binary nos(not greater than 2)
- mySqrt used to get the square of a number without using the standard library, we use binary search and check from the middle, the we multiply it twice to see if it is right
- Dynamic programming which is related to the fibonacci series, we have a nos of ways to climb a step. If there are no step at all there will be only one way(not to climb at all). so the no(ways) == no(ways) - 1 + no(ways) - 2
- removing duplicates from a linked list, we assign a pointer to our list and it get updated
- Merging 2 sorted array of length (m + n) and length (n), and the final sorted ahould not be returned but instead be sorted in nums1, so here we basically swap values and compare between the 2 arrays
- Inorder traversal is one of the DFS techniques that use 0(n) time complexity, which traverse the left node first to the root and the right node last. I am using the iterative techniques here we create a stack to store our current elements, the stack append when the current is not nil and pop out when the current is nil 
- PreOrder traversal, here we put the root first followed by the left and right, we can also use iterative and recursion method also. 
- PostOrder traversal, here we put the root Val last. We first traverse the left and also the right and finally the root is added last, we can use both iterative and recursion method here also
- IsSame Tree, we check for the base case and if the root (p.Val, q.Val) are not equal then we check for same left and same right
- IsSymmetry(checking if the tree is mirror like), here we create a helper func dfs that check the difference between the left and the right node because we cant use the root node alone to check if they are mirror like. This func is effective (left.Right, right.Left && left.Left, right.Right) which check for opposite sides of the tree
- for finding the maximum depth in a tree, we perform a recursion to the left and right of the tree, then add 1 to the max of the height(root.Left && root.Right). Always remember the base case for all this algorithm or you can get it wrong there
- converting a sorted array to a binary tree is straight forward with recursion and that is by dividing the array into 2, and representing the root as the middle, then follows the remaining value of root.right and root.left
- checking for balanced tree, we go create 3 helper func max, abs and dfs. the max func check between our max height on both the left and right, the abs check for the diff between the left and the right and the dfs is use to introduce the new int we added to the program 
- Minimum height, we create a helper func that check for the Minimum values and also when all the remaining value are nil, it give us the second output as answer then we add it to the root
- the path sum is checked from the deduction of the targetSum from the root.Val and then a recursion is performed on it till root.Left and root.Right become nil
- Generating Pascal triangle, we know the first array is 1 and the second is 1, 1 then we can also use that information in generating the remaining ones.
- Getting the row index of a pascal triangle can be easy with out pointer approach and appending the value to our slice
- The best time stock is can eaily be done with a two pointer(lef and right) approach techniques that goes from the beginning to the end
- max profit from a stock, we are also using the 2 pointer approach only that here there is no else in our code and thatt means the buy goes through loop across all 
- finding the single number among number that are 3 in an array, this is different from the first single nos among two same nos in an array, we can use the bruteforce and also the constant space of time method here. the xor operator for the constant space of time goes through the nums and create a 2 pointer which then equal ones ^= (num & ^twos) and same for the secodn pointer
- cycle in a linked list, we use the floyd tortiose and hare method, which means we have 2 pointer called slow and fast. Fast goes 2* faster than slow, we return true when fast and slow get to meet or false if they dont meet at all. Perfect illustration is a circular track 
- detecting the cycle in the linked list can be done with hashmap
- Getting Intersection Node can come in 2 ways using the hashmap, we save a ListNode on the map and we range the other one to check the equality. For good optimization and adequate memory we can create two pointer, where a and b will intersect then go through each other one by one and this is possible because a doesnot get to nil  and b also doesnot get to nil and when it does a becomes headB and b becomes headA and we return one of them when they are the same
- so this is a mathematics question, that have the equation and also some rune conversion
- the opposite of the conversion to string, this one the conversion is to int and they are not pretty hard just mathematics conversion and ranging through
- so here we are going to be using hashmap and boyer moore algorithm, they are both straight forward and easily understandable also
- reversing uint32 here, we have to do some bit manipulating and by doing this, we loop through our nos and for every loop we shift the result left once perform a bitwise operation and shift the num to the right once
- hamming weight is when we return the non-zero part of a number, and we are using unsigned int here so we check for till the num is not equal to 0 then we mod the last digit and move the num to the right once. Finally we return count
- is HAPPY alogrithm does some mathematical calculation to check if something pass or not, check leetcode for more understanding on it (Leetcode 202) 
- for removing elements in a linked list that are the same with our provided value, we just create 2 pointer and a dummy node that start before the beginning adn range through all. (dont worry about my util.ListNode, i created a package in the folder struct to avoid me writing the struct type all the time)
- isIsmorphic can be check with 2 different map back and forth just like the definition of ismorphic
- To reverse a list we can use 2 pointer for the iterative version, ann that is just creating a list outside and then swapping values(and remember curr.Next = prev and not curr.Val because it is an int)
- contains Nearby Duplicates can be done by creating 2 pointers or by using hashmap, either will work well but the hashmap is more memory efficient
