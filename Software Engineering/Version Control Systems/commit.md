# Commit

A commit is a version of the current repository. Each version or *commit* has an unique id and a message associated to it. Usually, this message describes what changes were made in the commit when compared to the previous version.

The unique ID, often called "hash" or "SHA" (secure hash algorithm), is used to refer to a specific version with more precision than the commit message, and identifies which changes were made, when, and who made these changes. It looks something like this:

`911acde09b7e82e725fcfdb996ef9e503b94d88b`

Oftentimes a SHA is referred to using only its first few letters, especially when used in commands. In that case, you could see it being referred to as `911acde`.

## How are they used?

Commits are the backbone of Git's version control system. They make up a sequential timeline that can track the changes being applied to a repository, allowing users to track the history of changes made in that repo.

Each commit is applied to a *branch,* which represents a sequence of versions of the same repository. This, especially when coupled with descriptive commit messages, allows for easy understanding of the progress made over time.

A standard is to be concise, yet descriptive. A good commit message will explain clearly what the commit does (and why, if necessary). It is recommended to use the imperative mood, as if we're giving out a command or instruction (e.g. "Add feature X" instead of "Added feature X"). This is because in some git command-line operations (such as `git cherry-pick` or `git log`) we can look at the message as the *result* of applying a certain commit.

Another way to think of it is: if asked the question *"What would your commit do if it was applied to the current branch?"*, what would you answer? That answer is likely a good, imperative commit message.

Commit messages may also use the past tense. Though this is not officially recommended, it is a widely used standard in commit messages because oftentimes it looks more natural in a commit history being read top-to-bottom.

*Author's note: I honestly prefer past tense. It feels more natural to me. However, I do understand why present tense might be a choice for many Git users. I think a good way of handling it is to have a standard in each project -- if we're doing past tense, we're doing past tense. If it's imperative, then everyone's gotta use imperative.*

### Merge commits

A merge commit is when changes are taken from one branch and applied to another in order to combine all changes that were made in both branches. This also combines their commit histories, creating a single, unified sequence of commits.

## Sources

- [GitHub Docs, About commits](https://docs.github.com/en/pull-requests/committing-changes-to-your-project/creating-and-editing-commits/about-commits)
- [Git SCM documentation, git-commit](https://git-scm.com/docs/git-commit)