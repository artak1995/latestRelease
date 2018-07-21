### Introduction
This is a simple application that return the highest patch version of every release between a minimum version and the highest released version of github repositories. It reads the Github Releases list, uses SemVer for comparison and takes a path to a file as its first argument when executed. It reads this file, which is in the format of:
```
repository,min_version
kubernetes/kubernetes,1.8.0
coreos/flannel,0.6.1
```
and it will return 
```
latest versions of kubernetes/kubernetes: [1.10.1 1.9.6 1.8.11]
latest versions of coreos/flannel: [0.10.0 0.9.1 0.8.0 0.7.1 0.6.2]
```