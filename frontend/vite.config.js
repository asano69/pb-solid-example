import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [solid()],
  server: {
    // Allow access from outside (needed when developing on a remote machine)
    host: "0.0.0.0",
    port: 3001,
    allowedHosts: true,
    changeOrigin: true,
    proxy: {
      "/api": "http://127.0.0.1:3000",
      "/_": "http://127.0.0.1:3000",
    },
  },
  build: {
    outDir: "dist",
  },
});


