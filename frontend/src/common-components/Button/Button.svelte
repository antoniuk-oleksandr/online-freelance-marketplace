<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import { Link } from 'svelte-routing'

  type MyButtonProps = LayoutProps & {
    clickAction?: () => void
    styles?: string
    type?: 'button' | 'submit' | 'reset'
    loading?: boolean
    color?: 'red' | 'default' | 'grey' | 'outline' | 'white'
    link?: string
  }

  const { children, clickAction, styles, type, loading, color, link }: MyButtonProps = $props()

  const colorStyles = {
    red: { buttonStyle: '!bg-red-500 text-white', hoverColor: 'hover:!bg-red-400' },
    white: { buttonStyle: '!bg-white text-cyan-500', hoverColor: 'hover:!bg-gray-200' },
    default: { buttonStyle: '!bg-cyan-500 text-white', hoverColor: 'hover:!bg-cyan-400' },
    grey: {
      buttonStyle: '!bg-gray-300 dark:!bg-zinc-700',
      hoverColor: 'hover:!bg-gray-200 hover:dark:!bg-zinc-600',
    },
    outline: {
      buttonStyle: 'ring-1 ring-light-palette-divider dark:ring-dark-palette-divider',
      hoverColor: 'hover:!bg-light-palette-action-hover hover:dark:!bg-dark-palette-action-hover',
    },
  }

  let selectedStyle = color ? colorStyles[color] : colorStyles.default
</script>

{#snippet button()}
  <button
    style="transition: all 0.2s ease-out !important;"
    disabled={loading}
    type={type ?? 'button'}
    class="{styles} {loading
      ? 'opacity-70'
      : selectedStyle.hoverColor +
        ' active:!scale-95'} {selectedStyle.buttonStyle} px-4 rounded-md !h-12 !flex !items-center !justify-center font-semibold"
    onclick={clickAction}
  >
    {#if loading}
      <Spinner useDelay={false} size="size-8" color="border-l-white" />
    {:else}
      {@render children()}
    {/if}
  </button>
{/snippet}

{#if link}
  <Link class="w-full" to={link}>
    {@render button()}
  </Link>
{:else}
  {@render button()}
{/if}
