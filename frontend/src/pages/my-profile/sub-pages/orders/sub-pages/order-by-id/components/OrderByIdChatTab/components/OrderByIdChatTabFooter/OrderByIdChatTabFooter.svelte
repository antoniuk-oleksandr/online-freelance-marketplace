<script lang="ts">
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import OrderByIdChatTabFooterLayout from './OrderByIdChatTabFooterLayout.svelte'
  import OrderByIdChatTabTextArea from '../OrderByIdChatTabTextArea/OrderByIdChatTabTextArea.svelte'
  import OrderByIdChatTabSendButton from '../OrderByIdChatTabSendButton/OrderByIdChatTabSendButton.svelte'
    import OrderByIdChatTabAttachButton from '../OrderByIdChatTabAttachButton/OrderByIdChatTabAttachButton.svelte'
</script>

<OrderByIdChatTabFooterLayout>
  <DividerElement />
  <div class="flex gap-3 items-center py-3">
    <OrderByIdChatTabAttachButton/>
    <OrderByIdChatTabTextArea />
    <OrderByIdChatTabSendButton />
  </div>
</OrderByIdChatTabFooterLayout>
