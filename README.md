# Concurrent Modification of Slice During Iteration in Go

This repository demonstrates a subtle bug that can occur when iterating over a Go slice while concurrently modifying the slice.  The example showcases the issue and its resolution.

**Bug:**
The `MyFunc` function iterates over a slice and prints each element.  If the slice is modified during the iteration, it can lead to out-of-bounds access or unexpected behavior.  This is because the index `i` is only calculated once before the loop starts, meaning concurrent modifications to the slice can invalidate the indices.