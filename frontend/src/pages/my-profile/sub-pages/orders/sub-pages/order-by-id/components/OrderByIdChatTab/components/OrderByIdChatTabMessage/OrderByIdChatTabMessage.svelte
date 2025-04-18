<script lang="ts">
  import type { ChatMessage } from '@/types/ChatMessage'
  import OrderByIdChatTabMessageLayout from './OrderByIdChatTabMessageLayout.svelte'
  import { formatChatMessageTime } from '../../helpers'
  import Icon from '@iconify/svelte'
  import { ChatMessageType } from '@/types/ChatMessageType'

  type OrderByIdChatTabMessageProps = {
    message: ChatMessage
    wasSentByPartner?: boolean
  }

  const { message, wasSentByPartner = false }: OrderByIdChatTabMessageProps = $props()
</script>

<OrderByIdChatTabMessageLayout {wasSentByPartner}>
  <span
    class="whitespace-pre-wrap break-words w-fit text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >{message.content}</span
  >
  <div class="flex ml-3 items-center">
    <span class="text-xs break-keep w-full whitespace-nowrap">
      {formatChatMessageTime(message.sentAt)}
    </span>
    {#if !wasSentByPartner}
      <Icon
        class="ml-1 text-light-palette-text-primary dark:text-dark-palette-text-primary"
        icon="hugeicons:tick-{message.type === ChatMessageType.Read ? 'double-' : ''}01"
        width="18"
        height="18"
      />
    {/if}
  </div>
</OrderByIdChatTabMessageLayout>
