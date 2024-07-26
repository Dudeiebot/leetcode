to me for medium

### Group Anagram

- Initialize a map where the key is an array of 26 integers (one for each letter) and the value is a slice of strings (the group of anagrams)
- range through the strings
- Initialize an array of 26 integers to count the frequency of each character in the string. so therefore k = [0,0,0,0,0,0,0...........0] for intialization.
- range through each character in the string
- Update the frequency count for the character in the array, so for example if char = e then k[e-a] a = 0 and e = 4 our int calculation, so k[4] will be incremented from 0 to 1, therefore k = [0,0,0,0,1,0,0........0] and we check for each character in the string like that till we are done with the string
- Use the character frequency array as a key in the map and append the string to the corresponding group. after done with each string, we make k the key of our map and the string the value as intitalized. like this hmap[[26]int{1, 0, 0, ..., 1, ..., 1, 0, 0}] = ["eat"]. we do that for all the word in the array string
- we have intialized a slice that holds our result
- Iterate over the map and append each group of anagrams to the result slice and remember each key in a hashmap must be unique, if you try to insert a key that already exists in the map, the new value will overwrite the existing value for that key. so we assign each value to a slice.

### Top K Frequent

- so here we are checking an array and checking the most Frequent k element
- we can use 2 approach but first we are going to be creating a map that will store the occurence of each num
- the first method and second method uses this process, by creating a double slice that have the size of len(nums)+1, then we go on by storing the occurence as the inner loop with the nos as the array slice position, for example: [1,1,1,2,4,4] is our array then our map will be [1:3, 2:1, 4:2] then our double slice will be [[], [2], [4], [1], [], [], [], []] which means that double_slice[3] will take our input 1 and double_slice[1] will take our input 2.
- when done, the first method after our double slice will range through our double_slice from the back, if double_slice[i] is not equal to nil, we append all the data in our double_slice to the result array and then check if len of result is equal to or greater than k, if it is the later we take only the first k element of the array
- the second method, is to range through the double_slice from the back, then we range through the double_slice[i] element, check if k > 0 and if not we append value from the double_slice[i] element to the array result and decrement out k value

### product Except Self

- we are returning an array, which is the product of our input array without the array position in relation to the result array, so there are some barriers in this array because we are not suppose to use division at all in our solution.
- so we are going to be using a prefix and suffix solution here, prefix will find the products from the front and suffix with do the product from the back
- prefix solution will go like this, when we range through the nums array, we will make sure to equalize our result position(res[i]) to the prefix and the prefix will always be Updated with new figures because we are multiplying it by the nums position(nums[i])
- suffix will eqaully go like this from the end, but here instead our result position is being updated with the product of our suffix, and the suffix is being updated with the nums position
- every variables get updated upon each iteration, and also i noticed that you dont need to get to the last and first array when performing both prefix and suffix operation

### longest Consecutive

- we can use slices.Sort to sort this and check the longest Consecutive there but the question want us to solve this in o(n) instead of introducing the sorting algorithm that uses o(logn)
- so for our o(n) operation, we are going through the nums and making sure that a number does not occur more than one, so we are using the normal map for this
- so right now the problem we have is that our map is not ordered, we range through the nums again and we check if n-1 is not in the map, if it is not we assign our currentLength to 1 and assign an var current to n
- so now we are working with current and currentLength, we go on further to check if current(our n) + 1 is in the map(if it is not we go to the next num but if it is we increment the current by 1) and also currentLength by 1
- when we are done with all the number, remember that upon each number that does not have it Consecutive nos in the array, we are going to be reasssigning our currentLength to 1, upon getting the longest Consecutive we are going to be comparing it to the maxLength and equating it to the maxLength (if the maxLength is less than currentLength)
- finally we return maxLength

### Two Sum II

