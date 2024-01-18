### Pre-requirement

- Have git version ≥ 2.34” in requirements, (update with : `brew install git`)
- Have an SSH key that has the usage type : **Signing**

<aside>
💡 Don’t delete your current SSH key if you have the right options for commit signing to avoid possible issues on your config

</aside>

## Add commit signing to your config (5 min)

**Sign commits with an SSH key on your computer**

1. Configure Git to use SSH for commit signing:

   ```sh
   git config --global gpg.format ssh
   ```

2. Specify which SSH key should be used as the signing key. Specify the private key that is linked to your Gitlab account. Indicate the path of your private key in the command below.

   ```sh
   git config --global user.signingkey ~/.ssh/privateKey
   ```

3. Activate commit signing for every time you commit :

   ```bash
   git config --global commit.gpgsign true

   ```

### Check if your commits are signed!

To check that your commits are signed, please go to our project and in a branch make a commit. If you can see the verified tag on your commit then your commit has been signed !

- You can go to the link of your commit and find the `verified` tag on the top-left corner.

## Possible issues

**Commit signing is active but I'm getting an error.**

- GPG failed to sign the data
  **Error :**

  ```jsx
  error: gpg failed to sign the data
  fatal: failed to write commit object
  ```

  **Why ?**

  - **You have no key that is linked to signing gitlab commits**

    **Solution :**

    - Specify which SSH key should be used as the signing key. Specify the private key that is linked to your Github account.

      ```jsx
      git config --global user.signingkey ~/.ssh/nameOfKey

      ```

  - **The key you are singing with is not the one linked to your gitlab account**
    **Solution :**

    - Specify which SSH key should be used as the signing key. Specify the private key that is linked to your Github account.

      ```jsx
      git config --global user.signingkey ~/.ssh/nameOfKey
      ```

  - Wrong format when commit signing
    Configure Git to use SSH for commit signing:

    ```jsx
    git config --global gpg.format ssh
    ```

  - Outdated version of Git
    You need to have git version ≥ 2.34” in requirements, (update with : `brew install git`)

- **Wrong user/email**
  - .gitconfig
    You may have used another email in your `.gitconfig` , please verify that you are using the same commit email as displayed in your Github
    1. Check your user : `git config --global user.email`
    2. Set your user email :git config --global user.email `john.doe@blabla.com`
  - Github profile
    Check your for Public email and Commit email sections.
- **Failed to push some refs to github.com**

  - You might have unsigned commits that you are trying to push, you can re-sign all commits
    You can use this script to help you sign your previous commits, but make sure to have :

    - Signed commits are configured
    - Your local master branch is up-to-date
    - Be on the your feature branch
    - Commit your changes on the feature branch

    ```bash
    func list_commits_in_current_branch() {
    	git log --pretty=format:"%h" --no-patch --no-merges `git rev-parse --abbrev-ref HEAD` ^master
    }

    IFS=$'\\n' commits=($(list_commits_in_current_branch))
    git reset --hard master
    for c in $commits
    do
    	git cherry-pick $c --no-commit
    	git commit -C $c
    done

    ```

  - Your branch might be faulty
    - Make a new branch and add your changes to this new branch and push new commits to the new feature branch.
