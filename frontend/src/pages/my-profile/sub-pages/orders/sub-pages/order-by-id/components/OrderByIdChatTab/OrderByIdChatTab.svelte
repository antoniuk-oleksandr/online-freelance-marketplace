<script module lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import OrderByIdChatTabHeader from './components/OrderByIdChatTabHeader/OrderByIdChatTabHeader.svelte'
  import OrderByIdChatTabFooter from './components/OrderByIdChatTabFooter/OrderByIdChatTabFooter.svelte'
  import OrderByIdChatTabBody from './components/OrderByIdChatTabBody/OrderByIdChatTabBody.svelte'
  import type { MyProfileChatRequestResponse } from '@/types/MyProfileChatRequestResponse'
  import { messageStore } from './stores/message-store'
  import type { ChatMessage } from '@/types/ChatMessage'
  import { processChatMessages } from './helpers'

  let messages = $state<ChatMessage[]>([])
  messageStore.subscribe((newMessages) => (messages = newMessages))

  export { orderByIdChatTabComponent }
</script>

{#snippet orderByIdChatTabComponent(response: MyProfileChatRequestResponse, orderId: number)}
  <PaperElement
    styles="-mx-6 lg:-mx-0 !ring-0 lg:dark:!ring-1 !p-0 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <OrderByIdChatTabHeader {...response.data.chatPartner} />
      <OrderByIdChatTabBody
        chatData={{
          ...response.data,
          messages: processChatMessages(orderId, response.data.messages, messages),
        }}
      />
      <OrderByIdChatTabFooter
        {orderId}
        chatParentId={response.data.chatPartner.partnerId}
      />
    {/if}
  </PaperElement>
{/snippet}
