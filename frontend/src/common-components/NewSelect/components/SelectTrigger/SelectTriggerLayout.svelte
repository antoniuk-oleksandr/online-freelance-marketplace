<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type SelectTriggerLayoutProps = LayoutProps & {
    selectWidth?: string
    triggerStyles?: string
  }

  const { children, selectWidth, triggerStyles }: SelectTriggerLayoutProps = $props()
</script>

<div
  class="{triggerStyles} {selectWidth} justify-between rounded-lg border bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider flex font-semibold gap-x-3 items-center  px-4 py-2 hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover duration-200 ease-out"
>
  {@render children()}
</div>
