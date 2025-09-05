"use strict";

const wasmRun = wasmPath => {
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch(wasmPath), go.importObject).then(webAssemblyInstantiatedSource => {
		go.run(webAssemblyInstantiatedSource.instance);
	});
};
