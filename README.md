Struct embedding lets one struct include another struct directly. The embedded struct's fields can be accessed from the outer struct.

Write a function that finds the top active seller.

Create an Account struct with:

Name
Active

Create a Seller struct that embeds Account and also has:

Sales
Rating

The function topSeller should accept a slice of Seller values.

Each seller's score is:

score = Sales * Rating


Return the name of the active seller with the highest score.

TEST ---->  go run .
Ignore sellers where Active is false.

If two active sellers have the same score, return the one that appears first.

If there are no active sellers, return an empty string.

