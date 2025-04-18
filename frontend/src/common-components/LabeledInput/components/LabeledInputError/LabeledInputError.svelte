<script lang="ts">
  type LabeledInputErrorProps = {
    error?: string | null
    styles?: string
  }

  let { error }: LabeledInputErrorProps = $props()
</script>

{#if error}
  <p class="text-red-500">{error}</p>
{/if}
