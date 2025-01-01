# Troubleshooting Guide

Here we will add some well-known troubleshooting issues that you may run into, along with possible solutions.

- [Troubleshooting Guide](#troubleshooting-guide)
  - [VSCode Issues](#vscode-issues)
  - [Quickfeed Issues](#quickfeed-issues)
  - [Test Issues](#test-issues)
  - [GitHub Authentication Issues](#github-authentication-issues)
    - [GitHub: Permission denied (publickey) when clone/pull/push a repository](#github-permission-denied-publickey-when-clonepullpush-a-repository)
    - [GitHub: Could not read from remote repository](#github-could-not-read-from-remote-repository)
    - [WSL: Multiple ssh clients or conflicting git configurations](#wsl-multiple-ssh-clients-or-conflicting-git-configurations)
  - [Git: Unrelated histories when merging](#git-unrelated-histories-when-merging)

## VSCode Issues

If you have been using VSCode and the Go plugin in the past, you may wish to update the installed tools.
To do this, open the command palette by pressing `Ctrl+Shift+P` (Windows) or `Cmd+Shift+P` (macOS) and type `Go: Install/Update Tools`.
This will open a list of tools that you can install or update.
I suggest to install all of them.

Make sure that your VSCode settings include the formatting and static check options.
Your `settings.json` file should include at least the following:

```json
"go.useLanguageServer": true,
"gopls": {
    "formatting.gofumpt": true,
    "staticcheck": true
},
```

## Quickfeed Issues

1. If you have problems logging in into the application, cannot see your courses or lab submissions, please try these few steps:
   - Make sure you are using the right URL: `uis.itest.run`.
   - Make sure your browser does not have cookies blocked.
   - Make sure you are logged in with the right GitHub account in case you have several accounts.
   - Refresh the page.
   - Log out of the Quickfeed application, then log out of your GitHub account.
     Clear all browser data.
     Log back in to [QuickFeed](https://uis.itest.run).

## Test Issues

1. What is this `Test0Lint` thing that gives me test failures?

   Most lab assignments include a `Test0Lint` checker that checks that your Go code

   - follows Go coding style as defined by the `gofmt` program,
   - follows (some of the) best practices for Go coding,
   - does not contain any TODO or FIXME items.

   If you are getting a message like: `File is not goimports-ed (goimports)`, this means that you are not using the proper formatting of your Go code.
   To fix this, use the Go plugin for VSCode and ensure that it works to format your code.
   The formatter works when you save your file.
   It is easy to check that it works, by adding a line that is incorrectly formatted, e.g. `var myName =   "John Doe"`.
   (make sure to include some extra spaces between the tokens.)
   When you save your Go file, the spaces should be removed automatically.
   If this does not work, please see the VSCode Issues section above.

## GitHub Authentication Issues

GitHub allows [many ways to authenticate](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/about-authentication-to-github) with their services.
We recommend using the `gh` command-line tool, which is a GitHub CLI tool that can be used to authenticate with GitHub.

```sh
gh auth login
```

This command will open a browser window where you can authenticate with GitHub.

We also suggest that you create a [Passkey](https://docs.github.com/en/authentication/authenticating-with-a-passkey/about-passkeys) for your GitHub account.
This allows you to authenticate with GitHub using a fingerprint or other biometrics instead of a password.
This avoids the need to type your password every time you interact with GitHub.

Alternatively, you may also use an SSH key to authenticate with GitHub.
Specifically, GitHub allows you to work with repositories using two different protocols `https` or `ssh`, each one requires their own set of [configurations](https://docs.github.com/en/github/using-git/which-remote-url-should-i-use) and uses a different URL style to connect with the GitHub servers.

- HTTPS URL: `https://github.com/YOUR_USER/SOME_REPO.git`
- SSH URL: `git@github.com:YOUR_USER/SOME_REPO.git`

If you cannot authenticate with GitHub using the above methods, you may want to use the ssh protocol.
This allows you to connect to GitHub without supplying your username or password every time.
But for that to work, you need to configure it properly, as described [here](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh)

### GitHub: Permission denied (publickey) when clone/pull/push a repository

If you are getting this error is probably because you forgot to add your public key to GitHub, or you are trying to access the repository with a different key-pair.
In either case, take a look [here](https://docs.github.com/en/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account) and see how to add a key to GitHub and test if it is properly configured by running:

```console
ssh -T git@github.com
```

The command should display a message like:

```text
Hi YOUR_USERNAME! You've successfully authenticated, but GitHub does not provide shell access.
```

If you get an error, ensure that you are using the correct public key in your machine to connect to GitHub.
The content of your public key file, normally located at your home folder: `~/.ssh/id_rsa.pub` (or some other key type ending in `.pub`) should be the same as displayed in your GitHub account settings.
We have created a [SSH tutorial video](https://youtu.be/qik3HHZW6C0) illustrating the necessary steps (and a bit more).

### GitHub: Could not read from remote repository

There are many reasons that can result in the error below when cloning or pulling a GitHub repository:

```text
fatal: Could not read from remote repository

Please make sure you have the correct access rights
and the repository exists.
```

One common reason is a misconfigured remote URL.
As explained above, we use the `ssh` protocol to avoid having to type password for each interaction with GitHub.
Hence, if the output from the command `git remote -v` displays a URL using `https` as shown below, you will need to change these entries in order to use ssh.

```console
$ git remote -v
course-assignments  https://github.com/dat520-2025/assignments.git (fetch)
course-assignments  https://github.com/dat520-2025/assignments.git (push)
origin  https://github.com/dat520-2025/YOUR_USERNAME-labs.git (fetch)
origin  https://github.com/dat520-2025/YOUR_USERNAME-labs.git (push)
```

If this is the case, change the remote's URL to use ssh by running (remember to replace YOUR_USERNAME with your own):

```console
git remote set-url course-assignments git@github.com:dat520-2025/assignments.git
git remote set-url origin git@github.com:dat520-2025/YOUR_USERNAME-labs.git
```

The new remote's URL should be like this:

```console
$ git remote -v
course-assignments  git@github.com:dat520-2025/assignments.git (fetch)
course-assignments  git@github.com:dat520-2025/assignments.git (push)
origin  git@github.com:dat520-2025/YOUR_USERNAME-labs.git (fetch)
origin  git@github.com:dat520-2025/YOUR_USERNAME-labs.git (push)
```

### WSL: Multiple ssh clients or conflicting git configurations

If experience the following problem while using git with [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10):

```console
C:\Windows\System32\OpenSSH\ssh.exe" Permission denied
```

Ensure that your git configuration points to the correct ssh client path.

```console
$ git config --list --global
...
[core]
   sshCommand = "C:\Windows\System32\OpenSSH\ssh.exe"
```

If the output of the above command displays a different path from the command `which ssh` in your Linux subsystem.

```console
$ which ssh
/usr/bin/ssh
```

Then you may need to edit your configuration to use a ssh command that your user has permission to execute.
This can be done by editing your local or global git configuration.

To edit the global configuration (applies to all repositories on your Linux subsystem):

```console
$ git config --edit --global
...
[core]
   sshCommand = /usr/bin/ssh
```

To edit your local configuration (applies only to the current `assignments` repository):

```console
$ cd dat520-2025/assignments
$ git config --edit
...
[core]
   sshCommand = /usr/bin/ssh
```

## Git: Unrelated histories when merging

If you get an fatal error like the one bellow when doing a merge/pull:

```console
$ git pull course-assignments main
remote: Enumerating objects: 36, done.
remote: Counting objects: 100% (36/36), done.
remote: Compressing objects: 100% (34/34), done.
remote: Total 36 (delta 2), reused 36 (delta 2), pack-reused 0
Unpacking objects: 100% (36/36), 1.58 MiB | 3.96 MiB/s, done.
From https://github.com/dat520-2025/assignments
* branch            main       -> FETCH_HEAD
* [new branch]      main       -> course-assignments/main

fatal: refusing to merge unrelated histories
```

It is because you are probably trying to merge two unrelated projects into a single branch.
This situation may happen if, right after cloning your lab repository for the first time (initially empty),
you created some files and made some commits, and only later on you realized that you should have first synced with
the `course-assignments` to retrieve the lab assignments.
Then when you try to pull from the `course-assignments` remote, you get this kind of error.

As stated in the [git documentation](https://git-scm.com/docs/git-merge#Documentation/git-merge.txt---allow-unrelated-histories),
by default git refuses to merge histories that do not share a common ancestor.
You can use the option `--allow-unrelated-histories` in the git pull command to override this setting, like below:

```console
git pull course-assignments main --allow-unrelated-histories
```
