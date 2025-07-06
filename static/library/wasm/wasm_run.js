"use strict";

const wasmRun = (wasmRoute) => {
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch(wasmRoute), go.importObject).then((webAssemblyInstantiatedSource) => {
		go.run(webAssemblyInstantiatedSource.instance);
	});
};
