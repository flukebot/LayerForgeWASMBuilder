# LayerForgeWASMBuilder
Self-documenting, browser-based AI framework with dynamic introspection and real-time execution

**LayerForgeWASMBuilder** is an advanced tool designed to dynamically introspect and compile the [LayerForge AI framework](https://github.com/flukebot/LayerForge) into WebAssembly (WASM). This setup allows `LayerForge`'s capabilities to be run, tested, and showcased directly in a web browser.

With WASM integration, you can execute LayerForge AI computations client-side, enabling interactive testing and benchmarking in the browser. Any new methods or updates made to `LayerForge` are automatically included in the WASM build.

## Features

- **Dynamic Method Introspection**: Detects and wraps all available methods in `LayerForge` for browser-based execution.
- **Browser-Based Execution**: Runs LayerForge benchmarks and other AI computations directly in a web browser via WebAssembly.
- **Self-Updating Environment**: Each build automatically includes new or modified methods from the LayerForge framework.
- **Test, Benchmark, and Showcase**: Easily demonstrate LayerForge's capabilities with real-time, browser-accessible results.

## Requirements

- [Go](https://golang.org/dl/) installed (version 1.16 or higher is recommended).
- [LayerForge AI framework](https://github.com/flukebot/LayerForge) cloned as a parent directory of LayerForgeWASMBuilder.
- A local HTTP server to serve the WASM file (like Python's simple HTTP server or similar).

## Getting Started

1. **Clone the Repositories**:

   Clone the LayerForge AI framework into the parent folder:
   
   `git clone https://github.com/flukebot/LayerForge.git`

   Then, clone the LayerForgeWASMBuilder repository inside the same parent directory:

   `git clone https://github.com/flukebot/LayerForgeWASMBuilder.git`

2. **Build the WASM File**:

   Navigate to the `LayerForgeWASMBuilder` directory, then run the following command to compile `LayerForge` into WebAssembly:

   `GOOS=js GOARCH=wasm go build -o blueprint.wasm main.go`

   This command generates the `blueprint.wasm` file, which can then be run in the browser.

3. **Prepare the WASM Execution Environment**:

   Ensure `wasm_exec.js` (found in your Go installation, typically under `GOROOT/misc/wasm/wasm_exec.js`) is in the same directory as your `index.html` file.

4. **Start a Local HTTP Server**:

   To serve the files locally, start an HTTP server from within the `LayerForgeWASMBuilder` directory.

   Using Python, for example:

   `python3 -m http.server 8000`

5. **Access the Interface**:

   Open your browser and navigate to:

   `http://localhost:8000`

   You should see the interface where you can:
   - Run introspection on the LayerForge framework to view available methods.
   - Execute benchmarks and other LayerForge methods directly in your browser.

## Usage

- **Get Blueprint Methods**: Click the "Get Blueprint Methods" button to view all available methods within the `LayerForge` framework.
- **Run Benchmark**: Use the "Run Benchmark" button to execute LayerForge's benchmark function in the browser and view results in real-time.

Each build will automatically include any new functions you add to `LayerForge`, making this a powerful tool for development, testing, and demonstrating `LayerForge` capabilities on the client side.

## Example Output

When you click "Get Blueprint Methods," you’ll see a JSON output listing all introspected methods and their parameter structures.

Click "Run Benchmark" to execute LayerForge benchmarks, with results displaying in the browser for performance verification.

## Troubleshooting

- **CORS Errors**: If you encounter CORS issues, ensure that you are serving the files over HTTP or HTTPS and not opening `index.html` directly as a file (`file:///...`).
- **Missing `wasm_exec.js`**: The `wasm_exec.js` file is essential for running Go’s WASM in the browser. Ensure it’s in the same directory as `index.html`.

