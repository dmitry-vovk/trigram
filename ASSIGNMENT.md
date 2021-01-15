# Go Developer Coding Challenge

In natural language processing, a _trigram_ is a sequence of three consecutive words in a given body of text. For example, the sentence "To be or not to be, that is the question" contains the following trigrams:

```
[to, be, or]
[be, or, not]
[or, not, to]
[not, to, be]
[to, be, that]
[be, that, is]
[that, is, the]
[is, the, question]
```

Given a series of trigrams, it is possible to generate a random piece of text that resembles the original. For example, if we start with the two words "to be", we can find two trigrams that match that prefix:

```
[to, be, or]
[to, be, that]
```

At this point, we can make a random choice, and end up with "to be that". We can repeat the process by taking the last two words, "be that", and looking up what words may come next. Although with such a small example we are very limited in our choice, given a large enough body of text, such as a novel, we can produce text that, while completely devoid of meaning, appears to be in the same style as the novel that it was trained on.

For example, after "learning" the text of some of Jane Austen’s novels, this technique produced the following paragraph:

> Middleton wished it very much inclined to ask them what they did not
> know how to take her from a long drive this morning; perhaps we may
> have as many sensations of exquisite felicity followed, and the walk
> must arise from something she said herself; and Jane’s offences rose
> again. They gave themselves up wholly to their satisfaction.

## Your task

We would like you to build a Go program that will "learn" from the text it receives, and generate random text using trigrams as described above.

Your program will expose an HTTP interface with two endpoints:

  1. `POST /learn` will "teach" your program about a body of text, sent as the POST body with `Content-Type: text/plain`.
  2. `GET /generate` will return randomly-generated text based on all the trigrams that have been learned since starting the program.

Here’s an example of how a client may teach the text of a novel, then generate some random text in that style. Note that if you’re using cURL to test the `POST /learn` endpoint, you will need to use `--data-binary` to ensure that newlines are retained.

```
$ curl --data-binary @pride-prejudice.txt localhost:8080/learn
$ curl localhost:8080/generate
To think it more than commonly anxious to get round to the preference of one, and offended by the other as politely and more cheerfully. Their visit afforded was produced by the lady with whom she almost looked up to the stables. They were to set out with such a woman.
```

## Requirements

In addition to the above, we would like you to keep the following requirements in mind:

  1. It is okay for your program to hold its state (the trigrams that it has learned) in memory, and to forget them when the process terminates.
  2. The random choice of words should be proportional to the frequency of that word in the trigrams. For example, if the word "stormy" follows the words "dark and" 9 times out of 10, then we would expect the randomly-generated text to preserve that frequency.
  3. You should expect the HTTP endpoints to be accessed concurrently. For example, you could imagine a Slack bot that sends every message it sees to the `POST /learn` endpoint, while constantly generating text using the `GET /generate` endpoint. Your program should be able to cope with concurrent access.
  4. Think about how your program would cope with learning a large amount of text. For example, if it was trained on every English language book on [Project Gutenberg](https://www.gutenberg.org), what would its memory usage look like? If you can think of different trade-offs that you could make here, let us know (even if you don’t end up implementing them).

## What we’d like to see

Although this is a relatively small task, we looking for production-ready code that's commented and tested as appropriate.

As well as sending us your code, we would like to hear about any design considerations or technical details you gave particular consideration to. If you had to make any particular trade-offs, or if you would have done things differently given more time, then let us know.
