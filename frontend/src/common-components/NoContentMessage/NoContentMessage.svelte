<script lang="ts">
  import NoContentMessageLayout from './NoContentMessageLayout.svelte'

  type NoContentMessageProps = {
    text: string
  }

  const { text }: NoContentMessageProps = $props()
</script>

<NoContentMessageLayout>
  {text}
</NoContentMessageLayout>
