<script lang="ts">
  import RichTextEditorFooterLayout from './RichTextEditorFooterLayout.svelte'

  type RichTextEditorFooterProps = {
    limit: number
    count?: number
    error?: string | null
  }

  const { count = 0, limit, error }: RichTextEditorFooterProps = $props()
</script>

<RichTextEditorFooterLayout {error}>
  <span>{count} / {limit}</span>
</RichTextEditorFooterLayout>
