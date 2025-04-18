<script lang="ts">
  import { Link } from 'svelte-routing'

  type OrderRequestSummaryItemProps = {
    title: string
    value: string | number
    styles?: string
    valueStyles?: string
    titleStyles?: string
    valueLink?: string
  }

  const {
    title,
    value,
    styles,
    valueStyles,
    titleStyles,
    valueLink,
  }: OrderRequestSummaryItemProps = $props()
</script>

<div
  class="{styles} flex capitalize items-center justify-between text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
>
  <span class={titleStyles}>{title}</span>
  {#if valueLink}
    <Link to={valueLink} class={valueStyles}>
      <span class={valueStyles}>{value}</span>
    </Link>
  {:else}
    <span class={valueStyles}>{value}</span>
  {/if}
</div>
