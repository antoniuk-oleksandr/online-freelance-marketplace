<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { Link } from 'svelte-routing'
  import { getHoverClass } from '@/utils/utils'

  type BadgeLayoutProps = {
    id?: number
  } & LayoutProps

  let { id, children }: BadgeLayoutProps = $props()
  const hoverClass = getHoverClass()
  const styles = `${id === undefined ? '' : hoverClass} group relative rounded-xl grid place-items-center text-center border border-light-palette-divider dark:border-dark-palette-divider px-3 py-1 bg-light-palette-background-block dark:bg-dark-palette-background-block`
</script>

{#if id === undefined}
  <div class={styles}>
    {@render children()}
  </div>
{:else}
  <Link to="/search?skill={id}" class={styles}>
    {@render children()}
  </Link>
{/if}
