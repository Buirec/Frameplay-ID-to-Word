# 🤓 Frameplay Backend Assessment

A GoLang application that takes in member information for a theoretical organisation (Let's say Globex), uses the ID and transforms it to return a word!

This application is tied with the API also made by me so make sure it's running alongside [this](https://github.com/Buirec/Frameplay-Word-API)

## 👓 Overview

This application takes in data under the JSON structure where a member has the data of their name, ID, and home city, which then logs the ID and subsequently sends it to our API, which then sums the ID down to a single digit and picks the associated word, sends it back which we then log the result.

## 🔨 Usage

Run "make all" while running the associated API, it will use the data that is in data.json as the input.

## 🧠 Considerations

In a larger scale commercial application, I think I would make the decision to use either Ion or Avro instead of JSON, with the most likely candidate being Ion due to being the more flexible of the two, along with support for types such as Ion timestamps and blobs, since a regular application with a database of member information would have more fields than the ones we specified in our schema.

In a real commercial environment, there'd be a lot more security measures on the API side itself like hashing our member ID data on the application side and then unhashing it on the API side to ensure no one can retrieve our member data through common means like man in the middle attacks.

## 🚶‍♂️ Thought Process Walkthrough

The principal concern when I started this project was whether to use an existing API or to make my own, I've been working on random word games in my spare time so I definitely had word related processing on the mind leading to our application transforming data into a word. I decided to make my own API in this case to demonstrate my understanding of HTTP based APIs in Golang, as well as more control over what I wanted to app to actually do, rather than conform to the (pretty pitiful) set of word based APIs that I found online while working on said word games. Once the API was in place, it was really just making sure they adhered to the requirements of the specification, while making basic considerations like making sure there's no memory leaks, and that the code is performant and efficient.