# Release Process

This document describes how to create a new release of code-hiit.

## Prerequisites

- Push access to the main repository
- All changes committed and pushed to `main` branch
- Tests passing (run `make test`)

## Release Steps

### 1. Choose a Version Number

Follow [Semantic Versioning](https://semver.org/):
- **Alpha releases**: `v0.1.0-alpha.1`, `v0.1.0-alpha.2`, etc.
- **Beta releases**: `v0.1.0-beta.1`, `v0.1.0-beta.2`, etc.
- **Release candidates**: `v0.1.0-rc.1`, `v0.1.0-rc.2`, etc.
- **Stable releases**: `v1.0.0`, `v1.1.0`, `v2.0.0`, etc.

### 2. Create and Push a Git Tag

```bash
# Create a tag
git tag -a v0.1.0-alpha.1 -m "Release v0.1.0-alpha.1"

# Push the tag
git push origin v0.1.0-alpha.1
```

### 3. Wait for GitHub Actions

Once you push the tag:
1. GitHub Actions will automatically trigger the release workflow
2. It will build binaries for supported platforms:
   - Linux AMD64 (x86_64)
   - macOS ARM64 (Apple Silicon)
3. It will create a GitHub Release with all binaries attached
4. It will generate SHA256 checksums for verification

You can monitor the workflow at: https://github.com/dlm/code-hiit/actions

### 4. Update Homebrew Tap

After the release is published:

1. Download the checksums from the release:
   ```bash
   curl -fsSL https://github.com/dlm/code-hiit/releases/download/v0.1.0-alpha.1/checksums.txt
   ```

2. Update the formula in the [homebrew-tap repository](https://github.com/dlm/homebrew-tap):
   - Update `version` to match the new release (e.g., `"0.1.0-alpha.1"`)
   - Update the `sha256` values with the checksums from `checksums.txt`

3. Commit and push:
   ```bash
   git add code-hiit.rb
   git commit -m "Update code-hiit to v0.1.0-alpha.1"
   git push
   ```

## Testing the Release

### Test the Install Script

```bash
# Download and run the install script
curl -fsSL https://raw.githubusercontent.com/dlm/code-hiit/main/install.sh | sh

# Verify installation
code-hiit --version
```

### Test Manual Installation

```bash
# Download a binary
curl -LO https://github.com/dlm/code-hiit/releases/download/v0.1.0-alpha.1/code-hiit-linux-amd64

# Verify checksum
curl -fsSL https://github.com/dlm/code-hiit/releases/download/v0.1.0-alpha.1/checksums.txt
sha256sum code-hiit-linux-amd64

# Make executable and test
chmod +x code-hiit-linux-amd64
./code-hiit-linux-amd64 --version
```

### Test Homebrew (Once Tap is Set Up)

```bash
# Install from tap
brew install dlm/tap/code-hiit

# Verify
code-hiit --version

# Cleanup
brew uninstall code-hiit
```

## Troubleshooting

### Release Workflow Fails

1. Check the GitHub Actions logs at https://github.com/dlm/code-hiit/actions
2. Common issues:
   - Missing permissions (ensure `GITHUB_TOKEN` has write access)
   - Build failures (test with `make build` locally first)

### Install Script Fails

1. Test the script locally before promoting the release
2. Check that binaries are properly attached to the release
3. Verify URLs in the install script match the actual release URLs

### Homebrew Formula Issues

1. Test the formula locally: `brew audit --strict code-hiit.rb`
2. Ensure SHA256 checksums match exactly
3. Verify the version number is correct

## Post-Release Checklist

- [ ] Test installation on at least one Linux and one macOS system
- [ ] Verify `--version` flag shows correct version
- [ ] Update homebrew-tap repository (if applicable)
- [ ] Announce the release (social media, Discord, etc.)
- [ ] Update any documentation that references version numbers

## Versioning Strategy for Alpha

For the alpha phase:
- Start with `v0.1.0-alpha.1`
- Increment the alpha number for each release: `alpha.2`, `alpha.3`, etc.
- Once ready for beta, move to `v0.1.0-beta.1`
- Once stable, release `v0.1.0` or `v1.0.0`
