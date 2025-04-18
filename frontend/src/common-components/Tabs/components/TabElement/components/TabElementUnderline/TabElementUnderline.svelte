<script lang="ts">
  import { slide } from 'svelte/transition'

  type TabElementUnderlineProps = {
    index: number
    activeTabIndex: number
  }

  const { index, activeTabIndex }: TabElementUnderlineProps = $props()
</script>

{#if index === activeTabIndex}
  <div
    in:slide={{ duration: 300 }}
    out:slide={{ duration: 300 }}
    class="h-1 rounded-t-md absolute left-0 -bottom-px w-full bg-cyan-500"
  ></div>
{/if}
