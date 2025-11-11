FlagForge — Multi-tenant Feature Flag Platform

FlagForge is a backend and cloud-native platform for safely turning application features on or off for specific users, segments, or environments in real time, without redeploying code. It separates the “control plane” where teams manage flags and targeting rules from the “data plane” that evaluates those rules at low latency inside the request path.
What it does

    Lets teams define feature flags, segments, and rollout rules (attributes, segments, percentages) across dev/stage/prod.

    Delivers instant flag changes to applications via streaming updates, enabling progressive delivery, canaries, A/B tests, and kill switches.

    Provides multi-tenant isolation, API keys, audit trails, and observability so large orgs can run it safely at scale.

Why it’s useful

    Reduce risk by rolling out features gradually and rolling back instantly if error rates rise.

    Decouple deploy from release: ship code dark and enable features later for specific cohorts.

    Run experiments and targeted betas without branching code or re-deploying.

High-level architecture

    Control Plane

        REST/gRPC APIs to manage organizations, environments, flags, segments, targeting rules, API keys, and RBAC.

        Stores metadata and audits in a relational database.

        Emits signed manifests and change events for the data plane and SDKs.

    Data Plane

        Low-latency rule evaluation service with a compact DSL (attribute checks, segment membership, percentage rollouts, prerequisites).

        Sticky bucketing via consistent hashing to maintain user experience across sessions.

        Caches hot flags/segments and exposes a fast evaluation endpoint for SDKs.

    Realtime Delivery

        Server-Sent Events/WebSockets push “flag delta” updates to connected SDKs; long-poll as fallback.

        Clients swap in new rules atomically without restarts.

    Storage

        PostgreSQL for control-plane metadata and audit logs.

        Redis for hot caches and pub/sub invalidation fan-out.

        Object storage (e.g., S3) for versioned snapshots/exports for disaster recovery.

How it works (request flow)

    Developer defines a flag and targeting rules in the control plane (e.g., enable Feature X for 10% of users in prod, 100% in staging).

    The control plane persists metadata and emits a signed flag manifest and change event.

    SDKs connected to the data plane receive deltas over SSE/WebSockets and refresh their in-memory rule set.

    At runtime, an application calls the SDK to evaluate a flag for a user/context.

    The SDK evaluates locally against cached rules or queries the data plane if needed; the decision returns in single-digit milliseconds.

Key capabilities

    Targeting: attribute rules, segment membership, percentage/gradual rollouts, prerequisites, and overrides.

    Multitenancy: org/env isolation, per-tenant rate limits/quotas, hierarchical API keys, soft/hard deletes and GC.

    Safety: kill switch, instant rollback, canary/blue-green rollout patterns.

    Observability: Prometheus metrics (eval latency, cache hit ratio, stream fan-out), structured logs, tracing/profiling endpoints.

    Resilience: circuit breakers and timeouts, idempotent upserts with outbox pattern, chaos tests for cache flush and network partitions.

Performance targets (illustrative)

    Data-plane evaluation: ~4–8 ms p95 with warm cache; sustained tens of thousands of eval/s per pod.

    Streaming fan-out: thousands of SDK clients per node with near-instant propagation (<1s) on flag updates.

SDK model

    Supported modes: streaming (preferred), long-poll, or offline bootstrap from a signed manifest.

    Caching: in-memory LRU with TTLs and checksum validation; consistent-sticky hashing to keep user experience stable.

    Minimal footprint per request; thread-safe and non-blocking evaluation path.

Security model

    Signed manifests and per-environment SDK keys.

    Per-tenant RBAC (admin, maintainer, reader) and audit logs for change tracking.

    Least-privilege access to backing stores; encryption in transit and at rest.

Operability

    Kubernetes-friendly: health/readiness probes, horizontal autoscaling, and graceful shutdown with in-flight drain.

    Continuous delivery: canary and blue/green scripts; config for progressive rollouts and automated rollback on error-budget burn.

    Backups: daily control-plane DB backups and periodic object storage snapshots of manifests.

Roadmap ideas

    Rule editor UI with simulation/playground.

    Native experimentation metrics (exposure events, conversions) and guardrail alerts.

    Edge SDK mode (co-located evaluation) for ultra-low latency.

Quick glossary

    Flag: a named boolean/variant that controls behavior in code.

    Segment: a saved set of users matching rules (attributes, lists).

    Rollout: gradual percentage-based enablement for a population.

    Sticky bucketing: hashing identity to a stable bucket to keep consistent experiences.

