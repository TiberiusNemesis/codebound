# Stashing

Essentially, [git](Software%20Engineering/Version%20Control%20Systems/git.md) stash is a way to preserve changes that we want to save, but are not yet ready for a [commit](Software%20Engineering/Version%20Control%20Systems/commit.md). This allows us to clear the working directory and switch branches without losing progress.

Stashing saves any modified files and reverts the stashed files to the last [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) state.

We can apply any existing stashes to the working directory, and the changes are reflected in the files without a [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) showing their addition (i.e a [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) still needs to be made to save it to the [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) in a more permanent manner).

## Sources

- [Git SCM documentation, git-stash](https://git-scm.com/docs/git-stash)
- [git stash - Saving Changes, Atlassian Git Tutorial](https://www.atlassian.com/git/tutorials/saving-changes/git-stash)
