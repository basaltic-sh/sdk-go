# Changelog

## 0.35.2

- Include `scale_in_blocked` in the documented autoscaling status values.

## 0.35.1

- Pool updates support `Idempotency-Key` for safe retries.
- Choose the exact members to remove when reducing an instance pool with `scale_in_instances`.
- Protect or unprotect pool members with the pool member protection endpoint, and read protection and blocked scale-in status in pool responses.
- Include Email quota policies in quota responses.

Compatibility: `quota.Client.ListQuotas` now returns `*quota.QuotaListResponse`. Read resource quotas from `response.Quotas` instead of `response.Items`; Email limits are in `response.Policies`.
