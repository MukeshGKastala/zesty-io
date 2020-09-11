# Autocomplete prefixes against Shakespeare's complete works

This approach uses a Trie and a Priority Queue

## How it works

1. Setup a buffered I/O Reader to scan for **words**
1. Discarded words w/ nonalphabetic characters
1. Split words into a list of characters
1. Insert characters into Trie tracking frequency until all valid words are consumed
1. Start HTTP server and wait for GET
1. After successful request validation, get prefix's sub-Trie
1. To find up to 25 of the sub-Trie's most frequent words, firstly initialize a Priority Queue whose comparator is frequency 
1. Preorder traverse the sub-Trie and insert into the Queue when node's frequency is greater than 0(denoting end-of-word)
1. Finally return at most 25 from the Priority Queue

## Requirements

- [x] HTTP GET request w/ **prefix** as a URL parameter 
- [x] Ignore non-word entities
- [x] Top 25 results ordered by frequency 
- [x] Only use the Golang standard library

## Getting Started

1. Clone the repo
1. Install Golang: `brew update && brew install golang`
1. Run the service with `make run`

## Running Unit Tests

After cloning the repo and installing Golang, run the test scripts: `make test`

## Supported HTTP Requests

**GET `/autocomplete?prefix=<prefix>`**

Response:
```
{
  "words": []string,
}
```

This gets the 25 most frequent words starting with the prefix

## Example CURL Command

Once the server is running:

```shell
curl -X GET http://localhost:9000/autocomplete\?prefix\=th
```

Default port of `9000` can be modified in the makefile

## Example Prefixe Results

sample data: `["th", "fr", "pi", "sh", "wu", "ar", "il", "ne", "se", "pl"]`

My findings:
<img src="/images/prefix_examples.png">

To follow along, run `make examples` which will call a shell script and display results to *stdout*

