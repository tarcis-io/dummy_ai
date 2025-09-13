/*
 * "use strict" enforces stricter parsing in JavaScript.
 */
"use strict";

/*
 * wasmRun fetches, instantiates and runs a WebAssembly module
 * compiled with Go.
 *
 * @param {string} wasmPath - The path to the Go WebAssembly module.
 * @returns {void}
 */
const wasmRun = wasmPath => {
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch(wasmPath), go.importObject).then(webAssemblyInstantiatedSource => {
		go.run(webAssemblyInstantiatedSource.instance);
	});
};
