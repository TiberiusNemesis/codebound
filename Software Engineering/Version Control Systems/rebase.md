# Branch rebasing

Rebasing is a process that, in short, changes the *base* of a branch to a new commit.

This is another way (besides the merge) to integrate changes from one branch to another. However, rebasing creates a cleaner, linear history.

## How does it work?

Rebasing takes the commits from the current branch and reapplies them on top of another branch's commits. Unlike merging, this process does not create a merge commit.

Instead, it rewrites the commit history, creating a straight, linear progression of changes, which can appear as if they were made sequentially even if they were developed separately.

## When to use it?

Often, rebasing is used before merging a feature branch into the main branch. This streamlines the history, making the commit timeline of the main branch look cleaner.

It can also be used to locally clean up or organize a series of commits before sharing the branch with others.

Rebasing may require resolving conflicts just like in merges, but often must be handled at each commit being rebased. This can be quite time-consuming, especially if the number of commits being rebased is high.

## Sources

- [Git SCM documentation, git-rebase](https://git-scm.com/docs/git-rebase)
- [freeCodeCamp, The Ultimate Guide to Git Merge and Git Rebase](https://www.freecodecamp.org/news/the-ultimate-guide-to-git-merge-and-git-rebase/)
