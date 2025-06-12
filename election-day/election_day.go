package electionday

import "strconv"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	res := initialVotes
	return &res
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter != nil {
		*counter += increment
	}
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	res := new(ElectionResult)
	res.Name = candidateName
	res.Votes = votes
	return res
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	if result == nil {
		return ""
	}
	return result.Name + " (" + strconv.Itoa(result.Votes) + ")"
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	if votes, ok := results[candidate]; ok {
		if votes > 0 {
			results[candidate] -= 1
		}
	}
}
