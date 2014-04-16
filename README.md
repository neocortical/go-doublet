go-doublet
==========

Doublets
--------

"Doublets" is a word puzzle invented by Lewis Carroll in 1877. Also known as "word ladders", "word-links", or "word golf", playing doublets consists of transforming one word into another via a series of intermediary words. Intermediary words are formed by changing a single letter of the previous word. Thus, each intermediary step differs from its previous and subsequent steps by a single letter. The solution to a doublet is a linked list of valid dictionary words beginning with the start word and terminating the end word of the puzzle.

Ex: dog --> god: dog --> dot --> got --> god

Interest
--------

I started thinking about the game when I ran across an software engineering interview question that asks the candidate to design an algorithm to solve doublets, given a dictionary of English words. It's a very good technical interview question. There is a brute force solution that is relatively simple to implement but has horrible time and space complexity. Then there are various more challenging ways to approach the problem. Lastly, any sophisticated solution will draw upon an applicant's academic knowledge of computer science.

Solution
--------

I had coded up a decent recursive solution when I had the realization that the dictionary could be transformed into a graph of all valid word transitions. Once the graph was generated, a shortest path algorithm could be used to find the smallest number of steps between any two inputs. Dijkstra's algorithm can be used, but A* is a better choice. A* is a variant of Dijkstra's that uses a programmer-chosen heuristic to estimate the path length between any vertex and the solution vertex. Using an estimation function of "minimum number of moves (legal OR illegal) to reach the target word" means that paths that are converging on the solution will always be evaluated before paths that are diverging. This guarantees that A* will find the shortest path when any path exists.

Code
----

The code is written in the Go programming language. I wrote it originally in Java, but I'm working on my Go proficiency, so I wanted to write a Go version. This was interesting, because Go made some things much simpler syntactically (creating the graph) while other things were tricky to translate (working with priority queues in Go feels clunky to my OOP mindset). 

I can't say that I followed every Go idiom as efficiently as possible, but the program came out pretty simple and clear, and it's blazing fast. The path from "orange" to "yellow" is 33 steps using my dictionary and results in the inspection of more than 10,000 vertices. This takes around 15 ms in my Go version, whereas my Java version takes more than 50 ms. Simple doublets, such as "dog" to "god", are measured in microseconds. 

Install
-------

go get github.com/neocortical/go-doublet

Run
---

$ go-doublet \<dicionary_file.txt\>

The dictionary file should be a plain text file containing one word per line. Invalid words (those containing non-alpha characters) will be ignored. There is a decent dictionary file included under the res/ directory.

Notes
-----

Because the program takes a dictionary file as an argument, users can use their own dictionaries. So for example, you could solve doublets using only the words found in Crime and Punishment if you wanted to. 

One bummer is that the algorithm chooses paths arbitrarily when multiple paths of the same length exist. I think a nice addition would be to preferentially weight common words, so that solutions would tend to be less esoteric. 

Links
-----

The A* search algorithm: http://en.wikipedia.org/wiki/A*_search_algorithm

The dictionary I used: http://www-01.sil.org/linguistics/wordlists/english/

Interesting blog post on graphing doublets using Wolfram Alpha/Mathematica: http://blog.wolfram.com/2012/01/11/the-longest-word-ladder-puzzle-ever/

