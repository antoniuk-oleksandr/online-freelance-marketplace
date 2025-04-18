<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type RichTextEditorLayoutProps = LayoutProps & {
    error?: string | null
  }

  const { children, error }: RichTextEditorLayoutProps = $props()
</script>

<div
  class="{error ? 'ring-red-500': 'ring-light-palette-divider dark:ring-dark-palette-divider'} flex flex-col ring-1 rounded-xl"
>
  {@render children()}
</div>
