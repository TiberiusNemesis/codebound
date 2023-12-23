# Branch rebasing

Rebasing is a process that, in short, changes the *base* of a [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) to a new [commit](Software%20Engineering/Version%20Control%20Systems/commit.md).

This is another way (besides the [merge](Software%20Engineering/Version%20Control%20Systems/merge.md)) to integrate changes from one [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) to another. However, rebasing creates a cleaner, linear history.

## How does it work?

Rebasing takes the commits from the current [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) and reapplies them on top of another [branch](Software%20Engineering/Version%20Control%20Systems/branch.md)'s commits. Unlike merging, this process does not create a [merge](Software%20Engineering/Version%20Control%20Systems/merge.md) [commit](Software%20Engineering/Version%20Control%20Systems/commit.md).

Instead, it rewrites the [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) history, creating a straight, linear progression of changes, which can appear as if they were made sequentially even if they were developed separately.

## When to use it?

Often, rebasing is used before merging a feature [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) into the main [branch](Software%20Engineering/Version%20Control%20Systems/branch.md). This streamlines the history, making the [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) timeline of the main [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) look cleaner.

It can also be used to locally clean up or organize a series of commits before sharing the [commit](Software%20Engineering/Version%20Control%20Systems/commit.md)0 with others.

Rebasing may require resolving conflicts just like in merges, but often must be handled at each [commit](Software%20Engineering/Version%20Control%20Systems/commit.md)1 being rebased. This can be quite time-consuming, especially if the number of commits being rebased is high.

## Sources

- [Git SCM documentation, git-rebase](https://git-scm.com/docs/git-rebase)
- [freeCodeCamp, The Ultimate Guide to Git Merge and Git Rebase](https://www.freecodecamp.org/news/the-ultimate-guide-to-git-merge-and-git-rebase/)
