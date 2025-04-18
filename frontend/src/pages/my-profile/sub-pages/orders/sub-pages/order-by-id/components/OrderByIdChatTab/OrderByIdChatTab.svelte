<script module lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import OrderByIdChatTabHeader from './components/OrderByIdChatTabHeader/OrderByIdChatTabHeader.svelte'
  import OrderByIdChatTabFooter from './components/OrderByIdChatTabFooter/OrderByIdChatTabFooter.svelte'
  import OrderByIdChatTabBody from './components/OrderByIdChatTabBody/OrderByIdChatTabBody.svelte'
  import type { MyProfileChatRequestResponse } from '@/types/MyProfileChatRequestResponse'

  export { orderByIdChatTabComponent }
</script>

{#snippet orderByIdChatTabComponent(response: MyProfileChatRequestResponse)}
  <PaperElement
    styles="!ring-0 lg:dark:!ring-1 !p-0 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  >
    {#if response && response.status === 200}
      <OrderByIdChatTabHeader {...response.data.chatPartner} />
      <OrderByIdChatTabBody chatData={response.data} />
      <OrderByIdChatTabFooter />
    {/if}
  </PaperElement>
{/snippet}
