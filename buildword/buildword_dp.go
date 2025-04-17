package buildword

func BuildWordDP(word string, fragments []string) int {
	n := len(word) // Get the length of the word
	if n == 0 {    // If the word is empty, then we can return 0
		return 0
	}

	// I created the dp array to store the minimum number of fragments needed to build the prefix word
	dp := make([]int, n+1)

	// Initialize the dp array with a large value (n + 1) to represent infinity
	for i := 1; i <= n; i++ {
		dp[i] = n + 1 // Set the initial value to n + 1 (infinity)
	}

	// For the base case, we can build an empty string with 0 fragments
	for i := 1; i <= n; i++ {
		for _, fragment := range fragments { // Iterate through each fragment
			fragLen := len(fragment)

			if i >= fragLen { // Does the fragment fit in the current position?
				prevIndex := i - fragLen // Calculate the previous index

				if dp[prevIndex] <= n && word[prevIndex:i] == fragment { // Check if the fragment matches the substring of the word
					if dp[prevIndex]+1 < dp[i] { // If we can build the word with fewer fragments, update dp[i]
						dp[i] = dp[prevIndex] + 1
					}
				}
			}
		}
	}

	if dp[n] > n { // If dp[n] is still n + 1, it means we couldn't build the word with the fragments
		return 0
	}
	return dp[n] // Return the minimum number of fragments needed to build the word
}
