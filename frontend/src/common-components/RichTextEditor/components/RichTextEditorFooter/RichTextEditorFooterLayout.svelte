<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type RichTextEditorFooterLayoutProps = LayoutProps & {
    error?: string | null
  }

  const { children, error }: RichTextEditorFooterLayoutProps = $props()
</script>

<div
  class="{error ? 'border-red-500' : 'border-light-palette-divider dark:border-dark-palette-divider'} p-6 text-base font-medium border-t"
>
  {@render children()}
</div>
