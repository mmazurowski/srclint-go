# srclint - source code consistency helper
Goal of this project is to keep consistent and concise codebase spread across many repositories. 

## How to use
```shell
srclint .
```

## How it works:
Each directory needs to have config file in working directory for the program. Basic config looks like this:
```yaml
version: 1
required:
  - "srclint.yml"
structure:
  - "src/**" 
ignore:
  - ".idea"
```


### Error Codes
1. Missing required file
2. Source code structure failed validation