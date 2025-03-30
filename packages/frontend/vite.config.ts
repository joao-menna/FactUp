import tsconfigPaths from "vite-tsconfig-paths";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths()],

  server: {
    proxy: {
      "/api": {
        target: "http://localhost",
        followRedirects: true,
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
