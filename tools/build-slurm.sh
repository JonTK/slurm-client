#!/bin/bash

# Build SLURM with REST API v0.0.44 support
# This script helps build SLURM from source with the necessary plugins

set -e

# Configuration
SLURM_VERSION="${SLURM_VERSION:-25.11.1}"
SLURM_DIR="${SLURM_DIR:-./slurm-build}"
INSTALL_PREFIX="${INSTALL_PREFIX:-/usr/local/slurm}"
PARALLEL_JOBS="${PARALLEL_JOBS:-$(nproc)}"

echo "🔨 Building SLURM $SLURM_VERSION with REST API v0.0.44 support"
echo "Build directory: $SLURM_DIR"
echo "Install prefix: $INSTALL_PREFIX"
echo "Parallel jobs: $PARALLEL_JOBS"
echo ""

# Check dependencies
echo "Checking build dependencies..."
MISSING_DEPS=()

command -v git >/dev/null 2>&1 || MISSING_DEPS+=("git")
command -v make >/dev/null 2>&1 || MISSING_DEPS+=("make")
command -v gcc >/dev/null 2>&1 || MISSING_DEPS+=("gcc")
command -v autoconf >/dev/null 2>&1 || MISSING_DEPS+=("autoconf")
command -v automake >/dev/null 2>&1 || MISSING_DEPS+=("automake")
command -v libtool >/dev/null 2>&1 || MISSING_DEPS+=("libtool")
command -v pkg-config >/dev/null 2>&1 || MISSING_DEPS+=("pkg-config")

# Check for development libraries
# Skip pkg-config check in NixOS environment or when explicitly disabled
if [ -z "$NIXOS_BUILD" ] && [ -z "$IN_NIX_SHELL" ]; then
    if ! pkg-config --exists json-c; then
        MISSING_DEPS+=("json-c-dev")
    fi

    if ! pkg-config --exists yaml-0.1; then
        MISSING_DEPS+=("libyaml-dev")
    fi
fi

if [ ${#MISSING_DEPS[@]} -ne 0 ]; then
    echo "❌ Missing dependencies: ${MISSING_DEPS[*]}"
    echo ""
    echo "On Ubuntu/Debian, install with:"
    echo "  sudo apt-get update"
    echo "  sudo apt-get install git build-essential autoconf automake libtool pkg-config"
    echo "  sudo apt-get install libmunge-dev libmunge2 munge"
    echo "  sudo apt-get install libjson-c-dev libyaml-dev libhttp-parser-dev"
    echo "  sudo apt-get install libcurl4-openssl-dev libssl-dev"
    echo ""
    echo "On RedHat/CentOS/Rocky, install with:"
    echo "  sudo yum groupinstall 'Development Tools'"
    echo "  sudo yum install git autoconf automake libtool pkg-config"
    echo "  sudo yum install munge-devel munge"
    echo "  sudo yum install json-c-devel libyaml-devel http-parser-devel"
    echo "  sudo yum install libcurl-devel openssl-devel"
    exit 1
fi

echo "✓ All dependencies found"
echo ""

# Create build directory
mkdir -p "$SLURM_DIR"
cd "$SLURM_DIR"

# Clone or update SLURM source
if [ ! -d "slurm" ]; then
    echo "📥 Cloning SLURM source..."
    git clone https://github.com/SchedMD/slurm.git
else
    echo "📥 Updating SLURM source..."
    cd slurm
    git fetch
    cd ..
fi

cd slurm

# Checkout specific version if specified
if [ "$SLURM_VERSION" != "latest" ]; then
    echo "📌 Checking out SLURM version $SLURM_VERSION..."
    git checkout "slurm-$SLURM_VERSION" || {
        echo "⚠ Version $SLURM_VERSION not found, using latest"
        git checkout master
    }
else
    git checkout master
fi

# Get current version info
CURRENT_VERSION=$(git describe --tags --always)
echo "📋 Building SLURM version: $CURRENT_VERSION"
echo ""

# Generate configure script
echo "🔧 Generating configure script..."
if [ ! -f configure ]; then
    ./autogen.sh
fi

# Configure build
echo "⚙️ Configuring build with REST API support..."
./configure \
    --prefix="$INSTALL_PREFIX" \
    --enable-rest \
    --enable-openapi \
    --with-json \
    --with-yaml \
    --with-http-parser \
    --enable-shared \
    --disable-static \
    --sysconfdir="$INSTALL_PREFIX/etc" \
    --localstatedir="$INSTALL_PREFIX/var"

echo ""
echo "🏗️ Building SLURM..."
make -j"$PARALLEL_JOBS"

echo ""
echo "✅ Build completed successfully!"
echo ""
echo "📦 To install SLURM:"
echo "  sudo make install"
echo ""
echo "🔧 To configure SLURM, create configuration files in $INSTALL_PREFIX/etc/"
echo ""
echo "🚀 To start slurmrestd with v0.0.44 support:"
echo "  $INSTALL_PREFIX/sbin/slurmrestd \\"
echo "    -s openapi/slurmctld,openapi/slurmdbd,openapi/util \\"
echo "    -s data_parser/v0.0.44 \\"
echo "    localhost:6820"
echo ""
echo "🧪 To test the REST API:"
echo "  curl http://localhost:6820/openapi/v3"
echo "  curl http://localhost:6820/openapi/v0.0.44"
echo ""

# Check if v0.0.44 plugin was built
V044_PLUGIN=$(find . -name "*v0_0_44*" -type f 2>/dev/null | head -1)
if [ -n "$V044_PLUGIN" ]; then
    echo "✅ v0.0.44 plugin found: $V044_PLUGIN"
else
    echo "⚠ v0.0.44 plugin not found - may not be available in this SLURM version"
fi

echo ""
echo "📚 Next steps:"
echo "1. Install SLURM: sudo make install"
echo "2. Configure SLURM cluster"
echo "3. Start slurmctld and slurmdbd"
echo "4. Start slurmrestd with v0.0.44 plugin"
echo "5. Run: make download-specs to get OpenAPI specifications"