
![OpenMetrics Logo](website/static/images/logo/logo-with-text.png)

# OpenMetrics

OpenMetrics a specification built upon and carefully extending [Prometheus exposition format](https://prometheus.io/docs/instrumenting/exposition_formats/) in almost 100% backwards-compatible ways.

> NOTE: This project recently moved to Prometheus and we are working on OpenMetrics 2.0! See the details in [#276](https://github.com/prometheus/OpenMetrics/issues/276) on how to participate!

# Spec

See [our spec file](specification/OpenMetrics.md) and [our proto](proto/openmetrics_data_model.proto).

# News and updates

Join the [mailing list](https://groups.google.com/forum/m/#!forum/openmetrics) or follow us on [Twitter](https://twitter.com/OpenMetricsIO)

# Code of Conduct
To make OpenMetrics a welcoming and harassment-free experience for everyone, we follow the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

# OpenMetrics 2.0 development

OpenMetrics 2.0 is currently under development. You can see [our charter](https://docs.google.com/document/d/1FCD-38Xz1-9b3ExgHOeDTQUKUatzgj5KbCND9t-abZY/edit?tab=t.6h53vr4k28b4#heading=h.uaaplxxbz60u) to understand our direction and our [meeting notes](https://docs.google.com/document/d/1FCD-38Xz1-9b3ExgHOeDTQUKUatzgj5KbCND9t-abZY/edit?tab=t.lvx6fags1fga#heading=h.uaaplxxbz60u) to understand our latest discussions.

## Contributing to OpenMetrics 2.0

### How can one propose a spec change to OpenMetrics?

The process starts with creating a new issue in the OpenMetrics’ GitHub repository: https://github.com/prometheus/OpenMetrics.

When opening an Issue, the author should try to describe the problem in detail. An issue asking for a change without explaining why such a thing is necessary will be ignored or closed. It's a good practice to focus on a feature that would be enabled by the change instead of going straight to implementation details.

Example of a bad issue:

```
Title: Relax requirement to add unit as suffixes
Body: It's annoying.
```

Example of a better issue:

```
Title: Allow exposing multiple metrics with same name in the same target
Body: 


# Problem Statement

Prometheus' federation endpoint cannot follow OpenMetrics specification because it's possible that multiple targets expose the same metric with same name but different metadata like Type/Help/Unit, but they all become a single target once exposed by Prometheus /federate endpoint. 

The same problem occurs with OpenTelemetry's Collector, who is able to collect metrics from several places and expose them in a single endpoint.

# Proposed idea

OpenMetrics should relax the requirement of exposing only one metric family by metric name. Instead, it should be allowed as long as metric TYPE or UNIT are different.
```

Once the issue is created, one of the maintainers should act by adding the necessary labels or closing the Issue if the idea is rejected.

### Labeling issues

#### `triage:deciding:*`

These labels are applied to issues when it is unclear yet if they are something the project will take on.

* `triage:deciding:community-feedback` - This issue is open to community discussion. If the community can provide sufficient reasoning, the project may accept it.
* `triage:deciding:needs-info` - This issue does not provide enough information for the project to accept it. It is left open to give the author time to add more details.

#### `triage:accepted:*`

These labels are applied to issues that describe a problem that is in scope and that we would like to tackle. Just because an issue is accepted does not mean that a solution suggested by the issue will be the solution applied.

* `triage:accepted:needs-champion` - This issue was discussed enough and the outcome is clear to the maintainers, however, someone to take this to the finish line is still needed.
* `triage:accepted:needs-formal-proposal` - This issue was discussed enough, the outcome is clear, and someone has already been assigned to implement the solution. The next step is to open a PR to the repository prometheus/proposals explaining the change to the wider Prometheus community. This proposal will cover the necessary changes that SDKs and/or Prometheus server will need to make.
* `triage:accepted:PoC-needed` - This issue was discussed amongst maintainers, but it's still unclear if the implementation is doable efficiently. A Proof of Concept showcasing the results is necessary to advance.
* `triage:accepted:ready` - The formal proposal has been accepted by the wider Prometheus community and is ready to be implemented.


A closed issue means it was rejected.

### Formal proposals

Compared to the whole Prometheus community, OpenMetrics maintainers and contributors are a relatively small group. However, whatever is decided in OpenMetrics impacts several SDKs, Prometheus Server itself, and several other projects in the ecosystem, e.g. Thanos, Cortex, and OpenTelemetry.

We follow Prometheus's proposal process to ensure we don't make changes that could harm the ecosystem.

When writing a formal proposal, the author needs to cover all the changes the ecosystem will need to make. If a proposal touches several parts of the ecosystem, such as parsers, storage, and SDKs, being as detailed as possible in all aspects will accelerate the approval needed to start the implementation.

Don't hesitate to create PoCs to better illustrate the final outcome.

### Accepted Proposals

To finalize the process, a PR is necessary to update the spec in: https://github.com/prometheus/OpenMetrics/blob/main/specification/OpenMetrics.md
