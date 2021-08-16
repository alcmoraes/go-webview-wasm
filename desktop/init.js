if (!WebAssembly.instantiateStreaming) { // polyfill
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

(async function() {
  try {
    const go = new Go()
    const b64 = await getWASM()
    const buf = Uint8Array.from(atob(b64), c => c.charCodeAt(0)).buffer
    const { instance } = await WebAssembly.instantiate(buf, go.importObject)
    go.run(instance);
  } catch(ERR){
    console.log(ERR);
  }
})()