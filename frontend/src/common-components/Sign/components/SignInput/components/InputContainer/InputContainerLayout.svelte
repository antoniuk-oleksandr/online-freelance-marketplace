<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'

  type InputContainerLayoutProps = LayoutProps & {
    isFocused: boolean
    wasSubmitted: boolean
    error: string | null
    value: string
  }

  const { isFocused, wasSubmitted, error, children, value }: InputContainerLayoutProps = $props()
</script>

<div
  class="relative rounded-md ring-1 {wasSubmitted && error
    ? 'ring-red-500'
    : isFocused || value.length > 0
      ? 'ring-cyan-500'
      : 'ring-light-palette-divider dark:ring-dark-palette-divider'}"
>
  {@render children()}
</div>
