# Contributing to **the project**

This is an open source project and accepts contributions from the community. These contributions
are made in the form of Issues or
[Pull Requests](https://help.github.com/articles/about-pull-requests/) on official project repositories on GitHub.

Issues are a quick way to point out a bug, propose new features, improve current ones, etc.
If you find any kind of unwanted error in **the project** then, please mkae sure of:
please check a few things first:

1. The error was not already been discovered by third party. Check existing issues to do so
2. The issue has already been fixed (check the develop branch, or look for closed issues).
3. Can you solve the issue yourself?

Reporting issues is helpful but an even better approach is to send a pull request, which is done by
*forking* the main repository and committing to your own copy. This will require you to use the
version control system called Git.

The bug tracker of the **the project** repository is the tipical Github Issues. Please post your issues there, and only here.

## Guidelines

Before we look into how, here are the guidelines. If your pull Requests fail to pass these
guidelines it will be declined and you will need to re-submit when you've made the changes. This
might sound a bit tough, but it is required for us to maintain quality of the code-base.

### Branching ###

**the project** uses the [Git-Flow](http://nvie.com/posts/a-successful-git-branching-model/) branching model
which requires all pull requests to be sent to the *develop* branch. This is where the next planned
version will be developed. The *master* branch will always contain the latest stable version and is
kept clean so a *hotfix* (e.g: an emergency security patch) can be applied to master to create a
new version, without worrying about other features holding it up. For this reason all commits need
to be made to *develop* and any sent to *master* will be closed automatically. If you have multiple
changes to submit, please place all changes into their own branch on your fork, since this will
allow your master branch to be up to date with upstream.

One thing at a time: A pull request should only contain changes related to a single issue.

That does not mean only one commit, but one change - however many commits it took. The reason for this is that if you change X
and Y but send a pull request for both at the same time, we might really want X but disagree with
Y, meaning we cannot merge the request. Using the Git-Flow branching model you can create new
branches for both of these features and send two requests.

### Brach manamegment using **git-flow**

GitFlow is a branching model for Git, created by Vincent Driessen.
It has attracted a lot of attention because it is very well suited to collaboration and scaling the development team.

This is an alternative way of managing git repos, and as usually it is installed with:

```
sudo apt-get update
sudo apt-get install git-flow
```

Check out more information about git-flow at: [Introducing git-flow](https://datasift.github.io/gitflow/IntroducingGitFlow.html)

We propose you to use git-flow utility in your system so that, previous described
steps are converted to this:

```
git clone $repository_url
git checkout develop
git flow init
git flow feature start $NAME_OF_YOUR_NEW_FEATURE
```

This will create a new branch called **feature/NAME_OF_YOUR_NEW_FEATURE** based on the latest content of branch **develop**.
Now, your local git, is ready to work on a separate branch, and you can start building all changes you wish.

After making all neccessary changes, commit them:

```
git add $CHANGED_FILES
git commit -s -m 'write a descriptive comment'
git flow feature finish  #will merge the branch with local develop
git push
```

### Signing ###

You must sign your work, certifying that you either wrote the work or otherwise have the right to
pass it on to an open source project. Git makes this trivial as you merely have to use `--signoff`
on your commits to your **the project** fork.

`git commit -s`

This will sign your commits with the information setup in your git config, e.g.

`Signed-off-by: you username <username@domain.tld>`

By signing your work in this manner, you certify to a "Developer's Certificate of Origin". The
current version of this certificate is in the `DCO.txt` file in the root of this repository. More
information on how to sign your commits can be found
[here](https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work). The GPG key can be set with the
`user.signingkey` Git configuration option.

## How-to Guide ##

There are two ways to make changes, the easy way and the hard way. Either way you will need to
[join GitHub](https://github.com/join).

Easy way GitHub allows in-line editing of files for making simple typo changes and quick-fixes.
This is not the best way as you are unable to test if the code works. If you do this you could be
introducing syntax errors, etc, but for a Git-phobic user this is good for a quick-fix.

Hard way the best way to contribute is to *clone* your fork of **the project** to your development area. That
sounds like some jargon, but *forking* on GitHub means "making a copy of that repo to your account"
and *cloning* means "copying that code to your environment so you can work on it".

1. Set up Git ([Windows](https://help.github.com/articles/set-up-git/#platform-windows),
[Mac](https://help.github.com/articles/set-up-git/#platform-mac) &
[Linux](https://help.github.com/articles/set-up-git/#platform-linux)).
2. Go to **the project** repo you want to contribute to.
3. Fork it.
4. Clone your **project** repo.
5. Checkout the *develop* branch.
6. Create a **new** branch for your changes. At this point you are ready to start making changes.
6. Fix existing bugs on the issue tracker after taking a look to see nobody else is working on them.
7. Commit the files.
8. Push your branch to your fork.
9. [Send a pull request](https://help.github.com/articles/about-pull-requests/).

The core developers will now be alerted about the change and at least one of the team will respond.
If your change fails to meet the guidelines it will be bounced, or feedback will be provided to
help you improve it.

Once the core developer handling your pull request is happy with it they will merge it into
*develop* and your patch will be part of the next release.

### Keeping your fork up-to-date ###

Unlike systems like Subversion, Git can have multiple remotes. A remote is the name for a URL of a
Git repository. By default your fork will have a remote named *origin* which points to your fork,
but you can add another remote named *upstream* which points to the **the project** repository. This is a
read-only remote but you can pull from this develop branch to update your own.

If you are using command-line you can do the following:

1. `git remote add upstream $repository_url`
2. `git pull upstream develop`
3. `git push origin develop`

Now your fork is up to date. This should be done regularly, or before you send a pull request at
least.
