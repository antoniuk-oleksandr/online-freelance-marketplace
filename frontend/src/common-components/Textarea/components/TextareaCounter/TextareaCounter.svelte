<script lang="ts">
  import { getValue } from 'felte'

  type TextareaContainer = {
    focus: boolean
    wasSubmitted: boolean
    error?: string
    useCounter: boolean
    value?: string
  }

  const { focus, wasSubmitted, error, value, useCounter }: TextareaContainer = $props()
</script>

{#if useCounter}
  <div
    class="{error && wasSubmitted
      ? 'border-red-500'
      : focus
        ? 'border-cyan-500 dark:border-cyan-500'
        : 'border-light-palette-divider dark:border-dark-palette-divider'} bg-light-palette-background-block dark:bg-dark-palette-background-block  rounded-b-xl p-6 border-t"
  >
    {value ? value.length : 0} / 2000
  </div>
{/if}
