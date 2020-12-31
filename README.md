# Vocabulary Tracker

This is a simple tool that can help language learners like me to keep a record of words I've encountered and learned.<br>

It is built with GOlang and I tried to make it adhere to the "REST API" conventions.

# Features
The main purpose of this program is to organize words in two ways:

**1. The scenario in which one encounters a word.**<br>

For example, while reading the novel "Sherlock Holmes" by Arthur Conan Doyle, I learned the word **"page"** which refers to a boy who is serving a nobel person, along with  other words like **wagon**, **speckled**, etc.

Therefore, I can group those words together under a **tag** called "Sherlock Holmes".
![HomePage](https://github.com/Hackerry/Vocabulary_Tracker/blob/main/imgs/HomePage.png?raw=true)

**2. The same word itself in different scenarios.**<br>

For example, the word **page** refers to not only a paper but also a boy who is serving a nobel person.

I record those two different meanings of page that I've encountered as comment entries alongside the complete definition of "page".
![SearchPage](https://github.com/Hackerry/Vocabulary_Tracker/blob/main/imgs/SearchPage.png?raw=true)
![Entries](https://github.com/Hackerry/Vocabulary_Tracker/blob/main/imgs/Entries.png?raw=true)

# How to run
0. GOlang 1.14+ installed.

1. Clone this repository with `git clone`.

2. Run `./build.sh` to compile the executable.

3. Obtain an api_key from a thrid-party dictionary API (now only MERRIAM-WEBSTER'S COLLEGIATE® THESAURUS API supported: https://dictionaryapi.com/).<br>
   The non-commercial use API account has a limit of 1000 searches/day. You can also write your own parser to connect to other online dictionary APIs as well (see section below).
   
4. Replace the part in URL where the search word is with `[word]` placeholder. Fill in your key as well.<br>
   **i.e.** If my key is 12345, I'll get the query URL template as: https://www.dictionaryapi.com/api/v3/references/thesaurus/json/**[word]**?key=**12345**.
   
5. Put this URL into the root directory's **api_key** file, followed by a space and the text **Merriam-Webster**, which tells the program which API it is connecting to.<br>
   **i.e.** The final content should be like: `https://www.dictionaryapi.com/api/v3/references/thesaurus/json/[word]?key=12345 Merriam-Webster`.
   
6. Run the program with `./bin/Vocabulary_Hunter.exe`(on Windows) or `./bin/Vocabulary_Hunter`(on Linux).

7. Access the server in browser at `localhost:8080`.

# Connect to other APIs
To connect to other dictionary APIs, implement the JSON Parser first (i.e. example parser at `./api/Merriam_Webster_Json_Structs.go` and `./api/apiUtil.go`).<br>
Then modify the server to use the new API in `./main.go`.

# Features to add list
~~1. Display all words tagged by a certain tag~~
2. Support edit/modify/delete entries and tags
3. Support uploading multimedia?

# Future improvements

(Not so important if it's only for personal use)

1. All file manipulations are not synced, so race conditions can occur if there are multiple connections.
2. URLs parameters need more strict sanity checks.
3. The given API parser doesn't display the full query response. The definitions shown are cherry-picked.
