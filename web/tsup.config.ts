import { defineConfig } from 'tsup'

export default defineConfig({
  esbuildOptions(options, _context) {
    // Preserve directory structure
    // E.G. `./src-gen/foo/bar.ts` -> `./dist/foo/bar.ts`
    options.outbase = './src-gen'
  }
})
