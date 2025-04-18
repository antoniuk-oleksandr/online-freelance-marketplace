<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type SelectTriggerLayoutProps = LayoutProps & {
    selectWidth?: string
    triggerStyles?: string
    error?: string | null
  }

  const { children, selectWidth, triggerStyles, error }: SelectTriggerLayoutProps = $props()
</script>

<div
  class="{error ? 'border-red-500': 'border-light-palette-divider dark:border-dark-palette-divider'} {triggerStyles} {selectWidth} text-light-palette-text-primary dark:text-dark-palette-text-primary justify-between rounded-lg border bg-light-palette-background-block dark:bg-dark-palette-background-block flex font-semibold gap-x-3 items-center  px-4 py-2 hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover duration-200 ease-out"
>
  {@render children()}
</div>
