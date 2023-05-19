module.exports = {
  env: {
    browser: true,
    es2021: true
  },
  extends: [
    'plugin:react/recommended',
    'standard-with-typescript',
    'eslint:recommended',
    'prettier'
  ],
  overrides: [
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module'
  },
  plugins: [
    'react'
  ],
  rules: {
    "import/no-default-export": "off",
    "no-use-before-define": "off",    
    "react/destructuring-assignment": "off",
    "react/prop-types": "off",
    "react/require-default-props": "off",
    "react-hooks/exhaustive-deps": "off",
    "@typescript-eslint/no-floating-promises": "off",
    "@typescript-eslint/lines-between-class-members": "off",
    "@typescript-eslint/no-use-before-define": ["error"],
  }
}
