import numpy as np
import sys

def sigmoid(x):
    return 1 / (1 + np.exp(-x))

def sigmoid_derivative(x):
    return x * (1 - x)

def tanh_activation(x):
    return np.tanh(x)

def tanh_derivative(x):
    return 1 - x**2

class NeuralNetwork:
    def __init__(self, layer_sizes, activation_id, learning_rate=0.1):
        self.layer_sizes = layer_sizes
        self.learning_rate = learning_rate
        self.activation_id = activation_id
        self.weights = []
        self.biases = []

        np.random.seed(42)
        for i in range(len(layer_sizes) - 1):
            n_in = layer_sizes[i]
            n_out = layer_sizes[i+1]
            self.weights.append(np.random.randn(n_in, n_out))
            self.biases.append(np.random.randn(1, n_out))

    def _activation(self, x):
        if self.activation_id == 0:
            return sigmoid(x)
        else:
            return tanh_activation(x)

    def _derivative(self, x):
        if self.activation_id == 0:
            return sigmoid_derivative(x)
        else:
            return tanh_derivative(x)

    def forward(self, X):
        activations = [X]
        current_input = X

        for i in range(len(self.weights)):
            net_input = np.dot(current_input, self.weights[i]) + self.biases[i]
            output = self._activation(net_input)
            activations.append(output)
            current_input = output

        return activations

    def train(self, X, y, epochs=20000):
        for _ in range(epochs):
            activations = self.forward(X)
            final_output = activations[-1]

            error = final_output - y
            deltas = [error * self._derivative(final_output)]

            for i in range(len(self.weights) - 1, 0, -1):
                delta_next = deltas[-1]
                current_activation = activations[i]
                delta = np.dot(delta_next, self.weights[i].T) * self._derivative(current_activation)
                deltas.append(delta)

            deltas.reverse()

            for i in range(len(self.weights)):
                input_to_layer = activations[i]
                delta = deltas[i]
                self.weights[i] -= self.learning_rate * np.dot(input_to_layer.T, delta)
                self.biases[i] -= self.learning_rate * np.sum(delta, axis=0, keepdims=True)

    def predict(self, X):
        activations = self.forward(X)
        return activations[-1]

def get_data(func_name):
    X = np.array([[0,0], [0,1], [1,0], [1,1]])

    if func_name == "AND":
        y = np.array([[0], [0], [0], [1]])
    elif func_name == "OR":
        y = np.array([[0], [1], [1], [1]])
    elif func_name == "XOR":
        y = np.array([[0], [1], [1], [0]])
    else:
        return None, None
    return X, y

def run_experiment(target_func, activation, hidden_layers_config):
    input_size = 2
    output_size = 1
    layer_sizes = [input_size] + hidden_layers_config + [output_size]

    nn = NeuralNetwork(layer_sizes, activation)
    X, y = get_data(target_func)

    nn.train(X, y)

    print(f"\n{target_func}:")
    predictions = nn.predict(X)
    for i in range(len(X)):
        input_pair = tuple(X[i])
        pred_val = predictions[i][0]
        print(f"({input_pair[0]},{input_pair[1]}) -> {pred_val:.4f}")

if __name__ == "__main__":
    func_input = sys.stdin.readline().strip()
    if not func_input: func_input = input().strip()

    act_input = sys.stdin.readline().strip()
    if not act_input: act_input = input().strip()
    activation_choice = int(act_input)

    hl_count_input = sys.stdin.readline().strip()
    if not hl_count_input: hl_count_input = input().strip()
    num_hidden_layers = int(hl_count_input)

    hidden_config = []
    for _ in range(num_hidden_layers):
        n_input = sys.stdin.readline().strip()
        if not n_input: n_input = input().strip()
        hidden_config.append(int(n_input))

    if func_input == "ALL":
        functions = ["AND", "OR", "XOR"]
    else:
        functions = [func_input]

    for f in functions:
        run_experiment(f, activation_choice, hidden_config)