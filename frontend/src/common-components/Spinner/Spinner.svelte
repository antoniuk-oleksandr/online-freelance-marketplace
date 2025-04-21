<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { fade } from 'svelte/transition'

  type SpinnerProps = {
    size?: string
    color?: string
    useDelay?: boolean
  }

  let render = $state(false)
  let timeoutRes = $state<number | undefined>()

  onMount(() => {
    timeoutRes = setTimeout(() => (render = true), 2000)
  })

  onDestroy(() => {
    if (timeoutRes === undefined) return

    clearTimeout(timeoutRes)
    timeoutRes = undefined
  })

  const { size = 'size-14', color = 'border-l-cyan-500', useDelay = true }: SpinnerProps = $props()
</script>

{#if render || !useDelay}
  <div
    transition:fade={{ duration: 300 }}
    class="animate-spin border-2 border-transparent rounded-full {size} {color}"
  ></div>
{/if}
