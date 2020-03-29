# kustomize-check-action

`kustomize-check-action` do check kustomizations when changed. 

```
  check:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
      with:
        fetch-depth: 0 # kustomize-check-action uses git diff cmd so fetch-depth should be 0 or enough depths.
    - name: kz-check
      uses: 'anarcher/kustomize-check-action@master'
      with:
        # paths is optional. when changes are in this paths, check-action do check it.
        paths: |
            base/
            overlay/
```


      
