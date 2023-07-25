# Matrix-gen

This is a small tool used in CI to generate a kernel compatibility matrix for each task given the ansible output folder.  

Example:

```
./matrix_gen --root-folder ~/ansible_output --output-file matrix_x86_64.md
```

Available options:
```
./matrix_gen -h
Usage of ./matrix_gen:
  -output-file string
        output file where the generated matrix is stored (default "matrix.md")
  -root-folder string
        ansible output root folder (default "~/ansible_output")
```
