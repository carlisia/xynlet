# Xynlet App
A program in Go that determines the smallest base (greater than or equal to 2) in which the first 1000 positive decimal integers are palindromes. Output is sent to a CSV file.

# External dependencies
None, only std lib has been used.

# Notes on commits
Normally I squash commits into a whole unit of work, usually having it correspond to a task or story. For this project, I chose to leave commits as smaller units. Also, longish commentaries are included on the body of commit messages as a (probably very poor!) attempt to give a sense of what it would've been like to pair-program with me during the code implementation.

# Opportunities for future improvements
It bothers me that on line 43 of the main.go file I create a new slice at each iteration. I would like to move the slice allocation outside of that loop and allocate the memory according to some sort of estimation. I also would like to experiment with changing from []int to []byte since I'm dealing with individual digits. This way I could leverage bytes.Buffer and probably make it more efficient. In making this change I would like to benchmark both versions and it will be interesting to see what the difference will be in terms of performance and also what version will read better.
