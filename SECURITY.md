# Security Policy

## Supported versions

This repository is a starter project. Security fixes should target the current `main` branch.

## Reporting a vulnerability

Please open a private report or contact the maintainer directly. Do not publish exploit details before the issue is fixed.

## Baseline rules

- Do not commit secrets, tokens, private keys, or production database URLs.
- Keep `.env` local and use `.env.example` for documentation.
- Use parameterized SQL queries.
- Do not log passwords, tokens, private user data, or full query strings.
- Store password hashes only. Raw password handling and authentication are outside the current project scope.
