# Git

Git (sometimes called Git SCM) is a tool used to track and manage changes made in software code. That's the main usage of git, but really, being a highly capable version control system, it can be utilized for tracking more than just code.

As of December 2023, it is the most widely utilized version control system in software engineering.

## How does it work?

Git works by tracking folders and subfolders, along with the files contained in them. It stores changes to these files in a type of database -- then allows for these changes to be reverted, tracked, modified, combined, shared, and managed in a multitude of ways. This is particularly effective in a team environment, where it is possible to have a large number of people working on the same set of files.

### Repositories, Commits and Branches

A folder structure being tracked by git is called a *repository*, and inside a repository, we can save a snapshot of the folder structure by using a *[commit](Software%20Engineering/Version%20Control%20Systems/commit.md)*.

Commits are organized in structures called *branches.* Each [branch](Software%20Engineering/Version%20Control%20Systems/branch.md) represents an independent sequence of commits, and to combine different sets of commits, a [merge](Software%20Engineering/Version%20Control%20Systems/merge.md) or a [rebase](Software%20Engineering/Version%20Control%20Systems/rebase.md) between branches can be performed. This allows for simultaneous development by multiple members of an organization even if they are altering the same repository.

## Sources

- [Git](https://git-scm.com/)
- [What is Version Control?](https://www.atlassian.com/git/tutorials/what-is-version-control)
