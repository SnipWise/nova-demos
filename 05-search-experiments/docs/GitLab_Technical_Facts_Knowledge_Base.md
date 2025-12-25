# GitLab Technical Facts Knowledge Base

## Resolution Reports for Common Issues

**Document Version:** 2.0  
**Last Updated:** December 2024  
**Total Facts:** 50+

---

## Table of Contents

1. [Overview](#overview)
2. [CI/CD Pipeline Issues](#cicd-pipeline-issues)
3. [Security Vulnerabilities (CVE)](#security-vulnerabilities-cve)
4. [GitLab Runner Issues](#gitlab-runner-issues)
5. [Container Registry Issues](#container-registry-issues)
6. [Merge Request Issues](#merge-request-issues)
7. [API & Rate Limiting Issues](#api--rate-limiting-issues)
8. [GitLab Pages Issues](#gitlab-pages-issues)
9. [Kubernetes Integration Issues](#kubernetes-integration-issues)
10. [Infrastructure & Performance Issues](#infrastructure--performance-issues)
11. [Authentication & Authorization Issues](#authentication--authorization-issues)
12. [Statistics](#statistics)
13. [References](#references)

---

## Overview

This knowledge base compiles technical facts frequently encountered on GitLab, including security vulnerabilities (CVEs), CI/CD configuration issues, GitLab Runner errors, and difficulties related to merge requests, container registry, and Kubernetes integration. Each technical fact is accompanied by a detailed resolution report.

**Sources:** GitLab Official Documentation, GitLab Issue Tracker, HackerOne Bug Bounty Program, GitLab Security Advisories, Community Forums.

---

## CI/CD Pipeline Issues

### FACT-001: YAML Invalid Error - Pipeline Blocked

| Field | Value |
|-------|-------|
| **ID** | FACT-001 |
| **Category** | CI/CD Pipeline |
| **Severity** | High |
| **CVE** | N/A |

**Symptoms:**
- "yaml invalid" badge displayed
- Pipeline blocked without execution
- No jobs created

**Root Cause:**
Incorrect YAML syntax, mixing tabs and spaces, improper indentation, or malformed structure in `.gitlab-ci.yml`.

**Resolution Steps:**
1. Use GitLab's Pipeline Editor with automatic validation
2. Validate syntax with CI Lint tool (Settings > CI/CD)
3. Convert all tabs to spaces (YAML requires spaces only)
4. Validate locally with `yamllint`
5. Check for missing colons, incorrect nesting, or duplicate keys

---

### FACT-002: Configuration Too Large - Error 500

| Field | Value |
|-------|-------|
| **ID** | FACT-002 |
| **Category** | CI/CD Pipeline |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error 500 when editing `.gitlab-ci.yml`
- Frequent timeouts during pipeline creation
- Web editor becomes unresponsive

**Root Cause:**
CI/CD configuration file exceeds size limits or contains circular include loops.

**Resolution Steps:**
1. Check configuration length in Pipeline Editor > Full configuration tab
2. Remove duplicated configurations
3. Move long scripts to standalone files in the project
4. Use parent/child pipelines to distribute workload
5. On self-managed instances, increase size limits in admin settings

---

### FACT-003: Variables Not Resolved in Rules

| Field | Value |
|-------|-------|
| **ID** | FACT-003 |
| **Category** | CI/CD Pipeline |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Jobs not added to pipeline despite correct conditions
- Pipeline runs but expected jobs are missing
- Unexpected pipeline behavior

**Root Cause:**
Incorrect usage of predefined variables (`CI_PIPELINE_SOURCE`, `CI_MERGE_REQUEST_ID`) or mixing `only/except` with `rules`.

**Resolution Steps:**
1. Verify variable values with a debug job
2. Check common `if` clauses in documentation
3. Avoid mixing `only/except` with `rules` (different behaviors)
4. Use `workflow:rules` for global pipeline control
5. Export variables list to identify missing or unexpected values

---

### FACT-004: Duplicate Pipelines Running

| Field | Value |
|-------|-------|
| **ID** | FACT-004 |
| **Category** | CI/CD Pipeline |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Two pipelines execute for a single push
- Branch pipeline AND merge request pipeline both trigger
- Resource waste and confusion

**Root Cause:**
Using `when` clause without `if` clause in rules, triggering both push and MR pipelines.

**Resolution Steps:**
1. Add `workflow:rules` to control pipeline triggers
2. Specify `if: $CI_PIPELINE_SOURCE == 'merge_request_event'`
3. Use `rules:if` instead of standalone `when`
4. Configure `CI_OPEN_MERGE_REQUESTS` for MR detection
5. Review rules logic with pipeline simulation tool

---

### FACT-005: Downstream Pipeline Creation Failed

| Field | Value |
|-------|-------|
| **ID** | FACT-005 |
| **Category** | CI/CD Pipeline |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Trigger job fails without creating downstream pipeline
- Error: "downstream pipeline can not be created"
- Ambiguous ref errors

**Root Cause:**
Downstream project not found, insufficient permissions, protected branch restrictions, or ambiguous tag/branch names.

**Resolution Steps:**
1. Verify downstream project exists and is accessible
2. Check user permissions in downstream project
3. Add downstream project to job token scope allowlist
4. Ensure tag names don't match branch names
5. Review protected branch settings for the user

---

### FACT-006: Job Configuration Not Updated After Changes

| Field | Value |
|-------|-------|
| **ID** | FACT-006 |
| **Category** | CI/CD Pipeline |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Rerun job uses old configuration
- Changes to `.gitlab-ci.yml` not reflected
- Included files not updated

**Root Cause:**
Pipeline configuration is fetched only at creation time. Rerunning jobs uses the same configuration.

**Resolution Steps:**
1. Start a new pipeline to use updated configuration
2. Don't rely on job retry for configuration changes
3. Use manual pipeline trigger after config changes
4. Verify included files are also updated
5. Clear any cached configuration if applicable

---

### FACT-007: cURL HTTP/2 Error During Git Fetch

| Field | Value |
|-------|-------|
| **ID** | FACT-007 |
| **Category** | CI/CD Pipeline |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Error: `curl 16 HTTP/2 send again with decreased length`
- Git fetch fails intermittently
- RPC failures during clone

**Root Cause:**
Incompatibility between HTTP/2 protocol and libcurl/Git in certain environments.

**Resolution Steps:**
1. Configure Git to use HTTP/1.1 in runner's `config.toml`
2. Add environment variables:
   - `GIT_CONFIG_COUNT=1`
   - `GIT_CONFIG_KEY_0=http.version`
   - `GIT_CONFIG_VALUE_0=HTTP/1.1`
3. Update Git and libcurl versions
4. Check proxy configuration if applicable

---

### FACT-008: Protected Manual Job Access Denied

| Field | Value |
|-------|-------|
| **ID** | FACT-008 |
| **Category** | CI/CD Pipeline |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Manual job fails when run by administrator
- Clone fails in private project
- Permission denied for protected jobs

**Root Cause:**
Administrators must be direct members of private projects to clone source code, even for protected manual jobs.

**Resolution Steps:**
1. Add administrator as direct project member
2. Impersonate a user who is a direct member
3. Review protected branch permissions
4. Check project visibility settings
5. Verify job token permissions

---

## Security Vulnerabilities (CVE)

### FACT-009: CVE-2024-5655 - Pipeline Execution as Arbitrary User

| Field | Value |
|-------|-------|
| **ID** | FACT-009 |
| **Category** | Security |
| **Severity** | **Critical (CVSS 9.6)** |
| **CVE** | CVE-2024-5655 |

**Symptoms:**
- Attacker can trigger CI/CD jobs under another user's identity
- Unauthorized pipeline execution
- Potential supply chain compromise

**Affected Versions:**
GitLab CE/EE 15.8 to 17.1

**Resolution Steps:**
1. **Immediately update** to GitLab 17.1.1, 17.0.3, or 16.11.5
2. Audit pipeline execution logs for suspicious activity
3. Review CI_JOB_TOKEN permissions
4. Enable job token scope restrictions
5. Monitor for unauthorized pipeline triggers

**Breaking Changes After Patch:**
- MR Retargeting Workflow Pipelines no longer run automatically
- GraphQL CI_JOB_TOKEN authentication disabled by default

---

### FACT-010: CVE-2024-6385 - Pipeline Execution Under Arbitrary Identity

| Field | Value |
|-------|-------|
| **ID** | FACT-010 |
| **Category** | Security |
| **Severity** | **Critical (CVSS 9.6)** |
| **CVE** | CVE-2024-6385 |

**Symptoms:**
- Attacker can execute pipelines as any user
- Similar to CVE-2024-5655
- CI/CD infrastructure compromise risk

**Resolution Steps:**
1. Apply patches immediately (versions 17.1.1+)
2. Audit pipeline executions from the last 90 days
3. Revoke and regenerate potentially compromised tokens
4. Implement CI/CD environment separation
5. Enable enhanced audit logging

---

### FACT-011: CVE-2024-6678 - Arbitrary User Pipeline Trigger

| Field | Value |
|-------|-------|
| **ID** | FACT-011 |
| **Category** | Security |
| **Severity** | **Critical (CVSS 9.9)** |
| **CVE** | CVE-2024-6678 |

**Symptoms:**
- Attackers can run pipeline jobs as arbitrary users
- Fourth similar vulnerability in one year
- Affects all versions from 8.14

**Affected Versions:**
GitLab CE/EE 8.14 prior to 17.1.7, 17.2 prior to 17.2.5, 17.3 prior to 17.3.2

**Resolution Steps:**
1. Upgrade to versions 17.3.2, 17.2.5, or 17.1.7
2. Implement Role-Based Access Control (RBAC)
3. Audit user permissions thoroughly
4. Enforce strict segmentation between environments
5. Monitor pipeline activity for anomalies

---

### FACT-012: CVE-2024-4835 - XSS in Web IDE

| Field | Value |
|-------|-------|
| **ID** | FACT-012 |
| **Category** | Security |
| **Severity** | High |
| **CVE** | CVE-2024-4835 |

**Symptoms:**
- Account takeover possible via malicious page
- XSS vulnerability in VS Code editor (Web IDE)
- Requires user interaction

**Resolution Steps:**
1. Update to GitLab 17.0.1, 16.11.3, or 16.10.6
2. Train users on malicious page risks
3. Enable strict CSP policies
4. Monitor for suspicious account activities
5. Implement additional authentication factors

---

### FACT-013: CVE-2024-9164 - Pipeline on Arbitrary Branches

| Field | Value |
|-------|-------|
| **ID** | FACT-013 |
| **Category** | Security |
| **Severity** | **Critical** |
| **CVE** | CVE-2024-9164 |

**Symptoms:**
- Malicious actors can run pipelines on any branch
- Project integrity compromise
- Data tampering risk

**Affected Versions:**
All versions from 12.5 prior to latest patch releases

**Resolution Steps:**
1. Update to latest patched version immediately
2. Review branch protection settings
3. Implement strict branch access controls
4. Monitor for unauthorized branch activity
5. Enable protected branches for sensitive code

---

### FACT-014: CVE-2024-4994 - CSRF on GraphQL API

| Field | Value |
|-------|-------|
| **ID** | FACT-014 |
| **Category** | Security |
| **Severity** | High |
| **CVE** | CVE-2024-4994 |

**Symptoms:**
- Attackers can execute arbitrary GraphQL mutations
- Unauthorized project settings updates
- Pipeline triggers without authorization

**Resolution Steps:**
1. Apply latest security patches
2. Implement additional CSRF protections
3. Review GraphQL mutation permissions
4. Monitor API activity logs
5. Enable additional authentication for sensitive operations

---

### FACT-015: CVE-2024-45409 - SAML Authentication Bypass

| Field | Value |
|-------|-------|
| **ID** | FACT-015 |
| **Category** | Security |
| **Severity** | **Critical (CVSS 10.0)** |
| **CVE** | CVE-2024-45409 |

**Symptoms:**
- Unauthenticated attacker gains instance access
- SAML authentication completely bypassed
- Affects self-hosted instances with SAML

**Root Cause:**
Vulnerability in Ruby libraries handling SAML authentication requests.

**Resolution Steps:**
1. Upgrade to latest GitLab version immediately
2. Review SAML configuration settings
3. Audit access logs for suspicious authentications
4. Consider temporary SAML disabling if patch not available
5. Implement additional authentication factors

---

### FACT-016: CVE-2024-9183 - CI/CD Cache Race Condition

| Field | Value |
|-------|-------|
| **ID** | FACT-016 |
| **Category** | Security |
| **Severity** | High |
| **CVE** | CVE-2024-9183 |

**Symptoms:**
- Authenticated user obtains higher-privileged credentials
- Race condition in CI/CD cache
- Privilege escalation possible

**Resolution Steps:**
1. Update to GitLab 18.6.1, 18.5.3, or 18.4.5
2. Review cache configurations
3. Implement least privilege access
4. Monitor for unusual credential usage
5. Audit pipeline cache access patterns

---

### FACT-017: CVE-2025-12571 - JSON DoS Attack

| Field | Value |
|-------|-------|
| **ID** | FACT-017 |
| **Category** | Security |
| **Severity** | High |
| **CVE** | CVE-2025-12571 |

**Symptoms:**
- Unauthenticated denial-of-service attacks
- Instance crashes from malicious JSON payloads
- No authentication required

**Resolution Steps:**
1. Apply security patches immediately
2. Implement rate limiting on JSON endpoints
3. Configure WAF rules for JSON payload inspection
4. Monitor for unusual request patterns
5. Set up alerting for service disruptions

---

### FACT-018: CVE-2024-8640 - Cube Server Command Injection

| Field | Value |
|-------|-------|
| **ID** | FACT-018 |
| **Category** | Security |
| **Severity** | High (CVSS 8.5) |
| **CVE** | CVE-2024-8640 |

**Symptoms:**
- Remote code execution possible
- Command injection in connected Cube server
- Incomplete input filtering

**Resolution Steps:**
1. Update GitLab EE to patched version
2. Review Cube server integration settings
3. Implement strict input validation
4. Isolate Cube server connections
5. Monitor for unusual command execution

---

### FACT-019: CVE-2025-25291/CVE-2025-25292 - SAML Bypass

| Field | Value |
|-------|-------|
| **ID** | FACT-019 |
| **Category** | Security |
| **Severity** | High (CVSS 8.8) |
| **CVE** | CVE-2025-25291, CVE-2025-25292 |

**Symptoms:**
- SAML authentication bypass
- Requires compromised valid user account
- Instance access gained

**Resolution Steps:**
1. Apply critical patches immediately
2. Audit user account security
3. Implement additional authentication factors
4. Review SAML provider configuration
5. Monitor for unusual login patterns

---

## GitLab Runner Issues

### FACT-020: SSH Authentication Handshake Failed

| Field | Value |
|-------|-------|
| **ID** | FACT-020 |
| **Category** | GitLab Runner |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `ssh: handshake failed: ssh: unable to authenticate`
- SSH executor jobs fail
- Connection refused during prepare stage

**Root Cause:**
SSH keys misconfigured, `known_hosts` file absent, or key mismatch.

**Resolution Steps:**
1. Create `/root/.ssh/known_hosts` if missing
2. Perform manual SSH connection to accept fingerprint
3. Add `disable_strict_host_key_checking = true` in `config.toml`
4. Verify private key matches authorized public key
5. Check SSH key permissions (600 for private key)

---

### FACT-021: Docker Image Pull Failed - Permissions

| Field | Value |
|-------|-------|
| **ID** | FACT-021 |
| **Category** | GitLab Runner |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `Failed to pull image` intermittent by user
- Bot users frequently affected
- Image available for some users, not others

**Root Cause:**
Different permissions between users for registry access, combined with runner image caching behavior.

**Resolution Steps:**
1. Verify access permissions for all users including bots
2. Configure job token scope allowlist for downstream projects
3. Grant read permissions to service accounts
4. Use public images or configure registry authentication
5. Clear runner cache if permissions changed

---

### FACT-022: Runner Authentication Failed After Registration

| Field | Value |
|-------|-------|
| **ID** | FACT-022 |
| **Category** | GitLab Runner |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Runner registers successfully but jobs fail
- Authentication errors in runner logs
- Token appears valid but rejected

**Root Cause:**
Token corruption, runner version mismatch, or network issues between runner and GitLab.

**Resolution Steps:**
1. Re-register the runner with a new token
2. Verify runner version compatibility
3. Check network connectivity to GitLab instance
4. Review TLS/SSL certificate validity
5. Ensure runner clock is synchronized (NTP)

---

### FACT-023: SELinux Permission Denied

| Field | Value |
|-------|-------|
| **ID** | FACT-023 |
| **Category** | GitLab Runner |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `Permission Denied` with Docker executor
- Jobs fail during script execution
- Works when SELinux disabled

**Root Cause:**
SELinux policies blocking runner operations on CentOS, Fedora, RHEL systems.

**Resolution Steps:**
1. Check SELinux audit log for denials
2. Update SELinux policies to allow runner operations
3. Use `audit2allow` to generate custom policies
4. Consider SELinux boolean adjustments
5. As last resort, set SELinux to permissive mode

---

### FACT-024: Docker-in-Docker MTU Mismatch

| Field | Value |
|-------|-------|
| **ID** | FACT-024 |
| **Category** | GitLab Runner |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Network timeouts in DIND jobs
- Package downloads fail intermittently
- Works locally but fails in CI

**Root Cause:**
DIND default MTU (1500) larger than Kubernetes overlay network, causing packet fragmentation issues.

**Resolution Steps:**
1. Set DIND MTU to match overlay network (usually 1450)
2. Add `--mtu=1450` to Docker daemon configuration
3. Configure via CI variable: `DOCKER_MTU=1450`
4. Check Kubernetes CNI MTU settings
5. Test with smaller MTU values if issues persist

---

### FACT-025: Long Polling Deadlock

| Field | Value |
|-------|-------|
| **ID** | FACT-025 |
| **Category** | GitLab Runner |
| **Severity** | High |
| **CVE** | N/A |

**Symptoms:**
- Runner workers stuck in long polling
- Jobs not processed promptly
- Performance bottlenecks to complete deadlocks

**Root Cause:**
GitLab CI/CD long polling feature (`apiCiLongPollingDuration` default 50s) causing request queueing.

**Resolution Steps:**
1. Review GitLab Workhorse settings
2. Adjust `apiCiLongPollingDuration` value
3. Check runner concurrent job settings
4. Review advanced configuration documentation
5. Consider runner scaling adjustments

---

## Container Registry Issues

### FACT-026: 502 Bad Gateway on Image Push (GCS)

| Field | Value |
|-------|-------|
| **ID** | FACT-026 |
| **Category** | Container Registry |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- 502 Bad Gateway when pushing images
- CPU spikes during large image pushes
- GCS backend specific issue

**Root Cause:**
Registry communicating with GCS using HTTP/2 protocol causing issues.

**Resolution Steps:**
1. Disable HTTP/2 by setting `GODEBUG=http2client=0`
2. Add environment variable to registry deployment
3. Restart registry service
4. Monitor CPU usage after change
5. Consider alternative storage backend if issues persist

---

### FACT-027: Token Expiration During Large Push

| Field | Value |
|-------|-------|
| **ID** | FACT-027 |
| **Category** | Container Registry |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `unauthorized: authentication required`
- Large images fail to push
- Smaller images work fine

**Root Cause:**
Authentication token expires before push completes (default 5 minutes self-managed, 15 minutes GitLab.com).

**Resolution Steps:**
1. Increase token duration in Admin > Settings > CI/CD
2. Split large images into smaller layers
3. Optimize image size with multi-stage builds
4. Use faster network connections
5. Consider image compression techniques

---

### FACT-028: Special Characters in Path Cause Connection Error

| Field | Value |
|-------|-------|
| **ID** | FACT-028 |
| **Category** | Container Registry |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Docker connection errors
- Push/pull fails for specific projects
- Works for other projects

**Root Cause:**
Special characters in group, project, or branch names (leading underscore, trailing hyphen).

**Resolution Steps:**
1. Change group path to remove special characters
2. Change project path if needed
3. Rename branches with problematic characters
4. Create push rules to prevent future issues
5. Use URL-safe naming conventions

---

### FACT-029: Manifest Unknown - Old Docker Client

| Field | Value |
|-------|-------|
| **ID** | FACT-029 |
| **Category** | Container Registry |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- Error: `404 Not Found` or `Unknown Manifest`
- Using Docker Engine 17.11 or earlier
- Images exist but cannot be pulled

**Root Cause:**
Old Docker clients using v1 API while GitLab requires v2 API.

**Resolution Steps:**
1. Update Docker Engine to current version
2. Convert v1 images to v2 format
3. Rebuild images with modern Docker
4. Check client compatibility requirements
5. Update CI job Docker images

---

### FACT-030: Multi-arch Blob Unknown Error

| Field | Value |
|-------|-------|
| **ID** | FACT-030 |
| **Category** | Container Registry |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `manifest blob unknown: blob unknown to registry`
- Multi-architecture image build fails
- Individual platform images work

**Root Cause:**
Multiple architecture images spread across different repositories instead of same repository.

**Resolution Steps:**
1. Push all architecture images to same repository
2. Include architecture in individual image tags
3. Build multi-arch manifest from same repository
4. Use `docker buildx` for proper multi-arch builds
5. Verify all blobs uploaded before manifest creation

---

### FACT-031: x509 Certificate Error - Self-Signed

| Field | Value |
|-------|-------|
| **ID** | FACT-031 |
| **Category** | Container Registry |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `x509: certificate signed by unknown authority`
- Self-signed certificate not trusted
- CI jobs fail to pull images

**Root Cause:**
Docker daemon doesn't trust self-signed certificates by default.

**Resolution Steps:**
1. Mount Docker daemon socket in runner config
2. Set `privileged = false` in runner `config.toml`
3. Add CA certificate to trusted store
4. Configure Docker daemon to trust the certificate
5. Consider using Let's Encrypt for valid certificates

---

### FACT-032: Kaniko HTTP/2 Performance Issue

| Field | Value |
|-------|-------|
| **ID** | FACT-032 |
| **Category** | Container Registry |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Kaniko builds extremely slow
- Timeouts during image push
- High network latency

**Root Cause:**
Performance issue between Kaniko and HTTP/2 protocol.

**Resolution Steps:**
1. Set `GODEBUG=http2client=0` environment variable
2. Add to CI job variables
3. Use HTTP/1.1 for Kaniko pushes
4. Monitor build times after change
5. Consider alternative build tools if issues persist

---

## Merge Request Issues

### FACT-033: 403 Error on Approval Rules API

| Field | Value |
|-------|-------|
| **ID** | FACT-033 |
| **Category** | Merge Request |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- 403 Forbidden when modifying approval rules via API
- User has sufficient project permissions
- Works in UI but not API

**Root Cause:**
`Prevent editing approval rules in projects and merge requests` enabled at instance level.

**Resolution Steps:**
1. Check `disable_overriding_approvers_per_merge_request` via API
2. Contact GitLab administrator to modify instance rules
3. Disable prevention setting if appropriate
4. Use administrator account for modifications
5. Review approval rule hierarchy (instance > group > project)

---

### FACT-034: Conflict Resolution Returns 500

| Field | Value |
|-------|-------|
| **ID** | FACT-034 |
| **Category** | Merge Request |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error 500 when clicking "Fix conflict automatically"
- No user-friendly error message
- Conflicts exist but cannot be resolved

**Root Cause:**
Insufficient permissions on source project (fork) when resolving conflicts.

**Resolution Steps:**
1. Verify permissions on both source AND destination projects
2. Resolve conflicts locally via `git rebase` or `git merge`
3. Use `/rebase` quick action in MR comments
4. Request additional permissions from maintainer
5. Create new MR if permissions cannot be granted

---

### FACT-035: MR Pipeline Stuck - Checking Ability to Merge

| Field | Value |
|-------|-------|
| **ID** | FACT-035 |
| **Category** | Merge Request |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- MR widget stuck on "Checking ability to merge automatically"
- Pipeline configuration changed after MR creation
- No progress on merge status

**Root Cause:**
MR pipelines configuration removed from `.gitlab-ci.yml` after MR was created.

**Resolution Steps:**
1. Execute `/rebase` quick action in comment
2. Click "Rebase source branch" in widget
3. Reconfigure merge request pipelines in `.gitlab-ci.yml`
4. Close and recreate the merge request
5. Push a new commit to trigger pipeline refresh

---

### FACT-036: Developer Cannot Resolve Conflicts

| Field | Value |
|-------|-------|
| **ID** | FACT-036 |
| **Category** | Merge Request |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- "Resolve conflicts" button not available for Developer role
- Previously working functionality now restricted
- Must use command line resolution

**Root Cause:**
Permission changes or protected branch settings preventing Developer from pushing merge commits.

**Resolution Steps:**
1. Check protected branch settings for the target branch
2. Verify Developer role permissions haven't changed
3. Use local git merge to resolve conflicts
4. Request Maintainer assistance for protected branches
5. Review project-level merge request settings

---

### FACT-037: 403 Error Submitting MR to Public Repo

| Field | Value |
|-------|-------|
| **ID** | FACT-037 |
| **Category** | Merge Request |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- 403 error when submitting MR from fork
- Can create MR form but submission fails
- No clear indication of required permissions

**Root Cause:**
Hidden permission requirements for cross-project merge requests not clearly communicated.

**Resolution Steps:**
1. Verify fork is properly configured
2. Check upstream project MR settings
3. Request contributor access if required
4. Review project's contribution guidelines
5. Contact project maintainers for access

---

## API & Rate Limiting Issues

### FACT-038: 429 Too Many Requests

| Field | Value |
|-------|-------|
| **ID** | FACT-038 |
| **Category** | API |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- HTTP 429 response code
- "Retry later" message
- API calls blocked

**Root Cause:**
Exceeding rate limits for authenticated/unauthenticated API requests.

**Resolution Steps:**
1. Implement exponential backoff retry logic
2. Check `RateLimit-*` response headers
3. Stagger automated pipeline execution
4. Request rate limit increase from administrator
5. Use caching to reduce API calls

**Default Limits:**
- Authenticated API: 7200 requests/hour
- Unauthenticated API: Configurable per IP
- Jobs API: 600 calls/authenticated user

---

### FACT-039: Rate Limit Headers Missing

| Field | Value |
|-------|-------|
| **ID** | FACT-039 |
| **Category** | API |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- 429 received without rate limit info
- Issue creation limit hit unexpectedly
- Different limit types with different headers

**Root Cause:**
Some rate limits (like issue creation) don't include response headers, even when authenticated API shows remaining quota.

**Resolution Steps:**
1. Understand different rate limit types
2. Implement conservative request pacing
3. Track requests client-side
4. Use separate monitoring for different operations
5. Check application-level vs Rack::Attack limits

---

### FACT-040: 500 Instead of 429 for Rate Limit (GraphQL)

| Field | Value |
|-------|-------|
| **ID** | FACT-040 |
| **Category** | API |
| **Severity** | Low |
| **CVE** | N/A |

**Symptoms:**
- 500 Internal Server Error instead of 429
- GraphQL issue creation rate limited
- Inconsistent error codes

**Root Cause:**
Bug in GitLab where issue creation rate limit via GraphQL returns 500 instead of 429.

**Resolution Steps:**
1. Handle both 429 and 500 as potential rate limits
2. Report issue to GitLab if not yet fixed
3. Implement retry logic for 500 errors
4. Use REST API as workaround
5. Check GitLab version for bug fixes

---

### FACT-041: Cloudflare 429 Without GitLab Limits

| Field | Value |
|-------|-------|
| **ID** | FACT-041 |
| **Category** | API |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- 429 from Cloudflare while RateLimit headers show capacity
- GitLab.com specific issue
- CDN-level blocking

**Root Cause:**
Cloudflare's own rate limiting separate from GitLab's application-level limits.

**Resolution Steps:**
1. Reduce request frequency significantly
2. Distribute requests across time
3. Contact GitLab support for investigation
4. Check for IP reputation issues
5. Consider using authenticated requests only

---

## GitLab Pages Issues

### FACT-042: 404 After Successful Deployment

| Field | Value |
|-------|-------|
| **ID** | FACT-042 |
| **Category** | GitLab Pages |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Deployment job succeeds
- Artifacts browsable in GitLab
- Page returns 404 at expected URL

**Root Cause:**
Multiple possible causes: missing `index.html`, wrong artifact path, `pages:` job not configured correctly.

**Resolution Steps:**
1. Verify `.gitlab-ci.yml` contains `pages:` job
2. Ensure `pages:deploy` job is running in pipeline
3. Check artifacts are in `public/` directory
4. Confirm `index.html` exists at root
5. Wait for propagation (can take several minutes)

---

### FACT-043: namespace_in_path 404 Error

| Field | Value |
|-------|-------|
| **ID** | FACT-043 |
| **Category** | GitLab Pages |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- 404 with separate GitLab Pages server
- URL structure mismatch
- Works in some configurations

**Root Cause:**
`namespace_in_path` setting mismatch between GitLab server and GitLab Pages server.

**Resolution Steps:**
1. Ensure identical `namespace_in_path` values on both servers
2. Check `/etc/gitlab/gitlab.rb` on both servers
3. Reconfigure GitLab on both servers
4. Verify URL structure expectations match configuration
5. Review GitLab Pages documentation for setup

---

### FACT-044: OAuth Authentication Error

| Field | Value |
|-------|-------|
| **ID** | FACT-044 |
| **Category** | GitLab Pages |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `Client authentication failed due to unknown client`
- Access control enabled
- Previously working authentication fails

**Root Cause:**
Authentication token in `/etc/gitlab/gitlab-secrets.json` became invalid.

**Resolution Steps:**
1. Edit `/etc/gitlab/gitlab-secrets.json`
2. Remove the `gitlab_pages` section
3. Reconfigure GitLab
4. Re-register Pages OAuth application
5. Test authentication flow

---

### FACT-045: Disk Access Disabled Error

| Field | Value |
|-------|-------|
| **ID** | FACT-045 |
| **Category** | GitLab Pages |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `disk access is disabled via enable-disk=false`
- Pages configured for object storage but still referencing disk
- Migration incomplete

**Root Cause:**
GitLab Rails instructing Pages to serve from disk while Pages configured to disable disk access.

**Resolution Steps:**
1. Complete object storage migration
2. Update GitLab configuration for object storage
3. Reconfigure GitLab
4. Verify storage backend consistency
5. Check for mixed configuration states

---

## Kubernetes Integration Issues

### FACT-046: Context Deadline Exceeded

| Field | Value |
|-------|-------|
| **ID** | FACT-046 |
| **Category** | Kubernetes |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `context deadline exceeded`
- Build pod scheduling fails
- Kubernetes API timeouts

**Root Cause:**
Cluster cannot schedule build pod before `poll_timeout` or Kubernetes API experiencing high latency.

**Resolution Steps:**
1. Increase `poll_timeout` in `config.toml`
2. Check kube-apiserver metrics for latency
3. Monitor error rates for core resource operations
4. Scale cluster if resource constrained
5. Review network connectivity to API server

---

### FACT-047: Agent Certificate Authority Error

| Field | Value |
|-------|-------|
| **ID** | FACT-047 |
| **Category** | Kubernetes |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `x509: certificate signed by unknown authority`
- Agent cannot connect to GitLab
- Internal CA not trusted

**Root Cause:**
GitLab instance using certificate from internal CA unknown to the agent.

**Resolution Steps:**
1. Customize Helm installation with CA certificate
2. Add `--set-file config.kasCaCert=my-custom-ca.pem` to install
3. Ensure PEM or DER-encoded certificate format
4. Verify certificate chain completeness
5. Restart agent after certificate addition

---

### FACT-048: Missing KUBECONFIG Variable

| Field | Value |
|-------|-------|
| **ID** | FACT-048 |
| **Category** | Kubernetes |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Deployment variables not passed to job
- KUBECONFIG empty or missing
- Kubernetes commands fail

**Root Cause:**
Job missing `environment:name` setting required for Kubernetes credential injection.

**Resolution Steps:**
1. Add `environment:name` to deployment job
2. Ensure environment matches cluster configuration
3. Verify Deploy Token named `gitlab-deploy-token` exists
4. Check cluster integration settings
5. Review project-level cluster upgrade from GitLab 12.0

---

### FACT-049: API Version Deprecated Error

| Field | Value |
|-------|-------|
| **ID** | FACT-049 |
| **Category** | Kubernetes |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Error: `no matches for kind "Deployment" in version "extensions/v1beta1"`
- Kubernetes 1.16+ cluster
- Auto DevOps deployment fails

**Root Cause:**
Deprecated API versions removed in Kubernetes 1.16+.

**Resolution Steps:**
1. Update Helm charts to use current API versions
2. Use `apps/v1` instead of `extensions/v1beta1`
3. Update auto-deploy images to latest version
4. Check all manifest files for deprecated APIs
5. Use `kubectl convert` to update manifests

---

### FACT-050: Read-Only Root Filesystem Policy

| Field | Value |
|-------|-------|
| **ID** | FACT-050 |
| **Category** | Kubernetes |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Runner installation fails
- Build pods blocked by admission controller
- Gatekeeper or Kyverno policy violations

**Root Cause:**
Admission policies forcing read-only root filesystems blocking GitLab Runner pods.

**Resolution Steps:**
1. Set `securityContext.readOnlyRootFilesystem: true` in pod specs
2. Configure Runner to comply with cluster policies
3. Add appropriate volume mounts for writable paths
4. Review Gatekeeper/Kyverno policy exceptions
5. Work with cluster administrators on policy adjustments

---

## Infrastructure & Performance Issues

### FACT-051: Gitaly Bottleneck with Monorepos

| Field | Value |
|-------|-------|
| **ID** | FACT-051 |
| **Category** | Infrastructure |
| **Severity** | High |
| **CVE** | N/A |

**Symptoms:**
- Extreme slowness with large repositories
- git-pack-objects consuming high CPU and memory
- Clone/fetch operations timeout

**Root Cause:**
Gitaly service analyzing entire commit history for pack operations.

**Resolution Steps:**
1. Implement shallow clone with `GIT_DEPTH` variable
2. Configure dependency caching
3. Consider repository splitting into submodules
4. Increase Gitaly server resources
5. Use partial clone features where supported

---

### FACT-052: High Availability Complexity

| Field | Value |
|-------|-------|
| **ID** | FACT-052 |
| **Category** | Infrastructure |
| **Severity** | High |
| **CVE** | N/A |

**Symptoms:**
- Complex multi-datacenter deployments failing
- Synchronous latency requirements not met
- Support scope limitations

**Root Cause:**
HA recommended only for 3000+ users, multi-datacenter requires synchronous-capable latency.

**Resolution Steps:**
1. Evaluate if HA is necessary for user count
2. Test latency between datacenters
3. Consider single-datacenter with DR strategy
4. Review GitLab HA documentation thoroughly
5. Engage GitLab Professional Services for complex deployments

---

## Authentication & Authorization Issues

### FACT-053: Failed Authentication IP Ban

| Field | Value |
|-------|-------|
| **ID** | FACT-053 |
| **Category** | Authentication |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- HTTP 403 for 1 hour
- Affects Git and container registry requests
- Triggered by failed authentication attempts

**Root Cause:**
30 failed authentication requests in 3-minute period from single IP triggers automatic ban.

**Resolution Steps:**
1. Wait for 1-hour ban period to expire
2. Verify credentials are correct before retrying
3. Implement exponential backoff for automation
4. Check for credential rotation issues
5. Note: successful auth resets counter

---

### FACT-054: Job Token Scope Restrictions

| Field | Value |
|-------|-------|
| **ID** | FACT-054 |
| **Category** | Authentication |
| **Severity** | Medium |
| **CVE** | N/A |

**Symptoms:**
- Cross-project access denied in CI jobs
- Container registry pull fails
- API calls to other projects rejected

**Root Cause:**
CI/CD job tokens scoped to executing project only (GitLab 15.9+).

**Resolution Steps:**
1. Add target project to job token scope allowlist
2. Use project access tokens for cross-project access
3. Configure CI/CD settings in target project
4. Review job token documentation for scope details
5. Consider using deploy tokens for broader access

---

---

## Statistics

### Distribution by Category

| Category | Count |
|----------|-------|
| CI/CD Pipeline | 8 |
| Security/CVE | 11 |
| GitLab Runner | 6 |
| Container Registry | 7 |
| Merge Request | 5 |
| API & Rate Limiting | 4 |
| GitLab Pages | 4 |
| Kubernetes | 5 |
| Infrastructure | 2 |
| Authentication | 2 |
| **Total** | **54** |

### Distribution by Severity

| Severity | Count |
|----------|-------|
| Critical | 5 |
| High | 14 |
| Medium | 28 |
| Low | 7 |

### Critical CVEs Summary

| CVE | CVSS | Description |
|-----|------|-------------|
| CVE-2024-45409 | 10.0 | SAML Authentication Bypass |
| CVE-2024-6678 | 9.9 | Arbitrary User Pipeline Trigger |
| CVE-2024-5655 | 9.6 | Pipeline Execution as Arbitrary User |
| CVE-2024-6385 | 9.6 | Pipeline Execution Under Arbitrary Identity |
| CVE-2024-9164 | Critical | Pipeline on Arbitrary Branches |

---

## References

- **GitLab Documentation:** https://docs.gitlab.com
- **GitLab Security Advisories:** https://about.gitlab.com/releases/categories/releases/
- **GitLab Issue Tracker:** https://gitlab.com/gitlab-org/gitlab/-/issues
- **GitLab CI/CD Troubleshooting:** https://docs.gitlab.com/ci/troubleshooting/
- **GitLab Runner FAQ:** https://docs.gitlab.com/runner/faq/
- **Container Registry Troubleshooting:** https://docs.gitlab.com/administration/packages/container_registry_troubleshooting/
- **Kubernetes Agent Troubleshooting:** https://docs.gitlab.com/ee/user/clusters/agent/troubleshooting.html
- **GitLab Pages Troubleshooting:** https://docs.gitlab.com/administration/pages/troubleshooting/
- **Rate Limits Documentation:** https://docs.gitlab.com/security/rate_limits/
- **HackerOne Bug Bounty Program:** https://hackerone.com/gitlab

---

*This document is generated for reference purposes and should be used alongside official GitLab documentation. Always verify information against the latest GitLab release notes and security advisories.*
