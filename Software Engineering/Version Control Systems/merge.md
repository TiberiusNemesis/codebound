# Branch merging

Merging is one of the core operations in [git](Software%20Engineering/Version%20Control%20Systems/git.md). It combines changes and the history from two different branches into a single [branch](Software%20Engineering/Version%20Control%20Systems/branch.md), allowing for the integration of new features and updates from separate streams of effort.

## Types

### Fast-forward merge

Happens when the target [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) (i.e. the one where we are applying our changes to) has had no additional commits since our [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) (aka source [branch](Software%20Engineering/Version%20Control%20Systems/branch.md)) was created. Since all changes from the target [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) are already in the source [branch](Software%20Engineering/Version%20Control%20Systems/branch.md), this type of merge is simple and straightforward.

### Three-way merge

Used when both target and source branches both have had new commits. Git automatically creates a merge [commit](Software%20Engineering/Version%20Control%20Systems/commit.md) that combines the changes from both branches.

## Merge conflicts

A conflict can happen during merging when the same lines of code are altered differently in the branches being merged. When that occurs, the user needs to intervene to resolve these, usually editing the files manually to choose which changes are kept from the two alterations and then committing the resolved version.

## Sources

- [Git SCM documentation, Basic Branching and Merging](https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging)
- [Git merge conflicts, Atlassian Git Tutorial](https://www.atlassian.com/git/tutorials/using-branches/merge-conflicts)
