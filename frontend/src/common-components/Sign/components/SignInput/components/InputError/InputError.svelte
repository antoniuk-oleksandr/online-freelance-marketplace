<script lang="ts">
  type InputErrorProps = {
    wasSubmitted: boolean
    error: string | null | undefined
    styles?: string
  }

  const { wasSubmitted, error, styles = '' }: InputErrorProps = $props()
</script>

{#if wasSubmitted && error}
  <span class="{styles} text-red-500 text-sm">{error}</span>
{/if}
