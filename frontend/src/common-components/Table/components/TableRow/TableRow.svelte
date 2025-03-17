<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { Link } from 'svelte-routing'

  type TableRowProps = LayoutProps & {
    gridTemplateColumns: string
    styles?: string
    link?: string
    hiddenOnSmall?: boolean
  }

  const {
    children,
    styles,
    link,
    gridTemplateColumns,
    hiddenOnSmall = false,
  }: TableRowProps = $props()
</script>

{#snippet renderElement(link?: string)}
  <div
    class="{link ? '' : styles} {gridTemplateColumns} {hiddenOnSmall
      ? 'hidden lg:grid'
      : ''} grid items-center"
  >
    {@render children()}
  </div>
{/snippet}

{#if link && !hiddenOnSmall}
  <Link hidden={hiddenOnSmall} class="duration-200 ease-out hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover {styles} {gridTemplateColumns}" to={link}>
    {@render renderElement(link)}
  </Link>
{:else}
  {@render renderElement(link)}
{/if}

