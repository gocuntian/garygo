# Contributing to Go-Server

Thank you for contributing to Go-Server. Your helping to make it even better than it is today! As a contributor, here are a few guidelines we ask for you to kindly follow.

 - [Issues and Bugs](#issue)
 - [Feature Requests](#feature)
 - [Submission Guidelines](#submit)
 - [Coding Rules](#rules)
 - [Commit Message Guidelines](#commit)

## <a name="issue"></a> Found an Issue?
If you find a bug in the source code or a mistake in the documentation, you can help by [submitting an issue](#submit-issue) to our [GitHub Repository][github]. Or even better, you can [submit a Pull Request](#submit-pr) with a fix.

## <a name="feature"></a> Want a Feature?
You can *request* a new feature by [submitting an issue](#submit-issue) to our [GitHub Repository][github]

For a **Major Feature**, first open an issue and outline your proposal so that it can be discussed. This will also allow us to better coordinate our efforts, prevent duplication of work, and help you to craft the change so that it is successfully accepted into the project.

**Small Features** can be crafted and directly [submitted as a Pull Request](#sumit-pr).

## <a name="submit"></a> Submission Guidelines

### <a name="submit-issue"></a> Submitting an Issue
Before you submit an issue, please take a second to make sure there is not an open bug for it already in Issues.

If your issue appears to be a bug, and hasn't been reported, open a new issue. Help us to maximize the effort we can spend fixing issues and adding new features, by not reporting duplicate issues. Providing the following information will increase the chances of your issue being dealt with quickly:

* **Motivation for or Use Case** - explain why this is a bug for you
* **Browsers and Operating System** - is this a problem with all browsers and/or Operating Systems?
* **Reproduce the Error** - provide an example (feel free to use [Go Playground][goplayground] or something similar) or a unambiguous set of steps.
* **Related Issues** - has a similar issue been reported before?
* **Suggest a Fix** - if you can't fix the bug yourself, perhaps you can point to what might be causing the problem (line of code or commit)

### <a name="submit-pr"></a> Submitting a Pull Request (PR)
Before you submit your Pull Request (PR) consider the following guidelines:

* Search [GitHub](https://github.com/byronmansfield/go-server/pulls) for an open or closed PR that relates to your submission. You don't want to duplicate effort.
* Make your changes in a new git branch:

		 ```shell
		 git checkout -b my-fix-branch master
		 ```

* Create your patch, **including appropriate test cases**.
* Follow our [Coding Rules](#rules).
* Ensure that all tests pass.
* Commit your changes using a descriptive commit message that follows our [commit message conventions](#commit). Adherence to these conventions is necessary because release notes are automatically generated from these messages.

		 ```shell
		 git commit -a
		 ```
  Note: the optional commit `-a` command line option will automatically "add"
	and "rm" edited files.

* Push your branch to GitHub:

		```shell
		git push origin my-fix-branch
		```

* In GitHub, send a pull request to `go-server:develop`.
* If we suggest changes then:
  * Make the required updates.
  * Re-run tests to ensure tests are still passing.
	* Rebase your branch and force push to your GitHub repository (this will
	update your Pull Request):

		```shell
		git rebase develop -i
		git push -f
		```

That's it!

#### After your pull request is merged

After your pull request is merged, you can safely delete your branch and pull the changes from the main (upstream) repository:

* Delete the remote branch on GitHub either through the GitHub web UI or your local shell as follows:

		```shell
		git push origin --delete my-fix-branch
		```

* Check out the develop branch:

		```shell
		git checkout develop -f
		```

* Delete the local branch:

		```shell
		git branch -D my-fix-branch
		```

* Update your develop with the latest upstream version:

		```shell
		git pull --ff upstream develop
		```

## <a name="rules"></a> Coding Rules
To ensure consistency throughout the source code, keep these rules in mind as you are working:

* All features or bug fixes **must be tested** by one or more specs (unit-tests).
* Please remove any development tools, notes, etc, such as console logs or commented out code and/or notes (Unless its go doc's)
* Go Doc your code - this helps the developers unfamiliar with the code base know whats going on

## <a name="commit"></a> Commit Message Guidelines

We have very precise rules over how our git commit messages can be formatted. This leads to **more readable messages** that are easy to follow when looking through the **project history**. But also, we use the git commit message to **generate the Go-Server Changelog**

### Commit Message Format
* Message needs to be descriptive
* Must have an Issue number with it.
* Any line of the commit message cannot be longer than 100 characters! This allows the message to be easier to read on GitHub as well as in various git tools.
* No period (.) at the end

[github]: https://github.com/byronmansfield/go-server
[goplayground]: https://play.golang.org/
