import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'

export default defineConfig({
  plugins: [solid()],
  server: {
      // Allow access from outside (needed when developing on a remote machine)
      host: '0.0.0.0',
      port: 3001,
  },
  build: {
    outDir: 'dist',
  },
})


