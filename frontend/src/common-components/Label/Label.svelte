<script lang="ts">
  type LabelProps = {
    text: string
    styles?: string
  }
  const { text, styles = '' }: LabelProps = $props()
</script>

<span
  class="{styles} text-light-palette-text-primary dark:text-dark-palette-text-primary font-medium text-base leading-6"
  >{text}</span
>