- so this is a medium and also a follow up question to the easy question (2 sum) and here our array is sorted
- binary search will be used here to solve the problem and for l < r{}
- the algorithm is that on addition of nums in l and r, we check if it is greater than or less than our target
- so if nums[l] + nums[r] is greater than target, r is decremented and if nums[l]+ nums[r] is less than target, l is incremented and if we get our answer, we store it in an array and return it

### ThreeSum

- three sums with our idea from two sums II, we are going to be sorting the algorithm (using the slices.Sort)
- range through the nums with index(i) and each num(n), checking if i is greater than 0 and n is not equal to the nos before nos(i-1)
- upon introducing binary search, we then intitalize our threeSum which is the addition of n + nums[l] + nums[r] and if threeSum is greater than 0(our target) we decrement r, and if it is lessthan 0 we increment l
- else we append our result to the double slice created, and remember we need multiple instances where our threeSum is equal to target and not only one possible scenario
- When we are done with the first result already appended, we go back to the array by incrementing l and checking if nums[l] is equal to nums[l-1] and l < r for more results

### MaxArea

- Binary search is going to be used here
- while l is less than r, we are going to intitalize our area as l-r \* min(nums[l], nums[r])
- we keep on updating the max value with the our intitalized res, and checking if nums[l] is less than nums[r], we increment l and we decrement r if not

### Length Of Longest Substring

- so here we are using the sliding window techniques(double pointer as in the case of binary search)
- so both left and right start at the beginning but the right is being incremented till the end of the string and they are stored in a map
- When a letter appear more than once, we decrease the value from the map and increment the left
- our result is being stored as the max value in comparison with (right-left+1) which is our current position in the string

### Find Min in a rotated sorted array

- so here we are using Binary search, and upon intialization we check if l is less than right
- we define our middle immediately after the for loop
- we check for difference, if left is greater than right, left is updated to middle + 1
- and else that means if left is less than right, we equate right to middle
- upon all our min num will end at the l, then we return it

### searching a particular number in a rotated sorted array

- we have our rotated array here and we are going through it by searching a certain target in it
- we are using binary search also here because we are to submit it in Ologn
- we define our left and right, initiate our loop and define our middle
- so basically it is all suppose to come down to the middle, and remember we check at the beginning of our loop
- so here our approach is that we want to divide the array into left and right and how do we check intiallialy if we are meant to use the left or the right
- if l is lessthan or equal to middle that means we should check the left because it is sorted but if it is the other way round, we are checking the right because that means our right is being sorted
- for the left part, we are checking with our target next, if our target is greater than left and our target is less than middle, we update our right as middle + 1 else our left is updated as middle - 1
- for the right part, we are checking with our target, if target is greater than middle and target is less than right, we update our left as middle + 1 else the right as middle - 1
- at the end if the target is not in the array, we return -1

### Re-Order List

- we are going to be re-ordering a list in this format.
  ![image](../medium/static/list_rep.png)

- so in this question we are not returning anything but only our input back
- we check for all base cases and return nothing also
- we are using a fast and slow pointer(floyd tortoise and hare method), the slow pointer move at `1*` speed and the fast pointer move at `2*` speed
- so here we are trying to get the middle of our listNode and it our slow that will help use to get there and remember the fast is moving at `2*` pace so we should remember to make sure our fast get to null because of odd lens
- we swap values upon getting the second half of the listNode
- we then go on to merge the two halves of the listNode
- The two halves of the list are merged. first traverses the first half, and prev (now the head of the reversed second half) traverses the reversed second half
- For each node in the first half, the next node is set to the corresponding node from the reversed second half, effectively weaving the two halves together.

##### to Illustrate

consider the list: `1 -> 2 -> 3 -> 4 -> 5`
Finding the Middle:

- slow points to 3 (middle of the list)
  Reversing the Second Half:
- The second half `4 -> 5` is reversed to `5 -> 4`
  Merging the Two Halves:
- The list is merged as `1 -> 5 -> 2 -> 4 -> 3`
