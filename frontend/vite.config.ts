import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

const rootPath = new URL(".", import.meta.url).pathname.substring(1);
console.log({ rootPath, url: import.meta.url });

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": rootPath + "src",
      "@wails": rootPath + "wailsjs",
    },
  },
});
