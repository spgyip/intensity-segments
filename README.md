
## TODO

- [X] Compact adjacent intervals with same intensity.
- More test cases for invalid range input.
- Implement `Get` method querying intensity of a specified input.

## Build

```shell
go build .
```

## Test coverage

```shell
#!] go test -v -cover                    
=== RUN   TestIntensitySegmentsEmpty
--- PASS: TestIntensitySegmentsEmpty (0.00s)
=== RUN   TestIntensitySegmentsSplit
--- PASS: TestIntensitySegmentsSplit (0.00s)
=== RUN   TestIntensitySegmentsAdd
--- PASS: TestIntensitySegmentsAdd (0.00s)
=== RUN   TestIntensitySegmentsSet
--- PASS: TestIntensitySegmentsSet (0.00s)
PASS
        test.com/intensity      coverage: 76.9% of statements
ok      test.com/intensity      0.564s  coverage: 76.9% of statements
```
