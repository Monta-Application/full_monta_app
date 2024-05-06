import react from "@vitejs/plugin-react-swc";
import million from "million/compiler";
import path from "path";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig, loadEnv, type PluginOption } from "vite";

export default defineConfig(({ command, mode }) => {
  // Load env file based on `mode` in the current working directory.
  // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
  const env = loadEnv(mode, process.cwd(), "");

  return {
    plugins: [
      million.vite({ auto: true }),
      react(),
      visualizer({
        open: true,
        gzipSize: true,
        brotliSize: true,
      }) as PluginOption,
    ],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    define: {
      __APP_ENV__: JSON.stringify(env.APP_ENV),
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes("node_modules")) {
              const packages = [
                "lodash",
                "date-fns",
                "react-hook-form",
                "@radix-ui",
                "react-aria",
                "react-dom",
                "@fortawesome",
              ];
              const chunk = packages.find((pkg) =>
                id.includes(`/node_modules/${pkg}`),
              );
              return chunk ? `vendor.${chunk}` : null;
            }
          },
        },
      },
    },
  };
});
