<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { onMount } from 'svelte'

  const { children }: LayoutProps = $props()

  let element = $state<undefined | HTMLDivElement>()

  onMount(() => {
    if (element) element.scrollTop = element.scrollHeight
  })
</script>

<div
  bind:this={element}
  class="grow px-6 flex flex-col gap-3 max-h-chat-body overflow-y-auto py-1"
>
  {@render children()}
</div>
