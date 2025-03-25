const eslintPluginTs = require("@typescript-eslint/eslint-plugin");
const parserTs = require("@typescript-eslint/parser");

/** @type {import("eslint").Linter.Config[]} */
module.exports = [
  {
    files: ["**/*.ts"],
    languageOptions: {
      parser: parserTs,
      sourceType: "module"
    },
    plugins: {
      "@typescript-eslint": eslintPluginTs
    },
    rules: {
      semi: ["error", "always"],
      quotes: ["error", "double"],
      "@typescript-eslint/no-unused-vars": "warn"
    }
  }
];
