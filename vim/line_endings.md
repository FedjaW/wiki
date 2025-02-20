# Confiure Line Endings

Stackoverflow:
https://stackoverflow.com/questions/9976986/force-lf-eol-in-git-repo-and-working-copy/34810209#34810209


Official Docs:
https://docs.github.com/en/get-started/getting-started-with-git/configuring-git-to-handle-line-endings#refreshing-a-repository-after-changing-line-endings


```shell
$ echo "* text=auto" >> .gitattributes
$ rm .git/index     # Remove the index to force Git to
$ git reset         # re-scan the working directory
$ git status        # Show files that will be normalized
$ git add -u
$ git add .gitattributes
$ git commit -m "Introduce end-of-line normalization"
```