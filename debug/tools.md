Certainly! Here's a translation and explanation of the provided instructions on how to use Git Bisect for debugging issues in your code:

---

### How to Use Git Bisect for Debugging

Git Bisect is an invaluable tool for narrowing down when an issue was introduced into your codebase. It helps you quickly identify the commit that introduced a bug by iteratively dividing the range of commits and testing each half until the problematic commit is found.

#### Step-by-Step Guide:

1. **Start Bisecting:**
   - Ensure you have a clean working directory.
   - Use `git bisect start` to initiate the process.
   
2. **Define Initial States:**
   - You can define custom terms for good and bad states, but usually, `good` and `bad` are used:
     ```sh
     git bisect start
     ```

3. **Specify the Range of Commits:**
   - Define your initial range by specifying a known `good` commit and a suspected `bad` commit.
     ```sh
     git bisect good <commit-hash-good>
     git bisect bad  <commit-hash-bad>
     ```
   
4. **Run Tests for Current Commit:**
   - Git Bisect will automatically switch to the current suspect commit, which is in the middle of your specified range.
   - It will tell you to run tests on this commit and report whether it passes or fails.

5. **Iteratively Narrow Down the Range:**
   - Depending on the result:
     - If the test fails (bad), run `git bisect bad` to keep narrowing down the suspect commits in the upper half.
       ```sh
       git bisect bad
       ```
     - If the test passes (good), run `git bisect good` to keep narrowing down the suspect commits in the lower half.
       ```sh
       git bisect good
       ```

6. **Identify the Problematic Commit:**
   - Continue this process until Git Bisect identifies a single commit as the problematic one.

7. **View Progress (Optional):**
   - To see the current range of remaining commits to check:
     ```sh
     git bisect visualize --oneline
     ```

8. **Clean Up:**
   - Once you have identified the problematic commit, use `git bisect reset` to restore your repository to the state before starting Bisect.
     ```sh
     git bisect reset
     ```
   - Consider reinstalling the correct version of any dependent packages that might need updating.

9. **Report the Issue:**
   - Go to the GitHub or GitLab commit URL and open an issue there, linking to the PR from which the problematic commit originated for further discussion.
   
10. **Handle Special Cases:**
    - **Skipping a Commit:** If a specific commit cannot be tested (e.g., due to changes in dependencies), use `git bisect skip`:
      ```sh
      git bisect skip <commit-hash>
      ```
    - **Reversing the Order of Good and Bad States:** 
      - If you need to reverse the order, define new terms like `fixed` and `broken`:
        ```sh
        git bisect start --term-new=fixed --term-old=broken
        git bisect broken 6c94774
        ```
        Then use `git fixed / git broken` instead of `git good / git bad`.

11. **Complications:**
    - If dependencies change between commits, you might need to reinstall them:
      ```sh
      pip install numpy==1.25  # Example for Python packages
      ```

---

By following these steps, you can effectively use Git Bisect to identify when an issue was introduced into your codebase and take the necessary actions to fix it.