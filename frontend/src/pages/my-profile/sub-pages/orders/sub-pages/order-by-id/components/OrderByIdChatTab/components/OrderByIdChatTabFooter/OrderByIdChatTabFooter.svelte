<script lang="ts">
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import OrderByIdChatTabFooterLayout from './OrderByIdChatTabFooterLayout.svelte'
  import OrderByIdChatTabTextArea from '../OrderByIdChatTabTextArea/OrderByIdChatTabTextArea.svelte'
  import OrderByIdChatTabSendButton from '../OrderByIdChatTabSendButton/OrderByIdChatTabSendButton.svelte'
  import OrderByIdChatTabAttachButton from '../OrderByIdChatTabAttachButton/OrderByIdChatTabAttachButton.svelte'
  import OrderByIdChatForm from '../OrderByIdChatForm/OrderByIdChatForm.svelte'
</script>

<OrderByIdChatTabFooterLayout>
  <DividerElement />
  <OrderByIdChatForm>
    <OrderByIdChatTabAttachButton />
    <OrderByIdChatTabTextArea />
    <OrderByIdChatTabSendButton />
  </OrderByIdChatForm>
</OrderByIdChatTabFooterLayout>
