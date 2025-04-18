<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type TextareaContainerLayoutProps = LayoutProps & {
    focus: boolean
    wasSubmitted: boolean
    error?: string
  }

  const { children, focus, error, wasSubmitted }: TextareaContainerLayoutProps = $props()
</script>

<div
  class="{error && wasSubmitted
    ? 'ring-red-500'
    : focus
      ? 'ring-cyan-500 dark:ring-cyan-500'
      : 'ring-light-palette-divider dark:ring-dark-palette-divider'} text-light-palette-text-primary dark:text-dark-palette-text-primary duration-200 ease-out ring-1 rounded-xl"
>
  {@render children()}
</div>
