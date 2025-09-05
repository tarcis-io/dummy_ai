/*
 * "use strict" enforces stricter parsing in JavaScript.
 */
"use strict";

/*
 * wasmRun fetches, instantiates and runs a Go WebAssembly module.
 *
 * @param {string} wasmPath - Path to the Go WebAssembly module file.
 * @returns {void}
 */
const wasmRun = wasmPath => {
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch(wasmPath), go.importObject).then(webAssemblyInstantiatedSource => {
		go.run(webAssemblyInstantiatedSource.instance);
	});
};
