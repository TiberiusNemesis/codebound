# Stashing

Essentially, `git stash` is a way to preserve changes that we want to save, but are not yet ready for a commit. This allows us to clear the working directory and switch branches without losing progress.

Stashing saves any modified files and reverts the stashed files to the last commit state.

We can apply any existing stashes to the working directory, and the changes are reflected in the files without a commit showing their addition (i.e a commit still needs to be made to save it to the branch in a more permanent manner).

## Sources

- [Git SCM documentation, git-stash](https://git-scm.com/docs/git-stash)
- [git stash - Saving Changes, Atlassian Git Tutorial](https://www.atlassian.com/git/tutorials/saving-changes/git-stash)
