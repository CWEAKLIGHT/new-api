name: Build new-api (patched)

on:
  workflow_dispatch:
    inputs:
      ref:
        description: 'Branch or tag to build'
        required: false
        default: 'main'
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: read

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.27'
          cache: true

      - name: Setup Bun
        uses: oven-sh/setup-bun@v2
        with:
          bun-version: latest

      - name: Build frontend
        working-directory: web
        run: |
          bun install --registry https://registry.npmmirror.com
          bun run build

      - name: Verify frontend output
        run: |
          test -f web/dist/index.html || (echo "web/dist/index.html missing" && exit 1)
          echo "dist files: $(find web/dist -type f | wc -l)"

      - name: Build backend (windows/amd64)
        env:
          CGO_ENABLED: '0'
          GOOS: windows
          GOARCH: amd64
        run: |
          go build -trimpath -ldflags "-s -w" -o dist-out/new-api-patched.exe .

      - name: Show artifact info
        run: |
          ls -lh dist-out/
          file dist-out/new-api-patched.exe || true

      - name: Upload artifact
        uses: actions/upload-artifact@v4
        with:
          name: new-api-windows-amd64
          path: dist-out/new-api-patched.exe
          retention-days: 14
          if-no-files-found: error
