

## TODO

- Compact adjacent intervals with same intensity, and test cases.
- More test cases.

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
coverage: 77.6% of statements
ok      test.com/intensity      0.314s
```
