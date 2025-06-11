<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { onMount, tick } from 'svelte'
  import { chatBodyStore } from '../../stores/chat-body-store'

  type OrderByIdChatTabBodyLayoutProps = LayoutProps & {
    messagesLength: number
  }

  let { children, messagesLength }: OrderByIdChatTabBodyLayoutProps = $props()

  onMount(() => {
    tick().then(() => {
      if (!$chatBodyStore) return
      $chatBodyStore.scrollTop = $chatBodyStore.scrollHeight
    })
  })
</script>

<div
  bind:this={$chatBodyStore}
  class="{messagesLength === 0 ? 'items-center justify-center' : ''} overflow-x-hidden grow px-6 flex flex-col gap-3 max-h-chat-body overflow-y-auto py-1"
>
  {@render children()}
</div>
