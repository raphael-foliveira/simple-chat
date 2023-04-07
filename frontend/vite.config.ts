import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/

const { VITE_API_HTTP_URL } = process.env;
export default defineConfig({
  plugins: [react()],
});
