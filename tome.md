
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
