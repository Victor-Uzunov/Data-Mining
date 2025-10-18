# Frog Leap Puzzle

## Problem Description

The Frog Leap Puzzle is a classic puzzle where:
- **Board length**: 2N + 1 cells
- **Initial state**: N frogs facing right (`>`), empty cell (`_`), N frogs facing left (`<`)
- **Goal**: Swap the frogs so left side has `<` frogs and right side has `>` frogs

### Rules
- Frogs move only in the direction they face
- A frog can step into an adjacent empty cell in front of it
- A frog can jump over exactly one frog into an empty cell (jump length = 2)

### Example

For N=2:
```
Input: 2
Output:
>>_<<
>_><<
><>_<
><><_
><_<>
_<><>
<_><>
<<>_>
<<_>>
```

## Testing

To test this solution:

```bash
# From repository root:
make run TASK=frog-leap-puzzle N=3
make test TASK=frog-leap-puzzle
make build TASK=frog-leap-puzzle

# Or run directly:
echo "3" | ./frog-leap-puzzle/go/solution
```

The solution supports timing mode with `FMI_TIME_ONLY=1` environment variable.
