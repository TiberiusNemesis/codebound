# Commit

A commit is a version of the current repository. Each version or *commit* has an unique id and a message associated to it. Usually, this message describes what changes were made in the commit when compared to the previous version.

The unique ID, often called "hash" or "SHA" (secure hash algorithm), is used to refer to a specific version with more precision than the commit message, and identifies which changes were made, when, and who made these changes. It looks something like this:

`911acde09b7e82e725fcfdb996ef9e503b94d88b`

Oftentimes a SHA is referred to using only its first few letters, especially when used in commands. In that case, you could see it being referred to as `911acde`.

Commits make up a sequential timeline that can track the changes being applied to the repository in order. Each commit is made to a branch, which represents a different sequence of versions of the same repository. A merge commit is when changes are taken from one branch and applied to another in order to combine all changes made in two version sequences (i.e. branches).

## Sources

- [GitHub Docs, About commits](https://docs.github.com/en/pull-requests/committing-changes-to-your-project/creating-and-editing-commits/about-commits)
- [Git SCM documentation, git-commit](https://git-scm.com/docs/git-commit)