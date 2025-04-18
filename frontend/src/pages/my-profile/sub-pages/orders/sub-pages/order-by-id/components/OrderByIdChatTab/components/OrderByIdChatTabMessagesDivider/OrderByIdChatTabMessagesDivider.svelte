<script lang="ts">
  import type { ChatMessage } from '@/types/ChatMessage'
  import { formatDateSeparator, isFirstMessageOfDay } from '../../helpers'

  type OrderByIdChatTabMessagesDividerProps = {
    index: number
    messages: ChatMessage[]
  }

  const { messages, index }: OrderByIdChatTabMessagesDividerProps = $props()

  let differentDate = $state(isFirstMessageOfDay(index, messages))
</script>

{#if differentDate}
  <div class="w-full text-center">{formatDateSeparator(messages[index].sentAt)}</div>
{/if}
