### List Zipper

In functional programming, there's a concept called the List Zipper, which is a really elegant way of navigating and modifying lists. Imagine you have a list, but instead of just thinking of it as a flat sequence, you pick a point in that list to "focus" on. You divide the list into two parts: everything before that point (which we typically store in reverse order for efficiency), and everything from that point onward. The focus allows you to move left or right, almost like a cursor in a text editor, and perform edits in a very localized way.

## Task

Your task is to implement a generic List Zipper in Go that supports creation from a standard list or empty state, navigation by moving the focus left or right, inspection of the current focused value, modification through insertion or deletion near the focus, and conversion back to a regular list. Importantly, all operations must respect immutability—meaning any modifying function must return a new instance of the zipper without altering the original. The structure should have two fields: left (a reversed slice of elements before the focus) and right (a slice of elements at and after the focus). For example, moving the focus left would shift the last element of left to the front of right, and moving right would do the opposite. If right is empty, the focus is at the end; if left is empty and right is not, it’s at the beginning.

Note that it must contain fields to represent the elements before the focus (`left`, reversed) and the elements at/after the focus (`right`).


For example, say you start with the list [1, 2, 3, 4]. If you move right once, your focus shifts from 1 to 2. Then maybe you insert 99, so the list becomes [1, 99, 2, 3, 4]. The original list and zipper stay untouched—you get a new zipper each time. That’s the whole idea: local, safe changes without side effects.

### Run tests with benchmarks

```
go test -bench .
```
