# Neural Networks (MLP) — Activations, Backpropagation, and Logic Gates

## Problem
Train a small feedforward neural network (MLP) to learn Boolean functions (AND, OR, XOR) over two inputs. Explore the effect of activation choice (sigmoid vs tanh) and hidden layer configuration.

## Model: Multilayer Perceptron (MLP)
- Architecture: input → one or more hidden layers → output
- Each layer performs: \( z^{(l)} = a^{(l-1)} W^{(l)} + b^{(l)} \), \( a^{(l)} = \phi(z^{(l)}) \)
- Output for binary target here is a single neuron with activation \(\phi\)

### Activations
- Sigmoid: \( \sigma(x) = \frac{1}{1+e^{-x}} \)
  - Range (0,1); derivative: \( \sigma'(x) = \sigma(x) (1 - \sigma(x)) \)
  - Pros: probabilistic interpretation; Cons: saturation → vanishing gradients
- Tanh: \( \tanh(x) \)
  - Range (−1,1); derivative: \( 1 - \tanh^2(x) \)
  - Often preferred over sigmoid for hidden layers due to centered outputs

### Learnability: XOR requires hidden layer(s)
- Single-layer perceptron cannot represent XOR (not linearly separable)
- MLP with at least one hidden layer can learn XOR via nonlinear transformations

## Training: Backpropagation + Gradient Descent
- Forward pass: compute activations layer by layer
- Loss (implicitly): squared error \( L = \frac{1}{2} \| \hat{y} - y \|^2 \)
- Backpropagation:
  1. Output delta: \( \delta^{(L)} = (\hat{y} - y) \odot \phi'(a^{(L)}) \)
  2. Hidden deltas: \( \delta^{(l)} = (\delta^{(l+1)} W^{(l+1)\top}) \odot \phi'(a^{(l)}) \)
  3. Gradients: \( \partial W^{(l)} = a^{(l-1)\top} \delta^{(l)} \), \( \partial b^{(l)} = \sum \delta^{(l)} \)
- Update rule: \( W \leftarrow W - \eta \partial W \), \( b \leftarrow b - \eta \partial b \) (\(\eta\) = learning rate)

## Implementation Notes (This Task)
- File: `neural-networks/python/neural-networks.py`
- Weights/biases: initialized from normal distribution with fixed seed
- Activations selectable via `activation_id`:
  - `0` → sigmoid
  - `1` → tanh
- Hidden layers specified as a list (e.g., `[3]` = one hidden layer with 3 neurons)
- Training epochs: 20,000; learning rate: 0.1
- Dataset:
  - Inputs X: all 2-bit combinations `[[0,0],[0,1],[1,0],[1,1]]`
  - Targets y depend on function: AND, OR, XOR

## Input/Output Format (Interactive)
The script reads from stdin in sequence:
1) Target function: `AND`, `OR`, `XOR`, or `ALL`
2) Activation id: `0` (sigmoid) or `1` (tanh)
3) Number of hidden layers: integer `h`
4) `h` lines, each with the size of the hidden layer (e.g., `3`, `2` ...)

Examples:
- Learn XOR with one hidden layer of size 3 using tanh:
```
XOR
1
1
3
```
- Learn all three functions with sigmoid and two hidden layers [3,2]:
```
ALL
0
2
3
2
```

Output:
- Prints the function name and predictions for each input pair, e.g.:
```
XOR:
(0,0) -> 0.01
(0,1) -> 0.98
(1,0) -> 0.98
(1,1) -> 0.02
```
Interpret values close to 0 as class 0 and close to 1 as class 1.

## Complexity
- Forward pass per sample: \( \sum_l O(n_{l-1} n_l) \)
- Backward pass similar order; total \( O(E \cdot N \cdot \sum_l n_{l-1} n_l) \) for E epochs and N samples
- Small dataset here → dominated by epochs and hidden sizes

## Practical Tips
- XOR: use at least one hidden layer; tanh often converges faster than sigmoid
- Initialization: fixed seed improves reproducibility; different seeds may improve results
- Learning rate: if predictions don’t progress, try smaller \(\eta\) or fewer epochs for stability
- Output thresholding: apply a threshold (e.g., 0.5 for sigmoid; 0 for tanh) when you need hard class labels

## How to Run
- Temporary venv:
```bash
make run neural-networks
```
Then paste interactive inputs as shown above.

- Persistent venv:
```bash
make venv neural-networks
make run neural-networks
```

## Exam Tips
- Derive backprop deltas and gradients; explain chain rule through layers
- Discuss why XOR is not linearly separable and how hidden layers solve it
- Compare sigmoid vs tanh: ranges, derivatives, effect on gradient flow
- Explain training dynamics (learning rate, epochs) and potential vanishing gradients
