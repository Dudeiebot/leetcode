  to me for medium


### Group Anagram
- Initialize a map where the key is an array of 26 integers (one for each letter) and the value is a slice of strings (the group of anagrams)
- range through the strings
- Initialize an array of 26 integers to count the frequency of each character in the string. so therefore k = [0,0,0,0,0,0,0...........0] for intialization.
- range through each character in the string
- Update the frequency count for the character in the array, so for example if char = e then k[e-a]  a = 0 and e = 4 our int calculation, so k[4] will be incremented from 0 to 1, therefore k = [0,0,0,0,1,0,0........0] and we check for each character in the string like that till we are done with the string
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
