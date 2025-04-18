<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import type { ChatMessage } from '@/types/ChatMessage'
  import type { ChatPartner } from '@/types/ChatPartner'
  import OrderByIdChatTabPartnerMessageLayout from './OrderByIdChatTabPartnerMessageLayout.svelte'
  import OrderByIdChatTabMessage from '../OrderByIdChatTabMessage/OrderByIdChatTabMessage.svelte'

  type OrderByIdChatTabPartnerMessageProps = {
    chatPartner: ChatPartner
    message: ChatMessage
    renderAvatar: boolean
  }

  const { chatPartner, message, renderAvatar }: OrderByIdChatTabPartnerMessageProps = $props()
</script>

<OrderByIdChatTabPartnerMessageLayout {renderAvatar}>
  {#if renderAvatar}
    <Avatar size="extra small" src={chatPartner.avatar} alt="avatar" />
  {/if}
  <OrderByIdChatTabMessage wasSentByPartner {message} />
</OrderByIdChatTabPartnerMessageLayout>
