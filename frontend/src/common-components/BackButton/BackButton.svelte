<script lang="ts">
  type BackButtonProps = {
    text?: string
    styles?: string
  }

  const { text = 'Back', styles }: BackButtonProps = $props()

  const back = () => {
    window.history.back()
  }
</script>

<button
  type="button"
  onclick={back}
  class="{styles} text-lg text-cyan-500 hover:underline duration-200"
>
  {text}
</button>
