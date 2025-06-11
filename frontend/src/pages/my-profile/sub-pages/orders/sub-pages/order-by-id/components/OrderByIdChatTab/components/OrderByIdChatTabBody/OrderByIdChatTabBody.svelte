<script lang="ts">
  import type { ChatData } from '@/types/ChatData'
  import OrderByIdChatTabBodyLayout from './OrderByIdChatTabBodyLayout.svelte'
  import OrderByIdChatTabMessageList from '../OrderByIdChatTabMessageList/OrderByIdChatTabMessageList.svelte'
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'

  type OrderByIdChatTabBodyProps = {
    chatData: ChatData
  }

  let { chatData }: OrderByIdChatTabBodyProps = $props()
</script>

<OrderByIdChatTabBodyLayout messagesLength={chatData.messages.length}>
  {#if chatData.messages.length === 0}
    <NoContentMessage text="You haven't started the conversation yet." />
  {:else}
    <OrderByIdChatTabMessageList {chatData} />
  {/if}
</OrderByIdChatTabBodyLayout>
