import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  build: {
    lib: {
      entry: resolve(__dirname, 'src/index.ts'),
      name: 'PrismUI',
      fileName: 'prism-ui',
      formats: ['es'],
    },
    rollupOptions: {
      // Don't externalize lit - bundle it
      external: [],
    },
    minify: 'esbuild',
    sourcemap: true,
  },
  define: {
    'process.env.NODE_ENV': JSON.stringify('production'),
  },
});
