<script lang="ts">
  import type { ChatData } from '@/types/ChatData'
  import { getUserId } from '@/utils/utils'
  import OrderByIdChatTabPartnerMessage from '../OrderByIdChatTabPartnerMessage/OrderByIdChatTabPartnerMessage.svelte'
  import OrderByIdChatTabUserMessage from '../OrderByIdChatTabUserMessage/OrderByIdChatTabUserMessage.svelte'
  import OrderByIdChatTabMessagesDivider from '../OrderByIdChatTabMessagesDivider/OrderByIdChatTabMessagesDivider.svelte'
  import { checkIfShouldRenderPartnerAvatar } from '../../helpers'

  type OrderByIdChatTabMessageListProps = {
    chatData: ChatData
  }

  const { chatData }: OrderByIdChatTabMessageListProps = $props()


  const userId = getUserId()
</script>

{#each chatData.messages as message, index}
  <OrderByIdChatTabMessagesDivider {index} messages={chatData.messages} />
  {#if message.senderId === userId}
    <OrderByIdChatTabUserMessage {message} />
  {:else}
    <OrderByIdChatTabPartnerMessage
      renderAvatar={checkIfShouldRenderPartnerAvatar(index, chatData.messages)}
      chatPartner={chatData.chatPartner}
      {message}
    />
  {/if}
{/each}
