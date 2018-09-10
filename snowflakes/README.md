### Snowflakes

A snowflake of order n is formed by overlaying an equilateral triangle (rotated by 180 degrees) onto each equilateral triangle of the same size in a snowflake of order n-1. A snowflake of order 1 is a single equilateral triangle.

![snowflakes](https://raw.githubusercontent.com/plutov/practice-go/master/snowflakes/snowflakes.png)

Some areas of the snowflake are overlaid repeatedly. In the above picture, blue represents the areas that are one layer thick, red two layers thick, yellow three layers thick, and so on.

Given the N order find how many triangles in the snowlake are M levels deep.

### Example

```
OverlaidTriangles(1, 1) // 1
OverlaidTriangles(3, 1) // 30
OverlaidTriangles(3, 3) // 6
```

### Run tests with benchmarks

```
go test -bench .
```
