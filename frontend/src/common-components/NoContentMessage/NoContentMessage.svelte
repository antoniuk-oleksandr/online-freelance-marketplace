<script lang="ts">
  import NoContentMessageLayout from './NoContentMessageLayout.svelte'

  type NoContentMessageProps = {
    text: string
    styles?: string
  }

  const { text, styles }: NoContentMessageProps = $props()
</script>

<NoContentMessageLayout {styles}>
  <p>{text}</p>
</NoContentMessageLayout>
