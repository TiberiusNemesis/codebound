# Codebound

![codebound-header](/Assets/codebound.png)

Welcome to *Codebound*, a repository where I document my journey through the realms of software engineering and computer science.

## Overview

This repository is a collection of markdown files containing my annotations and study notes. Containing algorithms, data structures, best practices, software engineering principles, you can find it all. From fundamental concepts all the way to the tiniest details -- this is my digital garden, where knowledge grows and ideas intertwine.

## Table of Contents

- Computer Science
  - Algorithms
  - Asymptotic Notation
  - Data Structures

- Software Engineering
  - The Internet

## How to use this repository

Navigate through the folders to find study notes on various topics. Each markdown file is an annotation or a collection of notes on a particular concept. All of them are original content -- they are my own words describing my understanding of the topic.

There is a script made in Golang for this repo that automatically links topics -- each time a match is found for one of the file names inside a Markdown file, a link will automatically be added pointing to the other file.

e.g.
Before the script has been executed, a line may look like this:

```markdown
We can clearly see that the lorem ipsum has dolor and did not sit amet.
```

If a file named `lorem ipsum.md` exists elsewhere in the project, the script will turn the line above into:

```markdown
We can clearly see that the [lorem ipsum](file/path/works%20with%20spaces/to/lorem%20ipsum.md) dolor sit amet.
```

To execute the script, use:

```bash
go run markdown_linker.go
```

## Contributing

While this repository serves as a personal knowledge base, contributions are welcome.  
If you see an error or want to add something, feel free to open a pull request.

## License

This project is open-sourced under the [MIT License](LICENSE).

## Acknowledgments

Inspiration drawn from the works of [Brandon Sanderson](https://en.wikipedia.org/wiki/Brandon_Sanderson) and the incredible community of software engineers all around the globe.

Credits to DALL-E for the header image.

---

Authored by Tiberius "Nemesis" Dourado.

*Journey before destination.*
