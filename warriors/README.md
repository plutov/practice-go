### Warriors

BitVille and ByteVille are 2 warring tribes. BitVille has a spy telescope called the Hobble Scope, to count the number of Byte Warriors in ByteVille. The Hobble Scope, however, has a problem. Its primary lens is contaminated with bugs which block part of each image. The Hobble Scope's flawed images are stored by pixels in a file.

Each image is square and each pixel or cell contains either a 0 or a 1. At each pixel location, 1 is recorded if part or all of a Byte Warrior is present and a 0 if any other object, including a bug, is visible. 1. A Byte warrior is represented by at least a single binary 1. 2. Cells with adjacent sides on the same row and/or column or Cells which are diagonally adjacent - which contain binary ones, comprise one Byte Warrior. A very large image of one Byte Warrior might contain all ones. 3. Distinct Byte Warriors do not touch one another. 4. There is no wrap-around. Pixels on the bottom are not adjacent to the top and the left is not adjacent to the right (unless, of course, there are only 2 rows or 2 columns).

Output:

The output should be a single number specifying the number of unique Byte Warriors in the image.

### Example

Input:

```
100100
001010
000000
110000
111000
010100
```

Output: 3

Why?

There is 1 warrior on top left most corner (just 1 pixel), 2nd warrior on the top right (3 "1" pixels), 3rd warrior on the left bottom corner (7 "1" pixels).

### Run tests with benchmarks

```
go test -bench .
```
