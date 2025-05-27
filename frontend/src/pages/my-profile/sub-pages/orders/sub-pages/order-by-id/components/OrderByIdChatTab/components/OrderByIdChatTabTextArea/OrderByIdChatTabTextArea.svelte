<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { calcChatTextAreaRows } from '../../helpers'

  let value = $state('')
  let textarea = $state<HTMLTextAreaElement | null>(null)

  const calculateRows = () => calcChatTextAreaRows(textarea)

  onMount(() => {
    calculateRows()
    window.addEventListener('resize', calculateRows)
  })

  onDestroy(() => {
    window.removeEventListener('resize', calculateRows)
  })
</script>

<!-- svelte-ignore element_invalid_self_closing_tag -->
<textarea
  name="message"
  bind:value
  maxlength="2000"
  bind:this={textarea}
  class="w-full bg-transparent px-3 outline-none resize-none"
  autocomplete="off"
  placeholder="Type a message"
  oninput={calculateRows}
  rows={1}
/>
