import { defineConfig } from "@rsbuild/core";
import { pluginReact } from "@rsbuild/plugin-react";

export default defineConfig({
    plugins: [pluginReact()],
    source: {
        entry: {
            index: "./uiweb/index.jsx",
        }
    },
    output: {
        target: 'web',
        distPath: {
            root: 'uiweb/dist'
        }
    },
    html: {
        title: "Harmony Cook"
    }
})