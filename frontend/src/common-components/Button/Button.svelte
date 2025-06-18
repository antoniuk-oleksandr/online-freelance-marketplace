<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import { Link } from 'svelte-routing'

  type MyButtonProps = LayoutProps & {
    clickAction?: () => void
    styles?: string
    linkStyles?: string
    type?: 'button' | 'submit' | 'reset'
    loading?: boolean
    color?: 'red' | 'default' | 'grey' | 'outline' | 'white' | 'green' | 'darkGray' | 'whiteWithGrayOutline'
    link?: string
  }

  const { children, linkStyles, clickAction, styles, type, loading, color, link }: MyButtonProps =
    $props()

  const colorStyles = {
    red: { buttonStyle: '!bg-red-500 text-white', hoverColor: 'hover:!bg-red-400' },
    white: { buttonStyle: '!bg-white text-cyan-500', hoverColor: 'hover:!bg-gray-200' },
    default: { buttonStyle: '!bg-cyan-500 text-white', hoverColor: 'hover:!bg-cyan-400' },
    green: { buttonStyle: '!bg-green-600 text-white', hoverColor: 'hover:!bg-green-500' },
    grey: {
      buttonStyle: '!bg-gray-300 dark:!bg-zinc-700',
      hoverColor: 'hover:!bg-gray-200 hover:dark:!bg-zinc-600',
    },
    outline: {
      buttonStyle: 'ring-1 ring-light-palette-divider dark:ring-dark-palette-divider',
      hoverColor: 'hover:!bg-light-palette-action-hover hover:dark:!bg-dark-palette-action-hover',
    },
    darkGray: {
      buttonStyle: 'bg-gray-800 text-white dark:bg-gray-700 ',
      hoverColor: 'hover:bg-gray-900 dark:hover:bg-gray-600',
    },
    whiteWithGrayOutline: {
      buttonStyle: 'bg-light-palette-background-block text-gray-800 ring-1 ring-light-palette-divider dark:ring-dark-palette-divider dark:bg-dark-palette-background-block dark:text-gray-300',
      hoverColor: 'hover:bg-gray-100 dark:hover:bg-gray-700'
    }
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
  <Link class="{linkStyles} w-full" to={link}>
    {@render button()}
  </Link>
{:else}
  {@render button()}
{/if}
