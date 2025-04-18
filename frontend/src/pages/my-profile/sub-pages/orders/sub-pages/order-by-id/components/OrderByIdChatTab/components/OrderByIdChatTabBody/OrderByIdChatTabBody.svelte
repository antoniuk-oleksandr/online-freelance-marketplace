<script lang="ts">
  import type { ChatData } from '@/types/ChatData'
  import OrderByIdChatTabBodyLayout from './OrderByIdChatTabBodyLayout.svelte'
  import OrderByIdChatTabMessageList from '../OrderByIdChatTabMessageList/OrderByIdChatTabMessageList.svelte'

  type OrderByIdChatTabBodyProps = {
    chatData: ChatData
  }

  const props: OrderByIdChatTabBodyProps = $props()
</script>

<OrderByIdChatTabBodyLayout>
  <OrderByIdChatTabMessageList {...props} />
</OrderByIdChatTabBodyLayout>
