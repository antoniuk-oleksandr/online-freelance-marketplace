<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type PaginationButtonProps = LayoutProps & {
    currentPage: boolean
    disabled: boolean
    clickAction?: () => void
  }

  const { children, disabled, currentPage, clickAction }: PaginationButtonProps = $props()
</script>

<button
  onclick={clickAction}
  disabled={disabled || currentPage}
  style="transition: all 0.2s ease-out !important;"
  class=" size-10 ring-1 rounded-md flex items-center justify-center font-semibold
  {disabled
    ? 'ring-light-palette-action-disabled dark:ring-dark-palette-action-disabled text-light-palette-action-disabled dark:text-dark-palette-action-disabled' // If disabled
    : currentPage
      ? 'ring-cyan-500 bg-cyan-500 !text-white'
      : 'active:scale-95 hover:!bg-light-palette-action-hover hover:dark:!bg-dark-palette-action-hover ring-light-palette-divider dark:ring-dark-palette-divider'}"
>
  {@render children()}
</button>
