import { resolve } from "node:path";
import { defineConfig } from "vite";

// @ts-expect-error
import handlebars from "vite-plugin-handlebars";

// import pageData from "./pagedata.json";

// @ts-expect-error
import pageData from "./pagedata.cjs";

export default defineConfig(({ command }) => {
  const isBuild = command == "build";
  return {
    plugins: [
      handlebars({
        partialDirectory: resolve(__dirname, "./partials"),
        context: (pagePath: string) => {
          // return { ...pageData[pagePath], develop: !isBuild };
          return { ...pageData[pagePath](isBuild), develop: !isBuild };
        },
      }),
    ],
  };
});
