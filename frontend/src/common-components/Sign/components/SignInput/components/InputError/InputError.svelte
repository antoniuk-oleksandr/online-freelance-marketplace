<script lang="ts">
  type InputErrorProps = {
    wasSubmitted: boolean
    error: string | null | undefined
  }

  const { wasSubmitted, error }: InputErrorProps = $props()
</script>

{#if wasSubmitted && error}
  <span class="text-red-500 text-sm">{error}</span>
{/if}
