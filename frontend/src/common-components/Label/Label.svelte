<script lang="ts">
  type LabelProps = {
    text: string
  }
  const { text }: LabelProps = $props()
</script>

<span
  class="text-light-palette-text-primary dark:text-dark-palette-text-primary font-medium text-base leading-6"
  >{text}</span
>
