# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Full support for SLURM REST API v0.0.44 (SLURM 25.11.0)
  - Complete adapter implementation for all resource managers
  - Comprehensive test coverage with 100% passing tests
  - Type-safe conversions between v0.0.44 API types and common types
  - Error handling with status code inclusion in error messages
  - Validation for all CRUD operations
- Enhanced error messages to include HTTP status codes for better debugging
- Cluster adapter validation for resource names and required fields
- Improved validation patterns across all v0.0.44 adapters

### Changed
- Updated multi-version support range from v0.0.40-v0.0.43 to v0.0.40-v0.0.44
- Enhanced error adapter to include status codes (404, 409, 500) in error messages
- Updated README.md to reflect v0.0.44 support
- Improved .gitignore patterns for temporary files and reports

### Fixed
- Case sensitivity issues in validation error messages
- Type conversion for PartitionState in partition adapter
- Empty update validation to require at least one field
- Reservation validation to require both StartTime and EndTime

## Previous Releases

See Git history for changes prior to this release.

---

## Version Support Matrix

| SLURM Version | REST API Version | Support Status |
|---------------|------------------|----------------|
| 25.11.0       | v0.0.44          | ✅ Supported    |
| 25.05.0       | v0.0.43          | ✅ Supported    |
| 24.11.0       | v0.0.42          | ✅ Supported    |
| 24.05.0       | v0.0.41          | ✅ Supported    |
| 23.11.0       | v0.0.40          | ✅ Supported    |

---

*For detailed changes, see the [commit history](https://github.com/jontk/slurm-client/commits/main).*
