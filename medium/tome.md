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
