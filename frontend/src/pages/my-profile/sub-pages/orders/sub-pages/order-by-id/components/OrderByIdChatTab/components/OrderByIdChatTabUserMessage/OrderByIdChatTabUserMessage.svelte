<script lang="ts">
  import type { ChatMessage } from '@/types/ChatMessage'
  import OrderByIdChatTabMessage from '../OrderByIdChatTabMessage/OrderByIdChatTabMessage.svelte'
  import OrderByIdChatTabUserMessageLayout from './OrderByIdChatTabUserMessageLayout.svelte'

  type OrderByIdChatTabUserMessageProps = {
    message: ChatMessage
  }

  const { message }: OrderByIdChatTabUserMessageProps = $props()
</script>

<OrderByIdChatTabUserMessageLayout>
  <OrderByIdChatTabMessage {message} />
</OrderByIdChatTabUserMessageLayout>
